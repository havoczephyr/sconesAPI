package main

import (
	"sconesAPI/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	app := fiber.New()
	routes.Router(app, db)
	app.Listen(":3000")

	// initDB()
}
