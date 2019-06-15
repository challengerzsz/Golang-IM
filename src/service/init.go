package service

import (
	"github.com/go-xorm/xorm"
	"log"
	"model"
)

var DbEngine *xorm.Engine
var err error

func init() {
	driverName := "mysql"
	dataSource := "root:agytorudhcv11@(localhost:3306)/golang_im?charset=utf8"
	DbEngine, err = xorm.NewEngine(driverName, dataSource)
	if err != nil {
		log.Fatalf("new engine err %s", err)
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(5)

	DbEngine.Sync2(new(model.User))
}
