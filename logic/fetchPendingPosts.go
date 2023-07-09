package logic

import (
	"fmt"
	"sconesAPI/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FetchPendingPosts(c *fiber.Ctx, db *gorm.DB) error {

	var uncuratedPosts []models.Question
	db = db.Debug()

	err := db.Find(&uncuratedPosts).Where("status = ?", "pending").Error
	if err != nil {
		fmt.Println("Error retrieving items:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	return c.JSON(uncuratedPosts)
}
