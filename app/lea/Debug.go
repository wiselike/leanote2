package lea

import (
	"encoding/json"
	"fmt"

	"github.com/wiselike/revel"
)

var leanoteLog = revel.AppLog.New("module", "leanote")

var LeanoteDBLog = revel.AppLog.New("module", "leanote-db")
var LeanoteFileLog = revel.AppLog.New("module", "leanote-file")

func init() {
	leanoteLog.SetStackDepth(2)

	LeanoteDBLog.SetStackDepth(3)   // 多回溯一层，显示执行db的行
	LeanoteFileLog.SetStackDepth(3) // 多回溯一层，显示执行文件操作的行
}

func Log(msg string, i ...interface{}) {
	leanoteLog.Info(msg, i...)
}

func Logf(msg string, i ...interface{}) {
	leanoteLog.Infof(msg, i...)
}

func LogW(msg string, i ...interface{}) {
	leanoteLog.Warn(msg, i...)
}

func LogE(msg string, i ...interface{}) {
	leanoteLog.Error(msg, i...)
}

func LogJ(i interface{}) {
	b, _ := json.MarshalIndent(i, "", " ")
	leanoteLog.Info(string(b))
}

// 为test用
func L(i interface{}) {
	fmt.Println(i)
}

func LJ(i interface{}) {
	b, _ := json.MarshalIndent(i, "", " ")
	fmt.Println(string(b))
}
