package util

import (
	"ginWeb/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngin *xorm.Engine

func CreateOrmDb() (*xorm.Engine, error) {
	db := config.AppConfig.DbInfo
	conn := db.Username + ":" + db.PassWord + "@tcp(" + db.Adress + ":" + db.Port + ")/" + db.DbName
	engine, err := xorm.NewEngine("mysql", conn)
	engine.ShowSQL(db.Showsql)
	if err != nil {
		return nil, err
	}
	DbEngin = engine
	return engine, nil
}
