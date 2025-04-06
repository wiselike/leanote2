package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/wiselike/revel"
	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote-of-unofficial/app/info"
	. "github.com/wiselike/leanote-of-unofficial/app/lea"
	"github.com/wiselike/leanote-of-unofficial/app/lea/netutil"
	"github.com/wiselike/leanote-of-unofficial/app/service"
)

// 首页
type File struct {
	BaseController
}

// 上传的是博客logo
// TODO logo不要设置权限, 另外的目录
func (c File) UploadBlogLogo() revel.Result {
	re := c.uploadImage("blogLogo", "")

	c.ViewArgs["fileUrlPath"] = re.Id
	c.ViewArgs["resultCode"] = re.Code
	c.ViewArgs["resultMsg"] = re.Msg

	return c.RenderTemplate("file/blog_logo.html")
}

// 拖拉上传, pasteImage
// noteId 是为了判断是否是协作的note, 如果是则需要复制一份到note owner中
func (c File) PasteImage(noteId string) revel.Result {
	re := c.uploadImage("pasteImage", "")

	if noteId != "" {
		userId := c.GetUserId()
		note := noteService.GetNoteById(noteId)
		if note.UserId != "" {
			noteUserId := note.UserId.Hex()
			if noteUserId != userId {
				// 是否是有权限协作的
				if shareService.HasUpdatePerm(noteUserId, userId, noteId) {
					// 复制图片之, 图片复制给noteUserId
					_, re.Id = fileService.CopyImage(userId, re.Id, noteUserId)
				} else {
					// 怎么可能在这个笔记下paste图片呢?
					// 正常情况下不会
				}
			}
		}
	}

	return c.RenderJSON(re)
}

// 头像设置
func (c File) UploadAvatar() revel.Result {
	re := c.uploadImage("logo", "")

	c.ViewArgs["fileUrlPath"] = re.Id
	c.ViewArgs["resultCode"] = re.Code
	c.ViewArgs["resultMsg"] = re.Msg

	if re.Ok {
		re.Ok = userService.UpdateAvatar(c.GetUserId(), re.Id)
		if re.Ok {
			c.UpdateSession("Logo", re.Id)
		}
	}

	return c.RenderJSON(re)
}

// leaui image plugin upload image
func (c File) UploadImageLeaui(albumId string) revel.Result {
	re := c.uploadImage("", albumId)
	return c.RenderJSON(re)
}

// 上传图片, 公用方法
// upload image common func
func (c File) uploadImage(from, albumId string) (re info.Re) {
	var fileUrlPath, dir string
	var fileId string
	var resultCode = 0      // 1表示正常
	var resultMsg = "error" // 错误信息
	var Ok = false

	defer func() {
		re.Id = fileId // 只是id, 没有其它信息
		re.Code = resultCode
		re.Msg = resultMsg
		re.Ok = Ok
	}()

	var data []byte
	c.Params.Bind(&data, "file")
	handel := c.Params.Files["file"][0]
	if data == nil || len(data) == 0 {
		return re
	}

	// file, handel, err := c.Request.FormFile("file")
	// if err != nil {
	// 	return re
	// }
	// defer file.Close()

	// data, err := ioutil.ReadAll(file)
	newGuid := NewGuid()

	userId := c.GetUserId()

	if from == "logo" || from == "blogLogo" { // 上传logo则放到logo目录
		fileUrlPath = "public/upload/" + Digest3(userId) + "/" + userId + "/images/logo"
		dir = path.Join(revel.BasePath, fileUrlPath)
	} else {
		album := albumService.GetAlbumById(c.GetUserId(), albumId)
		if album.Name == "" { // 相册名为空，则是默认相册，并且是文章中的图片
			fileUrlPath = GetRandomFilePath(userId, newGuid) + "/images-tmp"
		} else { // 相册名不为空的，上传到对应相册文件夹
			fileUrlPath = GetRandomFilePath(userId, newGuid) + "/albums/" + album.Name
		}
		dir = path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], fileUrlPath)
	}

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return re
	}

	// 生成新的随机文件名
	var ext string
	if from == "pasteImage" {
		handel.Filename = c.Message("unTitled")
		ext = ".png" // TODO 可能不是png类型
	} else {
		_, ext = SplitFilename(handel.Filename)
		if ext != ".gif" && ext != ".jpg" && ext != ".png" && ext != ".bmp" && ext != ".jpeg" {
			resultMsg = "Please upload image"
			return re
		}
	}

	filename := newGuid + ext

	var maxFileSize float64
	if from == "logo" {
		maxFileSize = configService.GetUploadSize("uploadAvatarSize")
	} else if from == "blogLogo" {
		maxFileSize = configService.GetUploadSize("uploadBlogLogoSize")
	} else {
		maxFileSize = configService.GetUploadSize("uploadImageSize")
	}
	if maxFileSize <= 0 {
		maxFileSize = 1000
	}

	// > 2M?
	if float64(len(data)) > maxFileSize*float64(1024*1024) {
		resultCode = 0
		resultMsg = fmt.Sprintf("The file Size is bigger than %vM", maxFileSize)
		return re
	}

	toPath := path.Join(dir, filename)
	err = ioutil.WriteFile(toPath, data, 0777)
	if err != nil {
		LogJ(err)
		return re
	}
	// 备份原始图片
	_, toPathGif := TransPicture(toPath, path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], "backup-origins", c.GetUserId(), time.Now().Format("2006")))
	filename = GetFilename(toPathGif)
	filesize := GetFilesize(toPathGif)
	fileUrlPath += "/" + filename
	resultCode = 1
	resultMsg = "Upload Success!"

	// File
	fileInfo := info.File{Name: filename,
		Title: handel.Filename,
		Path:  fileUrlPath,
		Size:  filesize}

	id := bson.NewObjectId()
	fileInfo.FileId = id
	fileId = id.Hex()

	if from == "logo" || from == "blogLogo" {
		fileId = fileUrlPath
	}

	Ok, resultMsg = fileService.AddImage(fileInfo, albumId, c.GetUserId(), from == "" || from == "pasteImage")
	resultMsg = c.Message(resultMsg)
	if !Ok { // 若数据库插入image失败，本地image也没必要保存
		DeleteFile(toPath)
	}

	fileInfo.Path = "" // 不要返回
	re.Item = fileInfo

	return re
}

// get all images by userId with page
func (c File) GetImages(albumId, key string, page int) revel.Result {
	re := fileService.ListImagesWithPage(c.GetUserId(), albumId, key, page, 12)
	return c.RenderJSON(re)
}

func (c File) UpdateImageTitle(fileId, title string) revel.Result {
	re := info.NewRe()
	re.Ok = fileService.UpdateImageTitle(c.GetUserId(), fileId, title)
	return c.RenderJSON(re)
}

func (c File) DeleteImage(fileId string) revel.Result {
	re := info.NewRe()
	re.Ok, re.Msg = fileService.DeleteImage(c.GetUserId(), fileId)
	return c.RenderJSON(re)
}

//-----------

// 输出image
// 权限判断
func (c File) OutputImage(noteId, fileId string) revel.Result {
	fpath := fileService.GetFile(c.GetUserId(), fileId) // 得到路径
	if fpath == "" {
		return c.RenderText("")
	}
	fn := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], fpath)
	file, _ := os.Open(fn)
	return c.RenderFile(file, revel.Inline) // revel.Attachment
}

// 协作时复制图片到owner
// 需要计算对方大小
func (c File) CopyImage(userId, fileId, toUserId string) revel.Result {
	re := info.NewRe()
	re.Ok, re.Id = fileService.CopyImage(userId, fileId, toUserId)
	return c.RenderJSON(re)
}

// 复制外网的图片
// 都要好好的计算大小
func (c File) CopyHttpImage(src string) revel.Result {
	re := info.NewRe()

	// 生成上传路径
	// newGuid := NewGuid()
	userId := c.GetUserId()
	fileUrlPath := GetRandomFilePath(userId, "") + "/images"
	dir := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], fileUrlPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return c.RenderJSON(re)
	}
	filesize, filename, _, ok := netutil.WriteUrl(src, dir)

	if !ok {
		re.Msg = "copy error"
		return c.RenderJSON(re)
	}

	// File
	fileInfo := info.File{Name: filename,
		Title: filename,
		Path:  fileUrlPath + "/" + filename,
		Size:  filesize}

	id := bson.NewObjectId()
	fileInfo.FileId = id

	re.Id = id.Hex()
	//	re.Item = fileInfo.Path
	re.Ok, re.Msg = fileService.AddImage(fileInfo, "", c.GetUserId(), true)

	return c.RenderJSON(re)
}
