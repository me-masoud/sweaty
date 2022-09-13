package Controller

import (
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func DataBase()  {
	dsn := "root:Masoud12@@tcp(127.0.0.1:3306)/sweaty"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

}