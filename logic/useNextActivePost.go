package logic

import (
	"math/rand"
	"sconesAPI/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UseNextActivePost(c *fiber.Ctx, db *gorm.DB) error {
	var priorityPosts []models.Question
	var activePosts []models.Question

	err := db.Where("status = ?", "active").Where("priority = ?", true).Find(&priorityPosts).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to fetch priority active posts"})
	}

	if len(priorityPosts) > 0 {
		err = db.Model(&priorityPosts[0]).Update("status", "used").Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to update post status"})
		}
		return c.JSON(priorityPosts[0])
	} else {
		err = db.Where("status = ?", "active").Find(&activePosts).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to fetch active posts"})
		}
		if len(activePosts) > 0 {
			randomIndex := rand.Intn(len(activePosts))
			randomPost := activePosts[randomIndex]

			err = db.Model(&randomPost).Update("status", "used").Error
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to update post status"})
			}

			return c.JSON(randomPost)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Oops! No Active Quotes :("})
		}

	}

}
