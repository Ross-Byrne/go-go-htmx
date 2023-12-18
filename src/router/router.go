package router

import (
	"go-go-htmx/src/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Static("/", "./src/assets/")

	// Routes
	app.Get("/", handlers.GetHome)
	app.Get("/profile", handlers.GetProfile)

	app.Route("/post", func(router fiber.Router) {
		router.Get("/create", handlers.GetCreatePost)
		router.Post("/create", handlers.PostCreatePost)

		router.Get("/:id", handlers.GetPost)
	}, "post")
}
