package logic

import (
	"math/rand"
	"sconesAPI/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UseNextActivePost(c *fiber.Ctx, db *gorm.DB) error {
	var activePosts []models.Question

	err := db.Where("status = ?", "active").Find(&activePosts).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to fetch active posts"})
	}

	randomIndex := rand.Intn(len(activePosts))
	randomPost := activePosts[randomIndex]

	err = db.Model(&randomPost).Update("status", "used").Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to update post status"})
	}

	return c.JSON(randomPost)
}
