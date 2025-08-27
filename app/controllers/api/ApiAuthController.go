package api

import (
	"time"

	"github.com/wiselike/revel"
	"gopkg.in/mgo.v2/bson"

	"github.com/wiselike/leanote-of-unofficial/app/info"
	. "github.com/wiselike/leanote-of-unofficial/app/lea"
)

// 用户登录后生成一个token, 将这个token保存到session中
// 以后每次的请求必须带这个token, 并从session中获取userId

// 用户登录/注销/找回密码

type ApiAuth struct {
	ApiBaseContrller
}

// 登录
// [ok]
// 成功返回 {Ok: true, Item: token }
// 失败返回 {Ok: false, Msg: ""}
func (c ApiAuth) Login(email, pwd string) revel.Result {
	var msg = "wrongUsernameOrPassword"

	// 没有客户端IP就不用登陆了
	if c.ClientIP != "" {
		userInfo, err := authService.Login(email, pwd)
		if err == nil {
			token := bson.NewObjectId().Hex()
			sessionService.SetUserId(token, userInfo.UserId.Hex())
			sessionService.Update(token, "LastClientIP", c.ClientIP)
			return c.RenderJSON(info.AuthOk{Ok: true, Token: token, UserId: userInfo.UserId, Email: userInfo.Email, Username: userInfo.Username})
		}
	}

	revel.AppLog.Warnf("username(%s) or password(%s) or ip(%s) is incorrect.", email, pwd, c.ClientIP)
	time.Sleep(time.Second * 2) // 登录错误就休息一下，缓一缓
	return c.RenderJSON(info.ApiRe{Ok: false, Msg: c.Message(msg)})
}

// 注销
// [Ok]
func (c ApiAuth) Logout() revel.Result {
	token := c.getToken()
	sessionService.Clear(token)
	re := info.NewApiRe()
	re.Ok = true
	return c.RenderJSON(re)
}

// 注册
// [Ok]
// 成功后并不返回用户ID, 需要用户重新登录
func (c ApiAuth) Register(email, pwd string) revel.Result {
	re := info.NewApiRe()
	if !configService.IsOpenRegister() {
		re.Msg = "notOpenRegister" // 未开放注册
		return c.RenderJSON(re)
	}

	if re.Ok, re.Msg = Vd("email", email); !re.Ok {
		return c.RenderJSON(re)
	}
	if re.Ok, re.Msg = Vd("password", pwd); !re.Ok {
		return c.RenderJSON(re)
	}

	// 注册
	re.Ok, re.Msg = authService.Register(email, pwd, "")
	return c.RenderJSON(re)
}
