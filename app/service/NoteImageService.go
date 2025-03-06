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
	//	"time"
)

type NoteImageService struct {
}

var noteImageReg = regexp.MustCompile("(outputImage|getImage)\\?fileId=([a-z0-9A-Z]+)")

// 通过id, userId得到noteIds
func (this *NoteImageService) GetNoteIds(imageId string) []bson.ObjectId {
	noteImages := []info.NoteImage{}
	db.ListByQWithFields(db.NoteImages, bson.M{"ImageId": bson.ObjectIdHex(imageId)}, []string{"NoteId"}, &noteImages)

	if noteImages != nil && len(noteImages) > 0 {
		noteIds := make([]bson.ObjectId, len(noteImages))
		cnt := len(noteImages)
		for i := 0; i < cnt; i++ {
			noteIds[i] = noteImages[i].NoteId
		}
		return noteIds
	}

	return nil
}

// TODO 这个web可以用, 但api会传来, 不用用了
// 解析内容中的图片, 建立图片与note的关系
// <img src="/file/outputImage?fileId=12323232" />
// 图片必须是我的, 不然不添加
// imgSrc 防止博客修改了, 但内容删除了
func (this *NoteImageService) UpdateNoteImages(userId, noteId, imgSrc, content string) bool {
	// 让主图成为内容的一员
	if imgSrc != "" {
		content = "<img src=\"" + imgSrc + "\" >" + content
	}
	// life 添加getImage
	find := noteImageReg.FindAllStringSubmatch(content, -1) // 查找所有的

	// 删除旧的
	db.DeleteAll(db.NoteImages, bson.M{"NoteId": bson.ObjectIdHex(noteId)})

	// 添加新的
	var fileId string
	noteImage := info.NoteImage{NoteId: bson.ObjectIdHex(noteId)}
	hasAdded := make(map[string]bool)
	if find != nil && len(find) > 0 {
		for _, each := range find {
			if each != nil && len(each) == 3 {
				fileId = each[2] // 现在有两个子表达式了
				// 之前没能添加过的
				if _, ok := hasAdded[fileId]; !ok {
					Log(fileId)
					// 判断是否是我的文件
					if fileService.IsMyFile(userId, fileId) {
						noteImage.ImageId = bson.ObjectIdHex(fileId)
						db.Insert(db.NoteImages, noteImage)
					}
					hasAdded[fileId] = true
				}
			}
		}
	}

	return true
}

// 复制图片, 把note的图片都copy给我, 且修改noteContent图片路径
func (this *NoteImageService) CopyNoteImages(fromNoteId, fromUserId, newNoteId, content, toUserId string) string {
	/* 弃用之
	// 得到fromNoteId的noteImages, 如果为空, 则直接返回content
	noteImages := []info.NoteImage{}
	db.ListByQWithFields(db.NoteImages, bson.M{"NoteId": bson.ObjectIdHex(fromNoteId)}, []string{"ImageId"}, &noteImages)
	if len(noteImages) == 0 {
		return content;
	}
	for _, noteImage := range noteImages {
		imageId := noteImage.ImageId.Hex()
		ok, newImageId := fileService.CopyImage(fromUserId, imageId, toUserId)
		if ok {
			replaceMap[imageId] = newImageId
		}
	}
	*/

	// 因为很多图片上传就会删除, 所以直接从内容中查看图片id进行复制

	// <img src="/file/outputImage?fileId=12323232" />
	// 把fileId=1232替换成新的
	replaceMap := map[string]string{}

	content = noteImageReg.ReplaceAllStringFunc(content, func(each string) string {
		// each = outputImage?fileId=541bd2f599c37b4f3r000003
		// each = getImage?fileId=541bd2f599c37b4f3r000003

		fileId := each[len(each)-24:] // 得到后24位, 也即id

		if _, ok := replaceMap[fileId]; !ok {
			if bson.IsObjectIdHex(fileId) {
				ok2, newImageId := fileService.CopyImage(fromUserId, fileId, toUserId)
				if ok2 {
					replaceMap[fileId] = newImageId
				} else {
					replaceMap[fileId] = ""
				}
			} else {
				replaceMap[fileId] = ""
			}
		}

		replaceFileId := replaceMap[fileId]
		if replaceFileId != "" {
			if each[0] == 'o' {
				return "outputImage?fileId=" + replaceFileId
			}
			return "getImage?fileId=" + replaceFileId
		}
		return each
	})

	return content
}

func (this *NoteImageService) getImagesByNoteIds(noteIds []bson.ObjectId) map[string][]info.File {
	noteNoteImages := []info.NoteImage{}
	db.ListByQ(db.NoteImages, bson.M{"NoteId": bson.M{"$in": noteIds}}, &noteNoteImages)

	// 得到imageId, 再去files表查所有的Files
	imageIds := []bson.ObjectId{}

	// 图片1 => N notes
	imageIdNotes := map[string][]string{} // imageId => [noteId1, noteId2, ...]
	for _, noteImage := range noteNoteImages {
		imageId := noteImage.ImageId
		imageIds = append(imageIds, imageId)

		imageIdHex := imageId.Hex()
		noteId := noteImage.NoteId.Hex()
		if notes, ok := imageIdNotes[imageIdHex]; ok {
			imageIdNotes[imageIdHex] = append(notes, noteId)
		} else {
			imageIdNotes[imageIdHex] = []string{noteId}
		}
	}

	// 得到所有files
	files := []info.File{}
	db.ListByQ(db.Files, bson.M{"_id": bson.M{"$in": imageIds}}, &files)

	// 建立note->file关联
	noteImages := make(map[string][]info.File)
	for _, file := range files {
		fileIdHex := file.FileId.Hex() // == imageId
		// 这个fileIdHex有哪些notes呢?
		if notes, ok := imageIdNotes[fileIdHex]; ok {
			for _, noteId := range notes {
				if files, ok2 := noteImages[noteId]; ok2 {
					noteImages[noteId] = append(files, file)
				} else {
					noteImages[noteId] = []info.File{file}
				}
			}
		}
	}
	return noteImages
}

// 整理node图片，按标题来存放，以便于到服务器上检索维护
func (this *NoteImageService) OrganizeImageFiles(userId, title, content string) (rmDir string) {
	// 获取所有的imgId
	find := noteImageReg.FindAllStringSubmatch(content, -1) // 查找
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
	newDbPathDir := path.Join(GetRandomFilePath(userId, ""), "/images/", title)
	newPathDir := path.Join(basePath, newDbPathDir)
	if err := os.MkdirAll(newPathDir, 0755); err != nil {
		return
	}
	for i, each := range find {
		if each != nil && len(each) == 3 {
			// 查找原路径
			file := &info.File{}
			if db.GetByIdAndUserId(db.Files, each[2], userId, file); file.Path != "" {
				// 创建文件名，并移动路径
				oldFullPath := path.Join(basePath, file.Path)
				fname := strings.Split(path.Base(file.Path), "_")
				file.Path = path.Join(newDbPathDir, fmt.Sprintf("%d_%s", i, fname[len(fname)-1]))
				newFullPath := path.Join(basePath, file.Path)
				if oldFullPath != newFullPath {
					if err := os.Rename(oldFullPath, newFullPath); err == nil {
						// 更新数据库
						if ok := db.UpdateByIdAndUserId(db.Files, each[2], userId, file); !ok {
							// 数据库写失败，回滚
							os.Rename(newFullPath, oldFullPath)
							continue
						}
						// 保存第一张图片的文件夹，作为旧路径，用于删除
						if rmDir == "" {
							rmDir = path.Dir(oldFullPath)
						}
					}
				}
			}
		}
	}

	// 没有移动任何图片的话，则仅删除空文件夹
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

// 整理node图片，同上。带删除旧文件夹
func (this *NoteImageService) ReOrganizeImageFiles(userId, noteId, title, content string, hasTitle, hasContent bool) bool {
	if !hasTitle && !hasContent {
		return true
	}
	if !hasTitle { // 获取title
		title = noteService.GetNote(noteId, userId).Title
	}
	if !hasContent { // 获取content
		content = noteService.GetNoteContent(noteId, userId).Content
	}

	if oldDir := this.OrganizeImageFiles(userId, title, content); oldDir != "" {
		// 删旧的空文件夹，如果仅部分文件移动，不应删除整个文件夹，因为可能发生从其他笔记里拷贝
		dir, _ := ioutil.ReadDir(oldDir)
		if len(dir) == 0 {
			os.RemoveAll(oldDir)
		}
	}
	return true
}
