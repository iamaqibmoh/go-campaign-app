package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DBConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa_campaign_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
