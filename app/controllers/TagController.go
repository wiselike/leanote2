package controllers

import (
	"github.com/wiselike/revel"
	//	"encoding/json"
	//	"gopkg.in/mgo.v2/bson"
	//	. "github.com/wiselike/leanote2/app/lea"
	"github.com/wiselike/leanote2/app/info"
	//	"os/exec"
)

type Tag struct {
	BaseController
}

// 更新Tag
func (c Tag) UpdateTag(tag string) revel.Result {
	ret := info.NewRe()
	ret.Ok = true
	ret.Item = tagService.AddOrUpdateTag(c.GetUserId(), tag)
	return c.RenderJSON(ret)
}

// 删除标签
func (c Tag) DeleteTag(tag string) revel.Result {
	ret := info.Re{}
	ret.Ok = true
	ret.Item = tagService.DeleteTag(c.GetUserId(), tag)
	return c.RenderJSON(ret)
}
