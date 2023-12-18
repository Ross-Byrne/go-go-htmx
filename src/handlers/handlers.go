package handlers

import (
	"fmt"
	"go-go-htmx/src/helpers"
	"go-go-htmx/src/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	db, err := helpers.ConnectToDB(c)

	if err != nil {
		return err
	}

	var posts []models.Post
	result := db.Order("created_at DESC").Find(&posts)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.Render("home/index", fiber.Map{
		"Posts": posts,
	})
}

func GetProfile(c *fiber.Ctx) error {
	return c.Render("profile/index", fiber.Map{})
}

func GetPost(c *fiber.Ctx) error {
	db, err := helpers.ConnectToDB(c)

	if err != nil {
		return err
	}

	var post models.Post
	db.First(&post, c.Params("id"))

	if post == (models.Post{}) {
		return c.SendStatus(404)
	}

	return c.Render("post/index", fiber.Map{
		"Post": post,
	})
}

func GetCreatePost(c *fiber.Ctx) error {
	return c.Render("createPost/index", fiber.Map{})
}

func PostCreatePost(c *fiber.Ctx) error {
	db, err := helpers.ConnectToDB(c)

	if err != nil {
		return err
	}

	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		log.Println("Error trying to parse request body")
		return c.SendStatus(500)
	}

	result := db.Create(&post)

	if result.Error != nil {
		log.Println("Error trying to creating Post")
		return c.SendStatus(500)
	}

	newLocation := fmt.Sprintf("/post/%d", post.ID)
	c.Set("HX-Location", newLocation)

	return c.Render("post/index", fiber.Map{
		"Post": post,
	})
}
