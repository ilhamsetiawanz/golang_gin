package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dns := "root:@tcp(127.0.0.1:3306)/belajargolang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Buku{})

	DB = db
}
