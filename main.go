package main

import (
	"flag"
	"go-go-htmx/src/helpers"
	"go-go-htmx/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func mainInit() {
	var initDB bool
	var seedDB bool

	// flags declaration using flag package
	flag.BoolVar(&initDB, "init-db", false, "Initialise database. Default is false")
	flag.BoolVar(&seedDB, "seed-db", false, "Seed database with test data. Default is false")

	flag.Parse() // after declaring flags we need to call it

	if initDB {
		log.Println("Initialising Database...")

		helpers.InitialiseDatabase(seedDB)
	}
}

func main() {
	mainInit()

	// Initialize standard Go html template engine
	engine := html.New("./src/views", ".html")

	engine.Reload(true)

	// Fiber instance
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	router.SetupRouter(app)

	// Start server
	log.Fatal(app.Listen(":3001"))
}
