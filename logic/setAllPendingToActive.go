package logic

import (
	"sconesAPI/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetAllPendingToActive(db *gorm.DB, c *fiber.Ctx) error {
	err := db.Model(&models.Question{}).Where("status = ?", "pending").Update("status", "active").Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": " Failed to switch Pending questions to active"})
	}

	return c.JSON(fiber.Map{"message": "Pending questions successfully switched to Active"})
}
