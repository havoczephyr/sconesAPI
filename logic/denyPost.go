package logic

import (
	"sconesAPI/models"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DenyPost(c *fiber.Ctx, db *gorm.DB) error {
	parsedBody := new(denyPostBody)
	err := c.BodyParser(parsedBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": " Failed to Parse Body"})
	}
	err = db.Model(&models.Question{}).Where("id = ?", parsedBody.PostID).Update("status", "denied").Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to set status to denied"})
	}
	jsonPost := fmt.Sprintf("%s %d %s", "Post #", parsedBody.PostID, " has been denied.")
	return c.JSON(fiber.Map{"message": jsonPost})
}

type denyPostBody struct {
	PostID int `json:"id"`
}
