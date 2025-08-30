package service

import (
	"github.com/wiselike/leanote2/app/db"
	"github.com/wiselike/leanote2/app/info"

	// . "github.com/wiselike/leanote2/app/lea"
	"time"

	"gopkg.in/mgo.v2/bson"
	//	"strings"
)

// Session存储到mongodb中
type SessionService struct {
}

func (this *SessionService) Update(sessionId, key string, value interface{}) bool {
	return db.UpdateByQMap(db.Sessions, bson.M{"SessionId": sessionId},
		bson.M{key: value, "UpdatedTime": time.Now()})
}

// 注销时清空session
func (this *SessionService) Clear(sessionId string) bool {
	// return db.Delete(db.Sessions, bson.M{"SessionId": sessionId})
	return this.Update(sessionId, "UserId", "") // 仅清理userID，保留追溯记录
}
func (this *SessionService) Get(sessionId string) info.Session {
	session := info.Session{}
	db.GetByQ(db.Sessions, bson.M{"SessionId": sessionId}, &session)

	// 如果session已过期，就删掉让用户重新登录
	if !session.UpdatedTime.IsZero() && time.Since(session.UpdatedTime) > configService.GetSessionExpireTime() {
		db.Delete(db.Sessions, bson.M{"SessionId": sessionId})
		session = info.Session{}
	}
	// 如果没有session, 那么插入一条之
	if session.Id == "" {
		session.Id = bson.NewObjectId()
		session.SessionId = sessionId
		session.CreatedTime = time.Now()
		session.LoginTimes = 1
		session.UpdatedTime = session.CreatedTime
		db.Insert(db.Sessions, session)
	}

	return session
}

//------------------
// 错误次数处理

// 登录错误时间是否已超过了
func (this *SessionService) LoginTimesIsOver(sessionId string) bool {
	session := this.Get(sessionId)
	// >= 3次需要验证码, 直到登录成功
	return session.LoginTimes > 3
}

// 登录成功后清空错误次数
func (this *SessionService) ClearLoginTimes(sessionId string) bool {
	return this.Update(sessionId, "LoginTimes", 0)
}

// 增加错误次数
func (this *SessionService) IncrLoginTimes(sessionId string) bool {
	session := this.Get(sessionId)
	return this.Update(sessionId, "LoginTimes", session.LoginTimes+1)
}

// ----------
// 验证码
func (this *SessionService) GetCaptcha(sessionId string) string {
	session := this.Get(sessionId)
	return session.Captcha
}
func (this *SessionService) SetCaptcha(sessionId, captcha string) bool {
	session := this.Get(sessionId)
	if session.UserId != "" && len(captcha) > 20 { // 已经登录就不用更新fake captcha了
		return true
	}
	return this.Update(sessionId, "Captcha", captcha)
}

// -----------
// API
func (this *SessionService) GetUserId(sessionId string) string {
	session := this.Get(sessionId)
	// 更新updateTime, 避免过期
	db.UpdateByQMap(db.Sessions, bson.M{"SessionId": sessionId},
		bson.M{"UpdatedTime": time.Now()})
	return session.UserId
}

// 登录成功后设置userId
func (this *SessionService) SetUserId(sessionId, userId string) bool {
	this.Get(sessionId)
	ok := this.Update(sessionId, "UserId", userId)

	// 清理未登录成功的、已过期的session
	selector := bson.M{
		"UserId":      bson.M{"$in": []interface{}{"", nil}},
		"UpdatedTime": bson.M{"$lt": time.Now().Add(-185 * 24 * time.Hour)}, // 留存六个月
	}
	db.DeleteAll(db.Sessions, selector)
	return ok
}
