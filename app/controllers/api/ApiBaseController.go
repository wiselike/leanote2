package api

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote2/app/controllers"
	"github.com/wiselike/leanote2/app/info"
	. "github.com/wiselike/leanote2/app/lea"
	"github.com/wiselike/leanote2/app/service"
)

// 公用Controller, 其它Controller继承它
type ApiBaseContrller struct {
	controllers.BaseController // 不能用*BaseController
}

// 得到token, 这个token是在AuthInterceptor设到Session中的
func (c ApiBaseContrller) getToken() string {
	return c.GetSession("_token")
}

// userId
// _userId是在AuthInterceptor设置的
func (c ApiBaseContrller) getUserId() string {
	return c.GetSession("_userId")
}

// 得到用户信息
func (c ApiBaseContrller) getUserInfo() info.User {
	userId := c.GetSession("_userId")
	if userId == "" {
		return info.User{}
	}
	return userService.GetUserInfo(userId)
}

// 上传附件
func (c ApiBaseContrller) uploadAttach(name string, noteId string) (ok bool, msg string, id string) {
	userId := c.getUserId()

	// 判断是否有权限为笔记添加附件
	// 如果笔记还没有添加是不是会有问题
	/*
		if !shareService.HasUpdateNotePerm(noteId, userId) {
			return
		}
	*/

	var data []byte
	c.Params.Bind(&data, name)
	handel := c.Params.Files[name][0]
	if data == nil || len(data) == 0 {
		return
	}

	// file, handel, err := c.Request.FormFile(name)
	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	return
	// }
	// > 5M?
	maxFileSize := configService.GetUploadSize("uploadAttachSize")
	if maxFileSize <= 0 {
		maxFileSize = 1000
	}
	if float64(len(data)) > maxFileSize*float64(1024*1024) {
		msg = "fileIsTooLarge"
		return
	}

	// 生成上传路径
	newGuid := NewGuid()
	filePath := GetRandomFilePath(userId, newGuid) + "/attachs"

	dir := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}
	// 生成新的文件名
	filename := handel.Filename
	_, ext := SplitFilename(filename) // .doc
	filename = newGuid + ext
	toPath := path.Join(dir, filename)
	err = ioutil.WriteFile(toPath, data, 0777)
	if err != nil {
		return
	}

	// add File to db
	fileType := ""
	if ext != "" {
		fileType = strings.ToLower(ext[1:])
	}
	filesize := GetFilesize(toPath)
	fileInfo := info.Attach{AttachId: bson.NewObjectId(),
		Name:         filename,
		Title:        handel.Filename,
		NoteId:       bson.ObjectIdHex(noteId),
		UploadUserId: bson.ObjectIdHex(userId),
		Path:         path.Join(filePath, filename),
		Type:         fileType,
		Size:         filesize}

	ok, msg = attachService.AddAttach(fileInfo, true)
	if !ok {
		return
	}

	id = fileInfo.AttachId.Hex()
	return
}

// 上传附件或图片
func (c ApiBaseContrller) upload(name string, noteId string, isAttach bool) (ok bool, msg string, id string) {
	if isAttach {
		return c.uploadAttach(name, noteId)
	}
	// file, handel, err := c.Request.FormFile(name)
	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	var data []byte
	c.Params.Bind(&data, name)
	handel := c.Params.Files[name][0]
	if data == nil || len(data) == 0 {
		return
	}

	newGuid := NewGuid()
	// 生成上传路径
	userId := c.getUserId()
	fileUrlPath := GetRandomFilePath(userId, newGuid) + "/images"

	dir := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], fileUrlPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}
	// 生成新的文件名
	filename := handel.Filename
	_, ext := SplitFilename(filename)
	switch ext {
	case "", ".gif", ".jpg", ".png", ".bmp", ".jpeg":
	default:
		msg = "notImage"
		return
	}

	filename = newGuid + ext
	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	return
	// }

	maxFileSize := configService.GetUploadSize("uploadImageSize")
	if maxFileSize <= 0 {
		maxFileSize = 1000
	}

	// > 2M?
	if float64(len(data)) > maxFileSize*float64(1024*1024) {
		msg = "fileIsTooLarge"
		return
	}

	toPath := path.Join(dir, filename)
	err = ioutil.WriteFile(toPath, data, 0777)
	if err != nil {
		return
	}
	// 改变成gif图片
	_, toPathGif := TransPicture(toPath, path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], "backup-origins", c.GetUserId(), time.Now().Format("2006")))
	filename = GetFilename(toPathGif)
	filesize := GetFilesize(toPathGif)

	// File
	fileInfo := info.File{FileId: bson.NewObjectId(),
		Name:  filename,
		Title: handel.Filename,
		Path:  path.Join(fileUrlPath, filename),
		Size:  filesize}
	ok, msg = fileService.AddImage(fileInfo, "", c.getUserId(), true)
	if ok {
		id = fileInfo.FileId.Hex()
	}
	return
}
