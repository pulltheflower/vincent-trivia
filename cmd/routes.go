package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pulltheflower/vincent-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/fact", handlers.CreateFact)
	app.Get("/facts", handlers.IndexFacts)
}
