package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./src/views", ".html")

	engine.Reload(true)

	// Fiber instance
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/", "./src/assets/")

	// Routes
	app.Get("/", hello)
	app.Get("/albums", getAlbums)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error {
	// Render index template
	return c.Render("home", fiber.Map{
		"HeaderTitle": "This is the header!",
		"Title":       "Hello, World!",
	})
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Hello, World!",
	})
}
