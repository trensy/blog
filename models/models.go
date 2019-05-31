package models

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var  db *gorm.DB


func Init() {
	var err error

	db, err = gorm.Open("mysql", "root:root123@tcp(127.0.0.1)/blog_article?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
	//db.LogMode(true)
	if err != nil {
		log.Fatalf("conn mysql error: %v", err);
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "blog_" + defaultTableName
	}

	//连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.SingularTable(true)

}