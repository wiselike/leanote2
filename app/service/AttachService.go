package service

import (
	"fmt"
	"github.com/wiselike/leanote-of-unofficial/app/db"
	"github.com/wiselike/leanote-of-unofficial/app/info"
	. "github.com/wiselike/leanote-of-unofficial/app/lea"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

type AttachService struct {
}

// add attach
// api调用时, 添加attach之前是没有note的
// fromApi表示是api添加的, updateNote传过来的, 此时不要incNote's usn, 因为updateNote会inc的
func (this *AttachService) AddAttach(attach info.Attach, fromApi bool) (ok bool, msg string) {
	attach.CreatedTime = time.Now()
	ok = db.Insert(db.Attachs, attach)

	note := noteService.GetNoteById(attach.NoteId.Hex())

	// api调用时, 添加attach之前是没有note的
	var userId string
	if note.NoteId != "" {
		userId = note.UserId.Hex()
	} else {
		userId = attach.UploadUserId.Hex()
	}

	if ok {
		// 更新笔记的attachs num
		this.updateNoteAttachNum(attach.NoteId, 1)
	}

	if !fromApi {
		// 增长note's usn
		noteService.IncrNoteUsn(attach.NoteId.Hex(), userId)
	}

	return
}

// 更新笔记的附件个数
// addNum 1或-1
func (this *AttachService) updateNoteAttachNum(noteId bson.ObjectId, addNum int) bool {
	num := db.Count(db.Attachs, bson.M{"NoteId": noteId})
	/*
		note := info.Note{}
		note = noteService.GetNoteById(noteId.Hex())
		note.AttachNum += addNum
		if note.AttachNum < 0 {
			note.AttachNum = 0
		}
		Log(note.AttachNum)
	*/
	return db.UpdateByQField(db.Notes, bson.M{"_id": noteId}, "AttachNum", num)
}

// list attachs
func (this *AttachService) ListAttachs(noteId, userId string) []info.Attach {
	attachs := []info.Attach{}

	// 判断是否有权限为笔记添加附件, userId为空时表示是分享笔记的附件
	if userId != "" && !shareService.HasUpdateNotePerm(noteId, userId) {
		return attachs
	}

	// 笔记是否是自己的
	note := noteService.GetNoteByIdAndUserId(noteId, userId)
	if note.NoteId == "" {
		return attachs
	}

	// TODO 这里, 优化权限控制

	db.ListByQ(db.Attachs, bson.M{"NoteId": bson.ObjectIdHex(noteId)}, &attachs)

	return attachs
}

// api调用, 通过noteIds得到note's attachs, 通过noteId归类返回
func (this *AttachService) getAttachsByNoteIds(noteIds []bson.ObjectId) map[string][]info.Attach {
	attachs := []info.Attach{}
	db.ListByQ(db.Attachs, bson.M{"NoteId": bson.M{"$in": noteIds}}, &attachs)
	noteAttchs := make(map[string][]info.Attach)
	for _, attach := range attachs {
		noteId := attach.NoteId.Hex()
		if itAttachs, ok := noteAttchs[noteId]; ok {
			noteAttchs[noteId] = append(itAttachs, attach)
		} else {
			noteAttchs[noteId] = []info.Attach{attach}
		}
	}
	return noteAttchs
}

func (this *AttachService) UpdateImageTitle(userId, fileId, title string) bool {
	return db.UpdateByIdAndUserIdField(db.Files, fileId, userId, "Title", title)
}

// Delete note to delete attas firstly
func (this *AttachService) DeleteAllAttachs(noteId, userId string) bool {
	note := noteService.GetNoteById(noteId)
	if note.UserId.Hex() == userId {
		attachs := []info.Attach{}
		db.ListByQ(db.Attachs, bson.M{"NoteId": bson.ObjectIdHex(noteId)}, &attachs)
		for _, attach := range attachs {
			os.Remove(path.Join(ConfigS.GlobalStringConfigs["files.dir"], attach.Path))
		}
		return true
	}

	return false
}

// delete attach
// 删除附件为什么要incrNoteUsn ? 因为可能没有内容要修改的
func (this *AttachService) DeleteAttach(attachId, userId string) (bool, string) {
	attach := info.Attach{}
	db.Get(db.Attachs, attachId, &attach)

	if attach.AttachId != "" {
		// 判断是否有权限为笔记添加附件
		if !shareService.HasUpdateNotePerm(attach.NoteId.Hex(), userId) {
			return false, "No Perm"
		}

		if db.Delete(db.Attachs, bson.M{"_id": bson.ObjectIdHex(attachId)}) {
			this.updateNoteAttachNum(attach.NoteId, -1)
			attach.Path = strings.TrimLeft(attach.Path, "/")
			err := os.Remove(path.Join(ConfigS.GlobalStringConfigs["files.dir"], attach.Path))
			if err == nil {
				// userService.UpdateAttachSize(note.UserId.Hex(), -attach.Size)
				// 修改note Usn
				noteService.IncrNoteUsn(attach.NoteId.Hex(), userId)

				return true, "delete file success"
			}
			return false, "delete file error"
		}
		return false, "db error"
	}
	return false, "no such item"
}

// 获取文件路径
// 要判断是否具有权限
// userId是否具有attach的访问权限
func (this *AttachService) GetAttach(attachId, userId string) (attach info.Attach) {
	if attachId == "" {
		return
	}

	attach = info.Attach{}
	db.Get(db.Attachs, attachId, &attach)
	path := attach.Path
	if path == "" {
		return
	}

	note := noteService.GetNoteById(attach.NoteId.Hex())

	// 判断权限

	// 笔记是否是公开的
	if note.IsBlog {
		return
	}

	// 笔记是否是我的
	if note.UserId.Hex() == userId {
		return
	}

	// 我是否有权限查看或协作
	if shareService.HasReadNotePerm(attach.NoteId.Hex(), userId) {
		return
	}

	attach = info.Attach{}
	return
}

// 复制笔记时需要复制附件
// noteService调用, 权限已判断
func (this *AttachService) CopyAttachs(noteId, toNoteId, toUserId string) bool {
	attachs := []info.Attach{}
	db.ListByQ(db.Attachs, bson.M{"NoteId": bson.ObjectIdHex(noteId)}, &attachs)

	// 复制之
	basePath := ConfigS.GlobalStringConfigs["files.dir"]
	toNoteIdO := bson.ObjectIdHex(toNoteId)
	for _, attach := range attachs {
		attach.AttachId = ""
		attach.NoteId = toNoteIdO

		// 文件复制一份
		_, ext := SplitFilename(attach.Name)
		newFilename := NewGuid() + ext
		dir := toUserId + "/attachs"
		filePath := path.Join(dir, newFilename)
		err := os.MkdirAll(path.Join(basePath, dir), 0755)
		if err != nil {
			return false
		}
		_, err = CopyFile(path.Join(basePath, attach.Path), path.Join(basePath, filePath))
		if err != nil {
			return false
		}
		attach.Name = newFilename
		attach.Path = filePath

		this.AddAttach(attach, false)
	}

	return true
}

// 只留下files的数据, 其它的都删除
func (this *AttachService) UpdateOrDeleteAttachApi(noteId, userId string, files []info.NoteFile) bool {
	// 现在数据库内的
	attachs := this.ListAttachs(noteId, userId)

	nowAttachs := map[string]bool{}
	if files != nil {
		for _, file := range files {
			if file.IsAttach && file.FileId != "" {
				nowAttachs[file.FileId] = true
			}
		}
	}

	for _, attach := range attachs {
		fileId := attach.AttachId.Hex()
		if !nowAttachs[fileId] {
			// 需要删除的
			// TODO 权限验证去掉
			this.DeleteAttach(fileId, userId)
		}
	}

	return false

}

var noteAttachReg = regexp.MustCompile("(getAttach)\\?fileId=([a-z0-9A-Z]+)")

// 整理node附件，按标题来存放，以便于到服务器上检索维护
func (this *AttachService) OrganizeAttachFiles(userId, title, content string) (rmDir string) {
	// 获取所有的fileId
	find := noteAttachReg.FindAllStringSubmatch(content, -1) // 查找
	if find == nil || len(find) < 1 {
		return
	}
	find = DeduplicateMatches(find)

	// 格式化titile
	title = FixFilename(title)
	if title == "" {
		title = "empty-titles-set"
	}

	basePath := ConfigS.GlobalStringConfigs["files.dir"]
	newDbPathDir := path.Join(GetRandomFilePath(userId, ""), "/attachs/", title)
	newPathDir := path.Join(basePath, newDbPathDir)
	if err := os.MkdirAll(newPathDir, 0755); err != nil {
		return
	}
	for i, each := range find {
		if each != nil && len(each) == 3 {
			// 查找原路径
			file := info.Attach{}
			if db.Get(db.Attachs, each[2], &file); file.Path != "" {
				// 创建文件名，并移动路径
				oldFullPath := path.Join(basePath, file.Path)
				fname := strings.Split(path.Base(file.Path), "_")
				file.Path = path.Join(newDbPathDir, fmt.Sprintf("%d_%s", i, fname[len(fname)-1]))
				newFullPath := path.Join(basePath, file.Path)
				if oldFullPath != newFullPath {
					if err := os.Rename(oldFullPath, newFullPath); err == nil {
						// 更新数据库
						if ok := db.Update(db.Attachs, bson.M{"_id": bson.ObjectIdHex(each[2])}, file); !ok {
							// 数据库写失败，回滚
							os.Rename(newFullPath, oldFullPath)
							continue
						}
						// 保存第一个附件的文件夹，作为旧路径，用于删除
						if rmDir == "" {
							rmDir = path.Dir(oldFullPath)
						}
					}
				}
			}
		}
	}

	// 没有移动任何文件的话，则仅删除空文件夹
	if rmDir == "" {
		os.Remove(newPathDir)
		return
	}

	// 带文件夹结束符，避免比较到部分文件名
	// 避免删除空标题集合文件夹
	if strings.HasPrefix(newPathDir+"/", rmDir+"/") || path.Base(rmDir) == "empty-titles-set" {
		return "" // 不删除
	}
	return
}

// 整理node附件，同上。带删除旧文件夹
func (this *AttachService) ReOrganizeAttachFiles(userId, noteId, title, content string, hasTitle, hasContent bool) bool {
	if !hasTitle && !hasContent {
		return true
	}
	if !hasTitle { // 获取title
		title = noteService.GetNote(noteId, userId).Title
	}
	if !hasContent { // 获取content
		content = noteService.GetNoteContent(noteId, userId).Content
	}

	if oldDir := this.OrganizeAttachFiles(userId, title, content); oldDir != "" {
		// 删旧的空文件夹，如果仅部分文件移动，不应删除整个文件夹，因为可能发生从其他笔记里拷贝
		dir, _ := ioutil.ReadDir(oldDir)
		if len(dir) == 0 {
			os.RemoveAll(oldDir)
		}
	}
	return true
}
