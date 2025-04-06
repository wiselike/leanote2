package api

import (
	"github.com/wiselike/revel"

	"github.com/wiselike/leanote-of-unofficial/app/service"

	//	"encoding/json"
	//	. "github.com/wiselike/leanote-of-unofficial/app/lea"
	//	"gopkg.in/mgo.v2/bson"
	//	"github.com/wiselike/leanote-of-unofficial/app/lea/netutil"
	//	"github.com/wiselike/leanote-of-unofficial/app/info"
	//	"io/ioutil"
	"os"
	"path"

	//	"strconv"
	"archive/tar"
	"compress/gzip"
	"io"
	"time"
)

// 文件操作, 图片, 头像上传, 输出

type ApiFile struct {
	ApiBaseContrller
}

/*
// 协作时复制图片到owner
func (c ApiFile) CopyImage(userId, fileId, toUserId string) revel.Result {
	re := info.NewRe()

	re.Ok, re.Id = fileService.CopyImage(userId, fileId, toUserId)

	return c.RenderJSON(re)
}

// get all images by userId with page
func (c ApiFile) GetImages(albumId, key string, page int) revel.Result {
	imagesPage := fileService.ListImagesWithPage(c.getUserId(), albumId, key, page, 12)
	re := info.NewRe()
	re.Ok = true
	re.Item = imagesPage
	return c.RenderJSON(re)
}

func (c ApiFile) UpdateImageTitle(fileId, title string) revel.Result {
	re := info.NewRe()
	re.Ok = fileService.UpdateImageTitle(c.getUserId(), fileId, title)
	return c.RenderJSON(re)
}

func (c ApiFile) DeleteImage(fileId string) revel.Result {
	re := info.NewRe()
	re.Ok, re.Msg = fileService.DeleteImage(c.getUserId(), fileId)
	return c.RenderJSON(re)
}

*/

//-----------

// 输出image
// [OK]
func (c ApiFile) GetImage(fileId string) revel.Result {
	fpath := fileService.GetFile(c.getUserId(), fileId) // 得到路径
	if fpath == "" {
		return c.RenderText("")
	}
	fn := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], fpath)
	file, _ := os.Open(fn)
	return c.RenderFile(file, revel.Inline) // revel.Attachment
}

// 下载附件
// [OK]
func (c ApiFile) GetAttach(fileId string) revel.Result {
	attach := attachService.GetAttach(fileId, c.getUserId()) // 得到路径
	fpath := attach.Path
	if fpath == "" {
		return c.RenderText("No Such File")
	}
	fn := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], fpath)
	file, _ := os.Open(fn)
	return c.RenderBinary(file, attach.Title, revel.Attachment, time.Now()) // revel.Attachment
}

// 下载所有附件
// [OK]
func (c ApiFile) GetAllAttachs(noteId string) revel.Result {
	note := noteService.GetNoteById(noteId)
	if note.NoteId == "" {
		return c.RenderText("")
	}
	// 得到文件列表
	attachs := attachService.ListAttachs(noteId, c.getUserId())
	if attachs == nil || len(attachs) == 0 {
		return c.RenderText("")
	}

	/*
		dir := path.Join(service.ConfigS.GlobalStringConfigs["files.dir"], "/files/tmp")
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return c.RenderText("")
		}
	*/

	filename := note.Title + ".tar.gz"
	if note.Title == "" {
		filename = "all.tar.gz"
	}

	basePath := service.ConfigS.GlobalStringConfigs["files.dir"]
	// file write
	fw, err := os.Create(path.Join(basePath, filename))
	if err != nil {
		return c.RenderText("")
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
	}

	// fw.Seek(0, 0) // 不需要seek
	return c.RenderBinary(fw, filename, revel.Attachment, time.Now()) // revel.Attachment
}
