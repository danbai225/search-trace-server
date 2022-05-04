package db

import (
	"database/sql"
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
	tx := GetDBW()
	var err error
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Callback()
		}
	}()

	err = tx.AutoMigrate(&model.User{}, &model.Blacklist{}, &model.Word{})
	if err == nil {
		arr := make([]map[string]interface{}, 0)
		err = tx.Raw("SHOW INDEX FROM `trace` WHERE Key_name=\"all_text_index\";").Scan(&arr).Error
		if err == nil && len(arr) == 0 {
			tx.Exec("CREATE FULLTEXT INDEX `all_text_index` ON `trace`(`title`,`content`) COMMENT '全文索引' /*!50100 WITH PARSER `ngram` */;")
		}
	}
}
func GetDBW() *gorm.DB {
	if config.C.Db.Debug {
		return db.Debug().Begin()
	}
	return db.Begin()
}
func GetDBR() *gorm.DB {
	if config.C.Db.Debug {
		return db.Debug().Begin(&sql.TxOptions{
			ReadOnly: true,
		})
	}
	return db
}
