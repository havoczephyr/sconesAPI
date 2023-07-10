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

	app.Get("/api/allPendingPosts", func(c *fiber.Ctx) error {
		return logic.FetchPendingPosts(c, db)
	})

	app.Put("/api/allPendingToActive", func(c *fiber.Ctx) error {
		return logic.SetAllPendingToActive(db, c)
	})

	app.Put("/api/denyPost", func(c *fiber.Ctx) error {
		return logic.DenyPost(c, db)
	})

	app.Get("/api/allDeniedPost", func(c *fiber.Ctx) error {
		return c.SendString("Get all quotes denied in Curation")
	})

	app.Get("/api/allUsedPost", func(c *fiber.Ctx) error {
		return c.SendString("Get all quotes already used in QOTD")
	})

	app.Get("/api/allActivePosts", func(c *fiber.Ctx) error {
		return c.SendString("Get all quotes pending use by QOTD")
	})

	app.Get("/api/useNextActivePost", func(c *fiber.Ctx) error {
		return logic.UseNextActivePost(c, db)
	})

	app.Post("/api/createNewQuote", func(c *fiber.Ctx) error {
		err := logic.CreateNewPost(c, db)
		if err != nil {
			fiber.NewError(401, "Unauthorized")
		}
		return c.SendString("Created a new quote")
	})
	app.Post("/api/createPriorityQuote", func(c *fiber.Ctx) error {
		err := logic.CreatePriorityPost(c, db)
		if err != nil {
			fiber.NewError(401, "Unauthorized")
		}
		return c.SendString("Created a new priority quote")
	})

}

type createUserBody struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
