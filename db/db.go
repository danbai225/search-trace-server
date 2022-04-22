package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"search-trace-server/config"
	"search-trace-server/model"
)

var db *gorm.DB

func InitDB() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.C.Db.User, config.C.Db.Pass, config.C.Db.Host, config.C.Db.Port, config.C.Db.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	initCache()
	autoMigrate()
	return err
}
func autoMigrate() {
	_ = db.AutoMigrate(&model.User{}, &model.Blacklist{}, &model.Word{})
}
func GetDBW() *gorm.DB {
	if config.C.Db.Debug {
		return db.Debug().Begin()
	}
	return db.Begin()
}
func GetDBR() *gorm.DB {
	if config.C.Db.Debug {
		return db.Debug().Begin()
	}
	return db
}
