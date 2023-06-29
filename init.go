package main

import (
	"sconesAPI/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	db, err := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.AutoMigrate(&models.AllowedUser{}, &models.Question{})
	if err != nil {
		panic("Failed to migrate database schema: " + err.Error())
	}

	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to get underlying SQL database: " + err.Error())
	}
	defer dbSQL.Close()
}
