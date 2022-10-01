// go:build (darwin && cgo) || linux
//go:build (darwin && cgo) || linux
// +build darwin,cgo linux

package Sweater

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Sweater struct {
	Name  string
	Email string
	gorm.Model
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&Sweater{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User migrated")
}
