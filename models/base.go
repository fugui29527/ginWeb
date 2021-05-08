package models

import (
	"ginWeb/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var mEngine *xorm.Engine

func init() {
	if mEngine == nil {
		db := config.AppConfig.DbInfo
		conn := db.Username + ":" + db.PassWord + "@tcp(" + db.Adress + ":" + db.Port + ")/" + db.DbName
		var err error
		mEngine, err = xorm.NewEngine("mysql", conn)
		if err != nil {
			log.Panic("初始数据库链接失败:", err)
			panic(err)
		}
		mEngine.ShowSQL(db.Showsql)
		mEngine.SetMaxIdleConns(10) //空闲连接
		mEngine.SetMaxOpenConns(50) //最大连接数
		mEngine.ShowExecTime(true)
		mEngine.Ping()
	}
}

func getEngine() *xorm.Engine {
	return mEngine
}
