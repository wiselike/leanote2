package tests

import (
	"testing"

	"github.com/wiselike/leanote2/app/db"
	//	. "github.com/wiselike/leanote2/app/lea"
	//	"github.com/wiselike/leanote2/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func TestDBConnect(t *testing.T) {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
}
