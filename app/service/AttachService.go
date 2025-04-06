package service

import (
	"fmt"
	"github.com/wiselike/leanote-of-unofficial/app/db"
	"github.com/wiselike/leanote-of-unofficial/app/info"
	. "github.com/wiselike/leanote-of-unofficial/app/lea"
	"gopkg.in/mgo.v2/bson"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
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

	userId := attach.UploadUserId.Hex()

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
func (this *AttachService) DeleteAllAttachs(userId, noteId string) bool {
	note := noteService.GetNoteById(noteId)
	if note.UserId.Hex() == userId {
		attachs := []info.Attach{}
		db.ListByQ(db.Attachs, bson.M{"NoteId": bson.ObjectIdHex(noteId)}, &attachs)
		var fullPath string
		basePath := ConfigS.GlobalStringConfigs["files.dir"]
		for _, attach := range attachs {
			fullPath = path.Join(basePath, attach.Path)
			DeleteFile(fullPath)
		}
		if fullPath != "" {
			DeleteFile(path.Dir(fullPath))
		}
		db.DeleteAll(db.Attachs, bson.M{"NoteId": bson.ObjectIdHex(noteId)})
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
			var fullPath string
			fullPath = path.Join(ConfigS.GlobalStringConfigs["files.dir"], attach.Path)
			if ok := DeleteFile(fullPath); ok {
				// 修改note Usn
				noteService.IncrNoteUsn(attach.NoteId.Hex(), userId)
				DeleteFile(path.Dir(fullPath))
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
		filePath := path.Join(toUserId, "attachs", this.getAttachNoteTitle(attach.Path), newFilename)
		os.MkdirAll(filepath.Dir(filePath), 0755)
		_, err := CopyFile(path.Join(basePath, attach.Path), path.Join(basePath, filePath))
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
		if !nowAttachs[fileId] { // 需要删除的
			this.DeleteAttach(fileId, userId)
		}
	}

	return false

}

var noteAttachReg = regexp.MustCompile("(getAttach)\\?fileId=([a-z0-9A-Z]+)")

// 整理node附件，只需要从数据库里查NoteId就能获取所有附件清单
func (this *AttachService) ReOrganizeAttachFiles(userId, noteId, title string) bool {
	if title == "" {
		return false // title不存在就不整理了，下次再整理
	}
	title = FixFilename(title)
	attachs := []info.Attach{}
	db.ListByQ(db.Attachs, bson.M{"UploadUserId": bson.ObjectIdHex(userId), "NoteId": bson.ObjectIdHex(noteId)}, &attachs)
	sort.Slice(attachs, func(i, j int) bool { return attachs[i].AttachId < attachs[j].AttachId })

	var oldFullPath, newFullPath string
	basePath := ConfigS.GlobalStringConfigs["files.dir"]
	for i, attach := range attachs {
		var newAttachPath string
		// 判断需要重命名和移动attach
		attachName := filepath.Base(attach.Path)
		if oldAttachTitle := this.getAttachNoteTitle(attach.Path); oldAttachTitle != title || !strings.HasPrefix(attachName, strconv.Itoa(i)+"_") {
			fName := strings.Split(attachName, "_")
			newAttachPath = filepath.Join(userId, "/attachs/", title, fmt.Sprintf("%d_%s", i, fName[len(fName)-1]))

			oldFullPath = filepath.Join(basePath, attach.Path)
			newFullPath = filepath.Join(basePath, newAttachPath)
			if err := os.MkdirAll(filepath.Dir(newFullPath), 0755); err != nil {
				return false
			}
			if err := MoveFile(oldFullPath, newFullPath); err == nil {
				// 更新数据库
				attach.Path = newAttachPath
				if ok := db.Update(db.Attachs, bson.M{"UploadUserId": bson.ObjectIdHex(userId), "_id": attach.AttachId}, &attach); !ok {
					// 数据库写失败，回滚
					MoveFile(newFullPath, oldFullPath)
					continue
				}
			}
		}
	}

	if oldFullPath != "" {
		DeleteFile(path.Dir(oldFullPath))
	}
	return true
}

func (this *AttachService) getAttachNoteTitle(path string) string {
	// file.Path值是相对路径：path.Join(userId, "attachs", title, "xxx.txt")
	paths := strings.Split(filepath.Clean(path), string(filepath.Separator))

	if len(paths) == 4 {
		return paths[2]
	}
	return ""
}
