package Database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sweaty/Config"
)

type Db struct {
	Database *gorm.DB
}

func ConnectToDB() Db {
	config := Config.AppConfig.Database
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/sweaty?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{

		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	return Db{Database: db}
}
