package controllers

import (
	"encoding/json"

	"github.com/wiselike/revel"
	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote-of-unofficial/app/info"
	//	. "github.com/wiselike/leanote-of-unofficial/app/lea"
	//	"io/ioutil"
)

type Notebook struct {
	BaseController
}

func (c Notebook) Index(notebook info.Notebook, i int, name string) revel.Result {
	return c.RenderJSON(notebook)
}

// 得到用户的所有笔记本
func (c Notebook) GetNotebooks() revel.Result {
	re := notebookService.GetNotebooks(c.GetUserId())
	return c.RenderJSON(re)
}

func (c Notebook) DeleteNotebook(notebookId string) revel.Result {
	re, msg := notebookService.DeleteNotebook(c.GetUserId(), notebookId)
	return c.RenderJSON(info.Re{Ok: re, Msg: msg})
}

// 添加notebook
func (c Notebook) AddNotebook(notebookId, title, parentNotebookId string) revel.Result {
	notebook := info.Notebook{NotebookId: bson.ObjectIdHex(notebookId),
		Title:  title,
		Seq:    -1,
		UserId: c.GetObjectUserId()}
	if parentNotebookId != "" {
		notebook.ParentNotebookId = bson.ObjectIdHex(parentNotebookId)
	}

	re, notebook := notebookService.AddNotebook(notebook)

	if re {
		return c.RenderJSON(notebook)
	} else {
		return c.RenderJSON(false)
	}
}

// 修改标题
func (c Notebook) UpdateNotebookTitle(notebookId, title string) revel.Result {
	return c.RenderJSON(notebookService.UpdateNotebookTitle(notebookId, c.GetUserId(), title))
}

// 修改笔记本排序和显示方式
func (c Notebook) UpdateNotebookAddition(notebookId, view, sortBy string) revel.Result {
	return c.RenderJSON(notebookService.UpdateNotebookAddition(notebookId, c.GetUserId(), view, sortBy))
}

// 排序
// 无用
// func (c Notebook) SortNotebooks(notebookId2Seqs map[string]int) revel.Result {
// 	return c.RenderJSON(notebookService.SortNotebooks(c.GetUserId(), notebookId2Seqs))
// }

// 传过来的data是JSON.stringfy数据
func (c Notebook) DragNotebooks(data string) revel.Result {
	info := info.DragNotebooksInfo{}
	json.Unmarshal([]byte(data), &info)

	return c.RenderJSON(notebookService.DragNotebooks(c.GetUserId(), info.CurNotebookId, info.ParentNotebookId, info.Siblings))
}

// 设置notebook <-> blog
func (c Notebook) SetNotebook2Blog(notebookId string, isBlog bool) revel.Result {
	re := notebookService.ToBlog(c.GetUserId(), notebookId, isBlog)
	return c.RenderJSON(re)
}
