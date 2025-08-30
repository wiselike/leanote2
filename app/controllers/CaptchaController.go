package controllers

import (
	"io"
	"net/http"

	"github.com/wiselike/revel"

	"github.com/wiselike/leanote2/app/lea/captcha"
)

// 验证码服务
type Captcha struct {
	BaseController
}

type Ca string

func (r Ca) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "image/png")
}

func (c Captcha) Get() revel.Result {
	c.Response.ContentType = "image/png"
	image, str := captcha.Fetch()
	out := io.Writer(c.Response.GetWriter())
	image.WriteTo(out)

	sessionId := c.GetSession("_ID")
	sessionService.SetCaptcha(sessionId, str)

	return c.Render()
}
