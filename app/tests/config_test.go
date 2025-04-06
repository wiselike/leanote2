package tests

import (
	"testing"

	"github.com/wiselike/revel"

	"github.com/wiselike/leanote-of-unofficial/app/db"

	//	. "github.com/wiselike/leanote-of-unofficial/app/lea"
	"github.com/wiselike/leanote-of-unofficial/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func init() {
	revel.Init("dev", "github.com/wiselike/leanote-of-unofficial", "/Users/life/Documents/Go/package_base/src")
	db.Init("mongodb://localhost:27017/leanote", "leanote")
	service.InitService()
	service.ConfigS.InitGlobalConfigs()
}

// 测试登录
func TestSendMail(t *testing.T) {
	ok, err := service.EmailS.SendEmail("life@leanote.com", "你好", "你好吗")
	t.Log(ok)
	t.Log(err)
}
