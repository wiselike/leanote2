package service

import (
	"path"
	"sort"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote-of-unofficial/app/db"
	"github.com/wiselike/leanote-of-unofficial/app/info"
	. "github.com/wiselike/leanote-of-unofficial/app/lea"
)

// 历史记录
type NoteContentHistoryService struct {
}

// 新建一个note, 不添加历史记录
// 添加历史，在数据库中倒序存放：前面的是老的，后面是新的（与原来的顺序相反）
func (this *NoteContentHistoryService) AddHistory(noteId, userId string, newHistory info.EachHistory) {
	// 检查是否是空
	if newHistory.Content == "" {
		return
	}

	// 每个历史记录最大值
	maxSize := ConfigS.GlobalAllConfigs["note.history.size"].(int)
	if maxSize < 1 {
		return
	}

	var historiesLenth int
	history := &info.NoteContentHistory{}
	db.GetByIdAndUserId(db.NoteContentHistories, noteId, userId, history)
	if history.NoteId == "" {
		historiesLenth = -1
	} else {
		historiesLenth = len(history.Histories)
	}

	if historiesLenth == -1 { // historiesLenth==0时，也不能newHistory，必须pushHistory
		this.newHistory(noteId, userId, newHistory)
	} else {
		// 读取最新的历史记录，判断是否是AutoBackup；
		if historiesLenth > 0 && history.Histories[historiesLenth-1].IsAutoBackup {
			// 清理不再用的图片
			this.CleanImage(history, &newHistory, historiesLenth-1)

			db.UpdateByIdAndUserIdPop(db.NoteContentHistories, noteId, userId, "Histories", 1)
		} else if historiesLenth >= maxSize { // 判断是否超出 maxSize, 如果是则pop掉一个最老的
			// 清理不再用的图片
			this.CleanImage(history, &newHistory, 0)

			db.UpdateByIdAndUserIdPop(db.NoteContentHistories, noteId, userId, "Histories", -1)
		}

		// 插入一个历史记录，只能后插
		db.UpdateByIdAndUserIdPush(db.NoteContentHistories, noteId, userId, "Histories", newHistory)
	}

	return
}

// 更新一下最后一条历史记录的状态，由自动历史转为手动历史
func (this *NoteContentHistoryService) UpdateHistoryBackupState(noteId, userId string, isAutoBackup bool) {
	// mongo2没法找到最后数组的最后一个，
	// 所以这里进行了折中，找到第一个IsAutoBackup为true的项
	// 将其替换为isAutoBackup值
	db.UpdateHistoryBackupState(db.NoteContentHistories, noteId, userId, isAutoBackup)
}

// 新建历史
func (this *NoteContentHistoryService) newHistory(noteId, userId string, newHistory info.EachHistory) {
	history := &info.NoteContentHistory{NoteId: bson.ObjectIdHex(noteId),
		UserId:    bson.ObjectIdHex(userId),
		Histories: []info.EachHistory{newHistory},
	}

	// 保存之
	db.Insert(db.NoteContentHistories, history)
}

// 列表展示
func (this *NoteContentHistoryService) ListHistories(noteId, userId string) []info.EachHistory {
	histories := info.NoteContentHistory{}
	db.GetByIdAndUserId(db.NoteContentHistories, noteId, userId, &histories)
	sort.Sort(info.EachHistorySlice(histories.Histories)) // 前端倒着展示，便于理解和操作
	return histories.Histories
}

// 删除一条历史；
// 使用历史记录的时间戳，作为标志进行查找并删除；
// 实际过程中应该不存在两条时间戳完全相同(时间戳是精确到毫秒级的)历史记录；
// 如果确实存在两条时间戳毫秒级也相同的，则内容肯定也相同，会一起都删除，目前还没遇到此情况
func (this *NoteContentHistoryService) DeleteHistory(noteId, userId, timeToDel string) {
	// 自动解析js返回的RFC 3339格式化时间戳。
	// golang可以自动解析末尾的Z或者时区偏移
	t, err := time.Parse(time.RFC3339, timeToDel)
	if err != nil {
		return
	}

	// 清理不再用的图片
	history := &info.NoteContentHistory{}
	db.GetByIdAndUserId(db.NoteContentHistories, noteId, userId, history)
	getI := func() int {
		for i, each := range history.Histories {
			if each.UpdatedTime == t {
				return i
			}
		}
		return -1
	}
	this.CleanImage(history, nil, getI())

	db.DeleteOneHistory(db.NoteContentHistories, noteId, userId, t)
	return
}

func (this *NoteContentHistoryService) CleanImage(history *info.NoteContentHistory, newHistory *info.EachHistory, num int) {
	if len(history.Histories) <= num {
		return
	}

	noteId := history.NoteId.Hex()
	userId := history.UserId.Hex()

	//
	findDelete := noteImageReg.FindAllStringSubmatch(history.Histories[num].Content, -1) // 查找
	if findDelete == nil || len(findDelete) < 1 {
		return
	}
	findDelete = DeduplicateMatches(findDelete)

	// 获取history、newHistory里的images
	findAll := make([][]string, 0, 10)
	for i := len(history.Histories) - 1; i >= 0; i-- {
		if i == num {
			continue
		}
		find := noteImageReg.FindAllStringSubmatch(history.Histories[i].Content, -1) // 查找
		if find == nil || len(find) < 1 {
			continue
		}
		findAll = append(findAll, find...)
	}
	if newHistory != nil {
		find := noteImageReg.FindAllStringSubmatch(newHistory.Content, -1) // 查找
		if find != nil && len(find) > 1 {
			findAll = append(findAll, find...)
		}
	}
	findAll = DeduplicateMatches(findAll)

	findDelete = SliceMinus(findDelete, findAll)
	if len(findDelete) < 1 {
		return
	}

	// 获取NoteImages，也就是Content里的images
	note_imageIDs_tmp := []info.NoteImage{}
	db.ListByQ(db.NoteImages, bson.M{"NoteId": bson.ObjectIdHex(noteId)}, &note_imageIDs_tmp)
	note_imageIDs := make(map[string]bool)
	for i := range note_imageIDs_tmp {
		note_imageIDs[note_imageIDs_tmp[i].ImageId.Hex()] = true
	}

	basePath := ConfigS.GlobalStringConfigs["files.dir"]
	var fullPath string
	for _, each := range findDelete {
		if each != nil && len(each) == 3 {
			image_id := each[2]
			if _, ok := note_imageIDs[image_id]; !ok { // 要删除
				needDelete := true

				// 判断其他笔记是否有用此图片
				noteImages := noteImageService.GetNoteIds(image_id)
				noteIdHex := bson.ObjectIdHex(noteId)
				for i := range noteImages {
					if noteImages[i] != noteIdHex {
						needDelete = false
						break // 其他笔记有用此图片，不删除
					}
				}

				if needDelete {
					file := &info.File{}
					if db.GetByIdAndUserId(db.Files, image_id, userId, file); file.Path != "" {
						if db.DeleteByIdAndUserId(db.Files, image_id, userId) {
							fullPath = path.Join(basePath, file.Path)
							DeleteFile(fullPath)
						}
					}
				}
			}
		}
	}
	if fullPath != "" {
		DeleteFile(path.Dir(fullPath))
	}

	return
}

// 删除history里的所有图片，并清理history数据库
func (this *NoteContentHistoryService) CleanHistoryAndImages(userId, noteId string) bool {
	histories := &info.NoteContentHistory{}
	db.GetByIdAndUserId(db.NoteContentHistories, noteId, userId, histories)

	// 获取history、newHistory里的images
	findDelete := make([][]string, 0, 10)
	for i := len(histories.Histories) - 1; i >= 0; i-- {
		find := noteImageReg.FindAllStringSubmatch(histories.Histories[i].Content, -1) // 查找
		if find == nil || len(find) < 1 {
			continue
		}
		findDelete = append(findDelete, find...)
	}
	findDelete = DeduplicateMatches(findDelete)

	basePath := ConfigS.GlobalStringConfigs["files.dir"]
	var fullPath string
	for _, each := range findDelete {
		if each != nil && len(each) == 3 {
			image_id := each[2]
			needDelete := true

			// 判断其他笔记是否有用此图片
			noteImages := noteImageService.GetNoteIds(image_id)
			noteIdHex := bson.ObjectIdHex(noteId)
			for i := range noteImages {
				if noteImages[i] != noteIdHex {
					needDelete = false
					break // 其他笔记有用此图片，不删除
				}
			}

			if needDelete {
				file := &info.File{}
				if db.GetByIdAndUserId(db.Files, image_id, userId, file); file.Path != "" {
					if db.DeleteByIdAndUserId(db.Files, image_id, userId) {
						fullPath = path.Join(basePath, file.Path)
						DeleteFile(fullPath)
					}
				}
			}
		}
	}
	if fullPath != "" {
		DeleteFile(path.Dir(fullPath))
	}

	return db.DeleteByIdAndUserId(db.NoteContentHistories, noteId, userId)
}
