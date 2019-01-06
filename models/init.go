package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库连接实例
var DB *gorm.DB

func init() {
	fmt.Println("init models...")
	var err error
	DB, err = gorm.Open("mysql", "root:gaojian@/moly?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&Service{}, &Rule{})
}
