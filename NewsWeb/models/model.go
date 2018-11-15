package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	UserName string `orm:"unique"`
	Pwd      string
}

func init() {
	orm.RegisterDataBase("default", "mysql","root:123@tcp(127.0.0.1:3306)/newsweb?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
