package auth

import (
	"sconesAPI/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DbHandler struct {
	Db *gorm.DB
}

func (d *DbHandler) Authentication(c *fiber.Ctx) error {

	authString := c.Get("Authorization")
	authSlice := strings.Split(authString, " ")

	if len(authSlice) < 2 {
		return fiber.NewError(401, "Unauthorized")
	}
	userName := authSlice[0]
	authKey := authSlice[1]

	var user = []models.AllowedUser{}
	d.Db.Where("authorized_user = ?", userName).Find(&user)

	if len(user) < 1 {
		return fiber.NewError(401, "Unauthorized")
	}
	for _, value := range user {
		if verifyPassword(authKey, value.Salt, value.KeyHash) {
			return c.Next()
		}
	}

	return fiber.NewError(401, "Unauthorized")
}
