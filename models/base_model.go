package models

import (
	_ "beego_action/helpers"
	"database/sql"
	"fmt"

	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	fmt.Println("start base_model init")
	GetConnDB("")
}

func GetConnDB(db_name string) {
	DbConfigInfo := GetDbConfig()

	if db_name == "" {
		db, _ = sql.Open("mysql", DbConfigInfo["User"]+":"+DbConfigInfo["Pwd"]+"@/"+DbConfigInfo["DbName"]+"?charset=utf8")
	} else {
		db, _ = sql.Open("mysql", DbConfigInfo["User"]+":"+DbConfigInfo["Pwd"]+"@/"+db_name+"?charset=utf8")
	}
	fmt.Println("mysql", DbConfigInfo["User"]+":"+DbConfigInfo["Pwd"]+"@/"+DbConfigInfo["DbName"]+"?charset=utf8")

	//db.SetMaxOpenConns(2000)
	//db.SetMaxIdleConns(1000)
	//db.Ping()
}

func GetDbConfig() map[string]string {
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	DbConfigInfo := make(map[string]string)
	DbConfigInfo["DbHost"] = iniconf.String("mysqlurls")
	DbConfigInfo["Port"] = iniconf.String("mysqlport")
	DbConfigInfo["User"] = iniconf.String("mysqluser")
	DbConfigInfo["Pwd"] = iniconf.String("mysqlpass")
	DbConfigInfo["DbName"] = iniconf.String("mysqldb")

	return DbConfigInfo
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
