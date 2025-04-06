package service

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote-of-unofficial/app/db"
	"github.com/wiselike/leanote-of-unofficial/app/info"
	. "github.com/wiselike/leanote-of-unofficial/app/lea"
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

// 解析内容中的图片, 建立图片与note的关系
// <img src="/file/outputImage?fileId=12323232" />
func (this *NoteImageService) UpdateNoteImages(userId, noteId, content string) bool {
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

func (this *NoteImageService) DeleteNoteImages(userId, noteId string) bool {
	imageIDs := []info.NoteImage{}
	db.ListByQ(db.NoteImages, bson.M{"NoteId": bson.ObjectIdHex(noteId)}, &imageIDs)
	defer db.DeleteAll(db.NoteImages, bson.M{"NoteId": bson.ObjectIdHex(noteId)})

	var fullPath string
	basePath := ConfigS.GlobalStringConfigs["files.dir"]
	for _, image := range imageIDs {
		file := &info.File{}
		fileId := image.ImageId.Hex()
		if db.GetByIdAndUserId(db.Files, fileId, userId, file); file.Path != "" {
			if db.DeleteByIdAndUserId(db.Files, fileId, userId) {
				fullPath = path.Join(basePath, file.Path)
				DeleteFile(fullPath)
			}
		}
	}
	if fullPath != "" {
		DeleteFile(path.Dir(fullPath))
	}

	return true
}

// 复制图片, 把note的图片都copy给我, 且修改noteContent图片路径
func (this *NoteImageService) CopyNoteImages(fromNoteId, fromUserId, newNoteId, content, toUserId string) string {
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

// 整理node图片，按标题来存放，以便于到服务器上检索维护(不更新NoteImages，由UpdateNoteImages来更新)
// 返回值true代表content已更新；false代表content无更新
func (this *NoteImageService) ReOrganizeImageFiles(userId, noteId, title string, content *string, noContent bool) (res bool) {
	if title == "" {
		title = "empty-title-images"
	}
	title = FixFilename(title)

	var oldFullPath, newFullPath, newImagePath string
	basePath := ConfigS.GlobalStringConfigs["files.dir"]

	defer func() {
		if oldFullPath != "" {
			DeleteFile(filepath.Dir(oldFullPath))
		}
	}()
	moveNoteImage := func(i int, imageId string, file *info.File) (res bool) {
		// 创建移动路径
		fName := strings.Split(filepath.Base(file.Path), "_")
		newImagePath = filepath.Join(userId, "/images/", title, fmt.Sprintf("%d_%s", i, fName[len(fName)-1]))
		oldFullPath = path.Join(basePath, file.Path)
		newFullPath = path.Join(basePath, newImagePath)
		if err := os.MkdirAll(filepath.Dir(newFullPath), 0755); err != nil {
			return false
		}
		defer func() {
			if !res { // 失败时，仅空目录删除
				DeleteFile(filepath.Dir(newFullPath))
			}
		}()

		Logf("moveNoteImage(%s): %s -> %s", imageId, file.Path, newImagePath)
		if oldFullPath != newFullPath && MoveFile(oldFullPath, newFullPath) == nil {
			// 更新数据库
			file.Path = newImagePath
			if ok := db.UpdateByIdAndUserId(db.Files, imageId, userId, file); !ok {
				// 数据库写失败，回滚
				MoveFile(newFullPath, oldFullPath)
				return false
			}
		}

		return true
	}
	copyNoteImage := func(i int, imageId string, file *info.File) bool {
		oldPath := file.Path
		if ok, newID := fileService.CopyImageToTitle(userId, imageId, title, userId, file); ok {
			// 复制过img后，需要更新note正文
			*content = strings.ReplaceAll(*content, "Image?fileId="+imageId, "Image?fileId="+newID)
			Logf("copyNoteImage(%s->%s): %s -> %s", imageId, newID, oldPath, file.Path)
			return true
		}
		return false
	}

	if noContent { // 只需移动图片
		note_imageIDs := []info.NoteImage{}
		db.ListByQ(db.NoteImages, bson.M{"NoteId": bson.ObjectIdHex(noteId)}, &note_imageIDs)
		for i, image := range note_imageIDs {
			file := &info.File{}
			if db.GetByIdAndUserId(db.Files, image.ImageId.Hex(), userId, file); file.Path != "" {
				Logf("moveNoteImage, 因为只有note标题变了")
				moveNoteImage(i, image.ImageId.Hex(), file)
			}
		}

		return false
	}

	// 处理content更新的情况

	// 获取旧note下的所有imageIds
	note_imageIDs_tmp := []info.NoteImage{}
	db.ListByQ(db.NoteImages, bson.M{"NoteId": bson.ObjectIdHex(noteId)}, &note_imageIDs_tmp)
	note_imageIDs := make(map[string]bool)
	for i := range note_imageIDs_tmp {
		note_imageIDs[note_imageIDs_tmp[i].ImageId.Hex()] = true
	}

	// 获取新note下的所有imageIds
	find := noteImageReg.FindAllStringSubmatch(*content, -1) // 查找
	if find == nil || len(find) < 1 {
		return false
	}
	find = DeduplicateMatches(find)

	other_noteORalbum_imageId := func(image_id string, file *info.File) bool {
		if file.AlbumId != bson.ObjectIdHex(DEFAULT_ALBUM_ID) {
			return true // 来自其他相册集里的图片，需要复制img
		}
		noteIds := this.GetNoteIds(image_id)
		noteIdHex := bson.ObjectIdHex(noteId)
		for i := range noteIds {
			if noteIds[i] != noteIdHex {
				return true // 来自其他笔记的图片，需要复制img
			}
		}
		return false
	}

	for i, each := range find {
		if each != nil && len(each) == 3 {
			file := &info.File{}
			image_id := each[2]

			if db.GetByIdAndUserId(db.Files, image_id, userId, file); file.Path != "" {
				if _, ok := note_imageIDs[image_id]; ok { // 重命名、移动
					note_imageIDs[image_id] = false // 标记该图片已处理

					Logf("moveNoteImage, 因为新旧content都有用到此图片")
					moveNoteImage(i, image_id, file)
				} else { // 这里要区分，是不是其他note/album的image
					if other_noteORalbum_imageId(image_id, file) { // 复制，并更新content
						Logf("copyNoteImage, 因为此图片来自其他note/album")
						if copyNoteImage(i, image_id, file) {
							res = true
						}
					} else { // 移动
						Logf("moveNoteImage, 因为此图片不来自其他note/album")
						moveNoteImage(i, image_id, file)
					}
				}
			}
		}
	}

	// 不需要再在这里删除图片了，因为note图片的删除动作，只会发生在NoteContentHistoryService里
	return
	/*
		for image_id, needDelete := range note_imageIDs {
			if needDelete {
				noteImages := this.GetNoteIds(image_id)
				noteIdHex := bson.ObjectIdHex(noteId)
				for i := range noteImages {
					if noteImages[i] != noteIdHex {
						needDelete = false
						break // 其他笔记有用此图片，不删除
					}
				}

				if needDelete {
					// 判断当前笔记的历史里有没有用此图片，有用则不删除

					file := &info.File{}
					if db.GetByIdAndUserId(db.Files, image_id, userId, file); file.Path != "" {
						if db.DeleteByIdAndUserId(db.Files, image_id, userId) {
							DeleteFile(path.Join(basePath, file.Path))
						}
					}
				}
			}
		}

		return
	*/
}
