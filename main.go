package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Contact struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var contacts = []Contact{
	{ID: "1", FirstName: "John", LastName: "Doe", Email: "john.doe@outlook.com"},
	{ID: "2", FirstName: "Jim", LastName: "O'Tool", Email: "jimmy@gmail.com"},
	{ID: "3", FirstName: "Frank", LastName: "Crews", Email: "frank.crews@hotmail.com"},
}

func findContact(id string) Contact {
	for _, value := range contacts {
		if value.ID == id {
			return value
		}
	}

	return Contact{}
}

func updateContact(contact Contact) {
	// Note: this is terrible, use a hashmap
	for index, value := range contacts {
		if value.ID == contact.ID {
			contacts[index].FirstName = contact.FirstName
			contacts[index].LastName = contact.LastName
			contacts[index].Email = contact.Email
			break
		}
	}
}

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
	app.Get("/", homepage)

	app.Route("/contact", func(router fiber.Router) {
		router.Get("/:id", contactShow)
		router.Put("/:id", contactPut)
		router.Get("/:id/edit", contactEdit)
	}, "contact")

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func homepage(c *fiber.Ctx) error {
	return c.Render("homePage/home", fiber.Map{})
}

func contactShow(c *fiber.Ctx) error {
	var contact Contact = findContact(c.Params("id"))

	return c.Render("homePage/show", fiber.Map{
		"id":        contact.ID,
		"firstName": contact.FirstName,
		"lastName":  contact.LastName,
		"email":     contact.Email,
	})
}

func contactPut(c *fiber.Ctx) error {
	contact := Contact{}

	if err := c.BodyParser(contact); err != nil {

		// var updatedContact = Contact{
		// 	ID:        c.Params("id"),
		// 	FirstName: c.Params("firstName"),
		// 	LastName:  c.Params("lastName"),
		// 	Email:     c.Params("email"),
		// }

		// Update Record
		updateContact(contact)

		// Render form
		return contactShow(c)
	}

	return c.SendStatus(500)
}

func contactEdit(c *fiber.Ctx) error {
	var contact Contact = findContact(c.Params("id"))

	return c.Render("homePage/form", fiber.Map{
		"id":        contact.ID,
		"firstName": contact.FirstName,
		"lastName":  contact.LastName,
		"email":     contact.Email,
	})
}
