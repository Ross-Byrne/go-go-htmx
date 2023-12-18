package handlers

import (
	"log"

	"go-go-htmx/src/helpers"
	"go-go-htmx/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	log.Println("Hello home page")

	// listOfPosts := allPosts()

	db, err := helpers.ConnectToDB()

	if err != nil {
		log.Fatal("Failed to connect database")
		return c.SendStatus(500)
	}

	var posts []models.Post
	result := db.Order("created_at DESC").Find(&posts)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	// Sort list by id
	// sort.SliceStable(listOfPosts, func(i, j int) bool {
	// 	return listOfPosts[i].ID > listOfPosts[j].ID
	// })

	return c.Render("home/index", fiber.Map{
		"Posts": posts,
	})
}

func ProfileGet(c *fiber.Ctx) error {
	return c.Render("profile/index", fiber.Map{})
}

func PostGet(c *fiber.Ctx) error {
	// id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	// if err != nil {
	// 	log.Println("Could not parse ID {} as int", c.Params("id"))
	// 	return c.SendStatus(422)
	// }

	// post := findPost(id)

	// // Return 404 if no post found
	// if *post == (models.Post{}) {
	// 	return c.SendStatus(404)
	// }

	// return c.Render("post/index", fiber.Map{
	// 	"Post": post,
	// })

	return c.Render("post/index", fiber.Map{})
}

func CreatePostGet(c *fiber.Ctx) error {
	return c.Render("createPost/index", fiber.Map{})
}

func CreatePostPost(c *fiber.Ctx) error {
	// post := new(models.Post)

	// if err := c.BodyParser(post); err != nil {
	// 	return c.SendStatus(500)
	// }

	// post.ID = uint64(len(postsMap) + 1)
	// post.AuthorID = 1
	// post.CreatedAt = time.Now().UTC()
	// postsMap[post.ID] = post

	// log.Println("Created new post with ID: {}", post.ID)

	// newLocation := fmt.Sprintf("/post/%d", post.ID)
	// c.Set("HX-Location", newLocation)

	// return c.Render("post/index", fiber.Map{
	// 	"Post": post,
	// })

	return c.Render("post/index", fiber.Map{})
}
