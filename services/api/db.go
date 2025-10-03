package api

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	// Auto migrate the schema
	err = db.AutoMigrate(&SampleModel{})
	if err != nil {
		return err
	}

	return nil
}
