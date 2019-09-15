package infra

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func Init() {
	var err error
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", os.Getenv("USER_NAME"), os.Getenv("PASSWORD"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		panic(err)
	}
}

func ConnectDB() *xorm.Engine {
	return engine
}
