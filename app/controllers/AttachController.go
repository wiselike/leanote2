package controllers

import (
	"github.com/wiselike/revel"
	//	"encoding/json"
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote2/app/info"
	. "github.com/wiselike/leanote2/app/lea"
	"github.com/wiselike/leanote2/app/service"
)

// 附件
type Attach struct {
	BaseController
}

// 上传附件
func (c Attach) UploadAttach(noteId string) revel.Result {
	re := c.uploadAttach(noteId)
	return c.RenderJSON(re)
}
func (c Attach) uploadAttach(noteId string) (re info.Re) {
	var fileId = ""
	var resultMsg = "error" // 错误信息
	var Ok = false
	var fileInfo info.Attach

	re = info.NewRe()

	defer func() {
		re.Id = fileId // 只是id, 没有其它信息
		re.Msg = resultMsg
		re.Ok = Ok
		re.Item = fileInfo
	}()

	// 判断是否有权限为笔记添加附件
	if !shareService.HasUpdateNotePerm(noteId, c.GetUserId()) {
		return re
	}

	var data []byte
	c.Params.Bind(&data, "file")

	if data == nil || len(data) == 0 {
		return re
	}
	// > 5M?
	maxFileSize := configService.GetUploadSize("uploadAttachSize")
	if maxFileSize <= 0 {
		maxFileSize = 1000
	}
	if float64(len(data)) > maxFileSize*float64(1024*1024) {
		resultMsg = fmt.Sprintf("The file's size is bigger than %vM", maxFileSize)
		return re
	}

	// 生成上传路径
	title := noteService.GetNote(noteId, c.GetUserId()).Title
	title = FixFilename(title)
	filePath := path.Join(c.GetUserId(), "attachs", title)
	dir := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return re
	}

	handel := c.Params.Files["file"][0]

	// 生成新的文件名
	filename := handel.Filename
	_, ext := SplitFilename(filename) // .doc
	filename = NewGuid() + ext
	toPath := dir + "/" + filename
	err = ioutil.WriteFile(toPath, data, 0777)
	if err != nil {
		return re
	}

	// add File to db
	fileType := ""
	if ext != "" {
		fileType = strings.ToLower(ext[1:])
	}
	filesize := GetFilesize(toPath)
	fileInfo = info.Attach{Name: filename,
		Title:        handel.Filename,
		NoteId:       bson.ObjectIdHex(noteId),
		UploadUserId: c.GetObjectUserId(),
		Path:         filePath + "/" + filename,
		Type:         fileType,
		Size:         filesize}

	id := bson.NewObjectId()
	fileInfo.AttachId = id
	fileId = id.Hex()
	Ok, resultMsg = attachService.AddAttach(fileInfo, false)
	if resultMsg != "" {
		resultMsg = c.Message(resultMsg)
	}

	fileInfo.Path = "" // 不要返回
	if Ok {
		resultMsg = "success"
	}
	return re
}

// 删除附件
func (c Attach) DeleteAttach(attachId string) revel.Result {
	re := info.NewRe()
	re.Ok, re.Msg = attachService.DeleteAttach(attachId, c.GetUserId())
	return c.RenderJSON(re)
}

// get all attachs by noteId
func (c Attach) GetAttachs(noteId string) revel.Result {
	re := info.NewRe()
	re.Ok = true
	re.List = attachService.ListAttachs(noteId, c.GetUserId())
	return c.RenderJSON(re)
}

// 下载附件
// 权限判断
func (c Attach) Download(attachId string) revel.Result {
	attach := attachService.GetAttach(attachId, c.GetUserId()) // 得到路径
	fpath := attach.Path
	if fpath == "" {
		return c.RenderText("")
	}
	fn := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], fpath)
	file, _ := os.Open(fn)
	return c.RenderBinary(file, attach.Title, revel.Attachment, time.Now()) // revel.Attachment
}

func (c Attach) DownloadAll(noteId string) revel.Result {
	note := noteService.GetNoteById(noteId)
	if note.NoteId == "" {
		return c.RenderText("")
	}
	// 得到文件列表
	attachs := attachService.ListAttachs(noteId, c.GetUserId())
	if attachs == nil || len(attachs) == 0 {
		return c.RenderText("")
	}

	/*
		dir := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], "/tmp")
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return c.RenderText("")
		}
	*/

	var filename string
	if note.Title == "" {
		filename = "all.tar.gz"
	} else {
		filename = FixFilename(note.Title) + ".tar.gz"
	}

	basePath := service.ConfigS.GlobalStringConfigs["files.dir"]
	dir := path.Join(basePath, GetRandomFilePath(c.GetUserId(), ""), "attach_all", time.Now().Format("2006"))
	if !MkdirAll(dir) {
		return c.RenderText("error")
	}

	// file write
	fw, err := os.Create(path.Join(dir, filename))
	if err != nil {
		return c.RenderText("error")
	}
	// defer fw.Close() // 不需要关闭, 还要读取给用户下载

	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// 遍历文件列表
	for _, attach := range attachs {
		fn := path.Join(basePath, attach.Path)
		fr, err := os.Open(fn)
		fileInfo, _ := fr.Stat()
		if err != nil {
			return c.RenderText("")
		}

		// 信息头
		h := new(tar.Header)
		h.Name = attach.Title
		h.Size = fileInfo.Size()
		h.Mode = int64(fileInfo.Mode())
		h.ModTime = fileInfo.ModTime()

		// 写信息头
		err = tw.WriteHeader(h)
		if err != nil {
			panic(err)
		}

		// 写文件
		_, err = io.Copy(tw, fr)
		if err != nil {
			panic(err)
		}

		fr.Close()
	} // for

	// fw.Seek(0, 0)
	return c.RenderBinary(fw, filename, revel.Attachment, time.Now()) // revel.Attachment
}
