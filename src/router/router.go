package router

import (
	"go-go-htmx/src/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Static("/", "./src/assets/")

	// Routes
	app.Get("/", handlers.GetHome)
	app.Get("/profile", handlers.ProfileGet)

	app.Route("/post", func(router fiber.Router) {
		router.Get("/create", handlers.CreatePostGet)
		router.Post("/create", handlers.CreatePostPost)

		router.Get("/:id", handlers.PostGet)
	}, "post")
	app.Get("/create", handlers.CreatePostGet)
}
