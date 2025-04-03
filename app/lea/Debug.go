package lea

import (
	"encoding/json"
	"fmt"
	"github.com/wiselike/revel"
	"sync"
)

const defautDepth = 1

var logMutex sync.Mutex

func init() {
	revel.AppLog.SetStackDepth(defautDepth)
}

func Log(msg string, i ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	revel.AppLog.Info(msg, i...)
}

func Logf(msg string, i ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	revel.AppLog.Info(fmt.Sprintf(msg, i...))
}
func Logf2(depth int, msg string, i ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	revel.AppLog.SetStackDepth(depth)
	defer revel.AppLog.SetStackDepth(defautDepth)
	revel.AppLog.Info(fmt.Sprintf(msg, i...))
}

func LogW(msg string, i ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	revel.AppLog.Warn(msg, i...)
}
func LogW2(depth int, msg string, i ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	revel.AppLog.SetStackDepth(depth)
	defer revel.AppLog.SetStackDepth(defautDepth)
	revel.AppLog.Warn(msg, i...)
}

func LogE(msg string, i ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	revel.AppLog.Error(msg, i...)
}

func LogE2(depth int, msg string, i ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	revel.AppLog.SetStackDepth(depth)
	defer revel.AppLog.SetStackDepth(defautDepth)
	revel.AppLog.Error(msg, i...)
}

func LogJ(i interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	b, _ := json.MarshalIndent(i, "", " ")
	revel.AppLog.Info(string(b))
}

// 为test用
func L(i interface{}) {
	fmt.Println(i)
}

func LJ(i interface{}) {
	b, _ := json.MarshalIndent(i, "", " ")
	fmt.Println(string(b))
}
