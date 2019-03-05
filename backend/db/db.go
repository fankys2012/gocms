package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

//全局DB
var XORM *xorm.Engine

func init() {
	var err error
	XORM, err = xorm.NewEngine("mysql", "nn_cms:nn_cms1234@tcp(192.168.95.55:3306)/nn_cms?charset=utf8")
	if err != nil {
		fmt.Printf("MySql err : %v", err)
	}
	XORM.ShowSQL(true)
	XORM.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "nns_"))
	XORM.SetColumnMapper(core.NewPrefixMapper(core.SnakeMapper{}, "nns_"))
}
