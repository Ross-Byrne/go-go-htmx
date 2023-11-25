package main

import (
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Post struct {
	Id        uint64    `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	AuthorId  uint64    `json:"authorId"`
	CreatedAt time.Time `json:"createdAt"`
}

var postsMap = map[uint64]Post{
	1: {Id: 1, Title: "First Post", Text: "This is the body of the post", AuthorId: 1, CreatedAt: time.Now().UTC()},
	2: {Id: 2, Title: "Post 2", Text: "This is the body of the post", AuthorId: 1, CreatedAt: time.Now().UTC()},
	3: {Id: 3, Title: "Post 3", Text: "This is the body of the post", AuthorId: 1, CreatedAt: time.Now().UTC()},
	4: {Id: 4, Title: "Post 4", Text: "This is the body of the post", AuthorId: 1, CreatedAt: time.Now().UTC()},
	5: {Id: 5, Title: "Post 5", Text: "This is the body of the post", AuthorId: 1, CreatedAt: time.Now().UTC()},
}

func findPost(id uint64) Post {
	log.Println("ID:", id)

	// find post
	post, exists := postsMap[id]

	if exists {
		return post
	}
	log.Println("failed to find post:", id)

	return Post{}
}

// func updateContact(contact Contact) {
// 	// Note: this is terrible, use a hashmap
// 	for index, value := range contacts {
// 		if value.ID == contact.ID {
// 			contacts[index].FirstName = contact.FirstName
// 			contacts[index].LastName = contact.LastName
// 			contacts[index].Email = contact.Email
// 			break
// 		}
// 	}
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
	app.Get("/", homeGet)
	app.Get("/profile", profileGet)

	app.Route("/post", func(router fiber.Router) {
		router.Get("/:id", postGet)
		// router.Put("/:id", postPut)
		// router.Get("/:id/edit", postEdit)
	}, "post")

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func homeGet(c *fiber.Ctx) error {
	log.Println("Hello home page")

	listOfPosts := allPosts()

	// Sort list by id
	sort.SliceStable(listOfPosts, func(i, j int) bool {
		return listOfPosts[i].Id < listOfPosts[j].Id
	})

	return c.Render("home/index", fiber.Map{
		"ShowSearch": true,
		"Posts":      listOfPosts,
	})
}

func profileGet(c *fiber.Ctx) error {
	return c.Render("profile/index", fiber.Map{})
}

func postGet(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		log.Println("Could not parse ID {} as int", c.Params("id"))
		return c.SendStatus(422)
	}

	post := findPost(id)

	// Return 404 if no post found
	if post == (Post{}) {
		return c.SendStatus(404)
	}

	return c.Render("post/index", fiber.Map{
		"Post": post,
	})
}

// func contactPut(c *fiber.Ctx) error {
// 	contact := Contact{}

// 	if err := c.BodyParser(contact); err != nil {

// 		// var updatedContact = Contact{
// 		// 	ID:        c.Params("id"),
// 		// 	FirstName: c.Params("firstName"),
// 		// 	LastName:  c.Params("lastName"),
// 		// 	Email:     c.Params("email"),
// 		// }

// 		// Update Record
// 		updateContact(contact)

// 		// Render form
// 		return contactShow(c)
// 	}

// 	return c.SendStatus(500)
// }

// func contactEdit(c *fiber.Ctx) error {
// 	var contact Contact = findContact(c.Params("id"))

// 	return c.Render("home/form", fiber.Map{
// 		"id":        contact.ID,
// 		"firstName": contact.FirstName,
// 		"lastName":  contact.LastName,
// 		"email":     contact.Email,
// 	})
// }

func allPosts() []Post {
	s := make([]Post, 0, len(postsMap))
	for _, v := range postsMap {
		s = append(s, v)
	}
	return s
}
