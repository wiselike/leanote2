package controllers

import (
	"math/rand"
	"strings"
	"time"

	"github.com/wiselike/revel"

	"github.com/wiselike/leanote2/app/info"
	. "github.com/wiselike/leanote2/app/lea"
)

// 用户登录/注销/找回密码

type Auth struct {
	BaseController
}

// --------
// 登录
func (c Auth) Login(email, from string) revel.Result {
	c.ViewArgs["title"] = c.Message("login")
	c.ViewArgs["subTitle"] = c.Message("login")
	c.ViewArgs["email"] = email
	c.ViewArgs["from"] = from
	c.ViewArgs["openRegister"] = configService.IsOpenRegister()

	sessionId := c.Session.ID()
	defer sessionService.Update(sessionId, "LastClientIP", c.ClientIP)
	if sessionService.LoginTimesIsOver(sessionId) {
		c.ViewArgs["needCaptcha"] = true
	}

	c.SetLocale()

	if c.Has("demo") {
		c.ViewArgs["demo"] = true
		c.ViewArgs["email"] = "demo@leanote.com"
	}
	return c.RenderTemplate("home/login.html")
}

// 为了demo和register
func (c Auth) doLogin(email, pwd string) revel.Result {
	sessionId := c.Session.ID()
	defer sessionService.Update(sessionId, "LastClientIP", c.ClientIP)
	var msg = "wrongUsernameOrPassword"

	// 没有客户端IP就不用登陆了
	if c.ClientIP != "" {
		userInfo, err := authService.Login(email, pwd)
		if err == nil {
			c.SetSession(userInfo)
			sessionService.ClearLoginTimes(sessionId)
			// 记录登陆成功的用户
			sessionService.SetUserId(sessionId, userInfo.UserId.Hex())
			return c.RenderJSON(info.Re{Ok: true})
		}
	}

	revel.AppLog.Warnf("username(%s) or password(%s) or ip(%s) is incorrect.", email, pwd, c.ClientIP)
	return c.RenderJSON(info.Re{Ok: false, Item: sessionService.LoginTimesIsOver(sessionId), Msg: c.Message(msg)})
}
func (c Auth) DoLogin(email, pwd string, captcha string) revel.Result {
	sessionId := c.Session.ID()
	defer sessionService.Update(sessionId, "LastClientIP", c.ClientIP)
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var msg = ""

	// >= 3次需要验证码, 直到登录成功
	if sessionService.LoginTimesIsOver(sessionId) && sessionService.GetCaptcha(sessionId) != captcha {
		msg = "captchaError"
	} else {
		// 没有客户端IP就不用登陆了
		if c.ClientIP != "" {
			userInfo, err := authService.Login(email, pwd)
			if err == nil {
				c.SetSession(userInfo)
				sessionService.ClearLoginTimes(sessionId)
				// 记录登陆成功的用户
				sessionService.SetUserId(sessionId, userInfo.UserId.Hex())
				return c.RenderJSON(info.Re{Ok: true})
			}
		}

		// 登录错误, 则错误次数++
		msg = "wrongUsernameOrPassword"
		revel.AppLog.Warnf("username(%s) or password(%s) or ip(%s) is incorrect.", email, pwd, c.ClientIP)
	}
	sessionService.IncrLoginTimes(sessionId)
	// 这部分代码是冗余的，防止前端没有请求新验证码(密码猜测攻击)，这里强制设置一个临时的不可能的fake captcha
	// 但仍然不能够阻止/api/auth方式的刷密码攻击
	token := make([]byte, len(letterBytes))
	for i := range token {
		token[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	sessionService.SetCaptcha(sessionId, Md5(letterBytes+string(token)+sessionId))

	if sessionService.LoginTimesIsOver(sessionId) {
		// 重试太多次了，休息一下
		time.Sleep(time.Second * 2)
	}
	return c.RenderJSON(info.Re{Ok: false, Item: sessionService.LoginTimesIsOver(sessionId), Msg: c.Message(msg)})
}

// 注销
func (c Auth) Logout() revel.Result {
	sessionId := c.Session.ID()
	sessionService.Clear(sessionId)
	c.ClearSession()
	return c.Redirect("/login")
}

// 体验一下
func (c Auth) Demo() revel.Result {
	email := "forbidden" // configService.GetGlobalStringConfig("demoUsername")
	pwd := "forbidden"   // configService.GetGlobalStringConfig("demoPassword")

	userInfo, err := authService.Login(email, pwd)
	if err != nil {
		return c.RenderJSON(info.Re{Ok: false, Msg: "Demo is forbidden now!"})
	} else {
		c.SetSession(userInfo)
		return c.Redirect("/note")
	}
	return nil
}

// --------
// 注册
func (c Auth) Register(from, iu string) revel.Result {
	if !configService.IsOpenRegister() {
		return c.Redirect("/index")
	}
	c.SetLocale()
	c.ViewArgs["from"] = from
	c.ViewArgs["iu"] = iu

	c.ViewArgs["title"] = c.Message("register")
	c.ViewArgs["subTitle"] = c.Message("register")
	return c.RenderTemplate("home/register.html")
}
func (c Auth) DoRegister(email, pwd, iu string) revel.Result {
	if !configService.IsOpenRegister() {
		return c.Redirect("/index")
	}

	re := info.NewRe()

	if re.Ok, re.Msg = Vd("email", email); !re.Ok {
		return c.RenderRe(re)
	}
	if re.Ok, re.Msg = Vd("password", pwd); !re.Ok {
		return c.RenderRe(re)
	}

	email = strings.ToLower(email)

	// 注册
	re.Ok, re.Msg = authService.Register(email, pwd, iu)

	// 注册成功, 则立即登录之
	if re.Ok {
		c.doLogin(email, pwd)
	}

	return c.RenderRe(re)
}

// --------
// 找回密码
func (c Auth) FindPassword() revel.Result {
	c.SetLocale()
	c.ViewArgs["title"] = c.Message("findPassword")
	c.ViewArgs["subTitle"] = c.Message("findPassword")
	return c.RenderTemplate("home/find_password.html")
}
func (c Auth) DoFindPassword(email string) revel.Result {
	re := info.NewRe()
	re.Ok, re.Msg = pwdService.FindPwd(email)
	return c.RenderJSON(re)
}

// 点击链接后, 先验证之
func (c Auth) FindPassword2(token string) revel.Result {
	c.SetLocale()
	c.ViewArgs["title"] = c.Message("findPassword")
	c.ViewArgs["subTitle"] = c.Message("findPassword")
	if token == "" {
		return c.RenderTemplate("find_password2_timeout.html")
	}
	ok, _, findPwd := tokenService.VerifyToken(token, info.TokenPwd)
	if !ok {
		return c.RenderTemplate("home/find_password2_timeout.html")
	}
	c.ViewArgs["findPwd"] = findPwd

	c.ViewArgs["title"] = c.Message("updatePassword")
	c.ViewArgs["subTitle"] = c.Message("updatePassword")

	return c.RenderTemplate("home/find_password2.html")
}

// 找回密码修改密码
func (c Auth) FindPasswordUpdate(token, pwd string) revel.Result {
	re := info.NewRe()

	if re.Ok, re.Msg = Vd("password", pwd); !re.Ok {
		return c.RenderRe(re)
	}

	// 修改之
	re.Ok, re.Msg = pwdService.UpdatePwd(token, pwd)
	return c.RenderRe(re)
}
