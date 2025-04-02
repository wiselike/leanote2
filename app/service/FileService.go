package service

import (
	"encoding/base64"
	"fmt"
	"github.com/wiselike/leanote-of-unofficial/app/db"
	"github.com/wiselike/leanote-of-unofficial/app/info"
	. "github.com/wiselike/leanote-of-unofficial/app/lea"
	"github.com/wiselike/revel"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const DEFAULT_ALBUM_ID = "52d3e8ac99c37b7f0d000001"

type FileService struct {
}

// add Image
func (this *FileService) AddImage(image info.File, albumId, userId string, needCheckSize bool) (ok bool, msg string) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case string:
				if strings.Contains(v, "invalid input to ObjectIdHex") {
					msg = "Album not found! Please create an album first." //相册不存在，请先添加相册！
				} else {
					msg = "Album not found."
				}
			default:
				msg = "Internel Server Error"
			}
			ok = false
			return
		}
	}()
	image.CreatedTime = time.Now()
	if albumId != "" {
		image.AlbumId = bson.ObjectIdHex(albumId)
	} else {
		image.AlbumId = bson.ObjectIdHex(DEFAULT_ALBUM_ID)
		image.IsDefaultAlbum = true
	}
	image.UserId = bson.ObjectIdHex(userId)

	ok = db.Insert(db.Files, image)
	return
}

// list images
// if albumId == "" get default album images
func (this *FileService) ListImagesWithPage(userId, albumId, key string, pageNumber, pageSize int) info.Page {
	skipNum, sortFieldR := parsePageAndSort(pageNumber, pageSize, "CreatedTime", false)
	files := []info.File{}

	q := bson.M{"UserId": bson.ObjectIdHex(userId), "Type": ""} // life
	if albumId != "" {
		q["AlbumId"] = bson.ObjectIdHex(albumId)
	} else {
		// q["IsDefaultAlbum"] = true // 不再返回默认相册里的图片
		return info.Page{Count: 0, List: files}
	}
	if key != "" {
		q["Title"] = bson.M{"$regex": bson.RegEx{".*?" + key + ".*", "i"}}
	}

	//	LogJ(q)

	count := db.Count(db.Files, q)

	db.Files.
		Find(q).
		Sort(sortFieldR).
		Skip(skipNum).
		Limit(pageSize).
		All(&files)

	return info.Page{Count: count, List: files}
}

func (this *FileService) UpdateImageTitle(userId, fileId, title string) bool {
	return db.UpdateByIdAndUserIdField(db.Files, fileId, userId, "Title", title)
}

// get all images names
// for upgrade
func (this *FileService) GetAllImageNamesMap(userId string) (m map[string]bool) {
	q := bson.M{"UserId": bson.ObjectIdHex(userId)}
	files := []info.File{}
	db.ListByQWithFields(db.Files, q, []string{"Name"}, &files)

	m = make(map[string]bool)
	if len(files) == 0 {
		return
	}

	for _, file := range files {
		m[file.Name] = true
	}
	return
}

// delete image
func (this *FileService) DeleteImage(userId, fileId string) (bool, string) {
	file := info.File{}
	db.GetByIdAndUserId(db.Files, fileId, userId, &file)

	if file.FileId != "" {
		if db.DeleteByIdAndUserId(db.Files, fileId, userId) {
			// delete image
			// TODO
			var err error
			if strings.HasPrefix(file.Path, "/upload/") {
				Log(file.Path)
				err = os.Remove(path.Join(revel.BasePath, "/public/", file.Path))
			} else {
				fullPath := path.Join(ConfigS.GlobalStringConfigs["files.dir"], file.Path)
				err = os.Remove(fullPath)
				DeleteFile(path.Dir(fullPath)) // 只有空文件才会删除成功
			}
			if err == nil {
				return true, ""
			}
			return false, "delete file error!"
		}
		return false, "db error"
	}
	return false, "no such item"
}

// update image title
func (this *FileService) UpdateImage(userId, fileId, title string) bool {
	return db.UpdateByIdAndUserIdField(db.Files, fileId, userId, "Title", title)
}

func (this *FileService) GetFileBase64(userId, fileId string) (str string, mine string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()

	fpath := this.GetFile(userId, fileId)
	if fpath == "" {
		return "", ""
	}
	fpath = path.Join(ConfigS.GlobalStringConfigs["files.dir"], fpath)

	ff, err := ioutil.ReadFile(fpath)
	if err != nil {
		return "", ""
	}

	e64 := base64.StdEncoding
	maxEncLen := e64.EncodedLen(len(ff))
	encBuf := make([]byte, maxEncLen)

	e64.Encode(encBuf, ff)

	mime := http.DetectContentType(ff)

	str = string(encBuf)
	return str, mime
}

// 得到图片base64, 图片要在之前添加data:image/png;base64,
func (this *FileService) GetImageBase64(userId, fileId string) string {

	str, mime := this.GetFileBase64(userId, fileId)
	if str == "" {
		return ""
	}
	switch mime {
	case "image/gif", "image/jpeg", "image/pjpeg", "image/png", "image/tiff":
		return fmt.Sprintf("data:%s;base64,%s", mime, str)
	default:
	}
	return "data:image/png;base64," + str
}

// 获取文件路径
// 要判断是否具有权限
// userId是否具有fileId的访问权限
func (this *FileService) GetFile(userId, fileId string) string {
	if fileId == "" {
		return ""
	}

	file := info.File{}
	db.Get(db.Files, fileId, &file)
	path := file.Path
	if path == "" {
		return ""
	}

	// 1. 判断权限

	// 是否是我的文件
	if userId != "" && file.UserId.Hex() == userId {
		return path
	}

	// 得到使用过该fileId的所有笔记NoteId
	// 这些笔记是否有public的, 若有则ok
	// 这些笔记(笔记本)是否有共享给我的, 若有则ok

	noteIds := noteImageService.GetNoteIds(fileId)
	if noteIds != nil && len(noteIds) > 0 {
		// 这些笔记是否有public的
		if db.Has(db.Notes, bson.M{"_id": bson.M{"$in": noteIds}, "IsBlog": true}) {
			return path
		}

		if userId == "" {
			return ""
		}
		// 2014/12/28 修复, 如果是分享给用户组, 那就不行, 这里可以实现
		for _, noteId := range noteIds {
			note := noteService.GetNoteById(noteId.Hex())
			if shareService.HasReadPerm(note.UserId.Hex(), userId, noteId.Hex()) {
				return path
			}
		}
		/*
			// 若有共享给我的笔记?
			// 对该笔记可读?
			if db.Has(db.ShareNotes, bson.M{"ToUserId": bson.ObjectIdHex(userId), "NoteId": bson.M{"$in": noteIds}}) {
				return path
			}

			// 笔记本是否共享给我?
			// 通过笔记得到笔记本
			notes := []info.Note{}
			db.ListByQWithFields(db.Notes, bson.M{"_id": bson.M{"$in": noteIds}}, []string{"NotebookId"}, &notes)
			if notes != nil && len(notes) > 0 {
				notebookIds := make([]bson.ObjectId, len(notes))
				for i := 0; i < len(notes); i++ {
					notebookIds[i] = notes[i].NotebookId
				}

				if db.Has(db.ShareNotebooks, bson.M{"ToUserId": bson.ObjectIdHex(userId), "NotebookId": bson.M{"$in": notebookIds}}) {
					return path
				}
			}
		*/
	}

	// 可能是刚复制到owner上, 但内容又没有保存, 所以没有note->imageId的映射, 此时看是否有fromFileId
	if file.FromFileId != "" {
		fromFile := info.File{}
		db.Get2(db.Files, file.FromFileId, &fromFile)
		if fromFile.UserId.Hex() == userId {
			return fromFile.Path
		}
	}

	return ""
}

// 复制图片
func (this *FileService) CopyImage(userId, fileId, toUserId string) (bool, string) {
	if userId == "" || toUserId == "" {
		return false, ""
	}
	return this.CopyImageToTitle(userId, fileId, "", toUserId, nil)
}

func (this *FileService) CopyImageToTitle(userId, fileId, newTitle, toUserId string, file *info.File) (bool, string) {
	if file == nil {
		file = &info.File{}
		db.GetByIdAndUserId(db.Files, fileId, userId, file)
	}

	_, ext := SplitFilename(file.Name)
	newFilename := NewGuid() + ext

	basePath := ConfigS.GlobalStringConfigs["files.dir"]
	if newTitle == "" {
		newTitle = this.getImageNoteTitle(file.Path)
	}
	oldFilePath := file.Path
	file.Path = path.Join(toUserId, "images", newTitle, newFilename)
	fullFilePath := path.Join(basePath, file.Path)
	os.MkdirAll(filepath.Dir(fullFilePath), 0755)
	_, err := CopyFile(path.Join(basePath, oldFilePath), fullFilePath)
	if err != nil {
		return false, ""
	}

	fileInfo := info.File{Name: newFilename,
		Title:      file.Title,
		Path:       file.Path,
		Size:       file.Size,
		FromFileId: file.FileId}
	id := bson.NewObjectId()
	fileInfo.FileId = id
	fileId = id.Hex()
	Ok, _ := this.AddImage(fileInfo, "", toUserId, false)

	if Ok {
		return Ok, id.Hex()
	}
	return false, ""
}

// 是否是我的文件
func (this *FileService) IsMyFile(userId, fileId string) bool {
	// 如果有问题会panic
	if !bson.IsObjectIdHex(fileId) || !bson.IsObjectIdHex(userId) {
		return false
	}
	return db.Has(db.Files, bson.M{"UserId": bson.ObjectIdHex(userId), "_id": bson.ObjectIdHex(fileId)})
}

func (this *FileService) getImageNoteTitle(path string) string {
	// file.Path值是相对路径：path.Join(userId, "attachs", title, "xxx.txt")
	paths := strings.Split(filepath.Clean(path), string(filepath.Separator))

	if len(paths) == 4 {
		return paths[2]
	}
	return ""
}
