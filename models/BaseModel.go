package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type BaseModel struct {

}

func init() {
	dbHost := beego.AppConfig.String("db.host")
	dbUser := beego.AppConfig.String("db.user")
	dbPwd := beego.AppConfig.String("db.password")
	dbPort := beego.AppConfig.String("db.port")
	dbName := beego.AppConfig.String("db.name")
	//dbTimezone := beego.AppConfig.String("db.timezone")

	// set default database
	dsn := dbUser + ":" + dbPwd + "@tcp("+ dbHost + ":" + dbPort +")/" + dbName + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(ProjectConfig))
	orm.Debug = true
	fmt.Println(21312312)
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}