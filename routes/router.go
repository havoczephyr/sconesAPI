package routes

import (
	"sconesAPI/auth"
	"sconesAPI/logic"
	"sconesAPI/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) {

	dbHandlers := auth.DbHandler{Db: db}
	app.Group("/api", dbHandlers.Authentication)

	app.Post("/createUser", func(c *fiber.Ctx) error {
		parsedBody := new(createUserBody)
		err := c.BodyParser(parsedBody)
		if err != nil {
			return err
		}
		newSalt, err := auth.GenerateSalt()
		if err != nil {
			return err
		}
		newHash := auth.HashPassword(parsedBody.Password, newSalt)
		newUser := models.AllowedUser{
			AuthorizedUser: parsedBody.User,
			KeyHash:        newHash,
			Salt:           newSalt,
		}
		result := db.Create(&newUser)
		if result.Error != nil {
			return result.Error
		}

		return c.SendString("User Created Successfully")
	})

	app.Get("/api/allUncuratedQuotes", func(c *fiber.Ctx) error {
		return c.SendString("Get all quotes pending Curation")
	})

	app.Get("/api/allDeniedQuotes", func(c *fiber.Ctx) error {
		return c.SendString("Get all quotes denied in Curation")
	})

	app.Get("/api/allUsedQuotes", func(c *fiber.Ctx) error {
		return c.SendString("Get all quotes already used in QOTD")
	})

	app.Get("/api/allActiveQuotes", func(c *fiber.Ctx) error {
		return c.SendString("Get all quotes pending use by QOTD")
	})

	app.Get("/api/useNextActiveQuote", func(c *fiber.Ctx) error {
		return c.SendString("Sending the next active quote and marking them as used")
	})

	app.Post("/api/createNewQuote", func(c *fiber.Ctx) error {
		err := logic.CreateNewPost(c, db)
		if err != nil {
			fiber.NewError(401, "Unauthorized")
		}
		return c.SendString("Create a new quote")
	})

}

type createUserBody struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
