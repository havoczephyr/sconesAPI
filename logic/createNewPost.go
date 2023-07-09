package logic

import (
	"sconesAPI/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateNewPost(c *fiber.Ctx, db *gorm.DB) error {
	parsedBody := new(createPostBody)
	err := c.BodyParser(parsedBody)
	if err != nil {
		return err
	}

	newPost := models.Question{
		Body:   parsedBody.Body,
		Author: parsedBody.Author,
		Status: "pending",
	}

	result := db.Create(&newPost)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type createPostBody struct {
	Body   string
	Author string
}
