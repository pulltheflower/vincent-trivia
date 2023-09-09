package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pulltheflower/vincent-trivia/database"
	"github.com/pulltheflower/vincent-trivia/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Vincent App!")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func IndexFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}
