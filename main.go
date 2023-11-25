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
	1: {Id: 1, Title: "First Post", Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean sollicitudin, elit sed mollis hendrerit, turpis lectus pellentesque arcu, ut pharetra nibh risus bibendum urna. Cras bibendum gravida orci ut vehicula. Ut egestas neque sed imperdiet vulputate. Integer faucibus consequat ante ut posuere. Suspendisse quis diam quis eros dictum feugiat eu at dui. Donec id sollicitudin erat. Aliquam pulvinar purus eu venenatis posuere. Donec eleifend aliquam nunc, nec pellentesque ex sagittis non. Mauris luctus sodales mi vitae pretium. Sed volutpat metus eu justo iaculis, eget venenatis velit sagittis. Curabitur sem neque, euismod in lectus non, hendrerit ullamcorper leo. Aliquam tempor, mi sit amet sagittis mollis, lacus lectus convallis mi, ut tincidunt nibh purus vel lorem. Proin a tortor neque. Aenean magna nulla, vestibulum non mi et, rutrum condimentum velit.", AuthorId: 1, CreatedAt: time.Now().UTC()},
	2: {Id: 2, Title: "Post 2", Text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean sollicitudin, elit sed mollis hendrerit, turpis lectus pellentesque arcu, ut pharetra nibh risus bibendum urna. Cras bibendum gravida orci ut vehicula. Ut egestas neque sed imperdiet vulputate. Integer faucibus consequat ante ut posuere. Suspendisse quis diam quis eros dictum feugiat eu at dui. Donec id sollicitudin erat. Aliquam pulvinar purus eu venenatis posuere. Donec eleifend aliquam nunc, nec pellentesque ex sagittis non. Mauris luctus sodales mi vitae pretium. Sed volutpat metus eu justo iaculis, eget venenatis velit sagittis. Curabitur sem neque, euismod in lectus non, hendrerit ullamcorper leo. Aliquam tempor, mi sit amet sagittis mollis, lacus lectus convallis mi, ut tincidunt nibh purus vel lorem. Proin a tortor neque. Aenean magna nulla, vestibulum non mi et, rutrum condimentum velit.

	Morbi sed tempus mi. Aenean eget lorem et neque rutrum blandit. Etiam ut mattis enim. Maecenas ornare sagittis malesuada. Aliquam tortor nibh, porttitor vitae libero ut, pharetra auctor diam. Duis a consequat massa. Morbi est dolor, consequat vel libero ultricies, cursus semper neque. Cras consectetur porttitor odio ut rutrum. Nulla sollicitudin vehicula viverra. Vestibulum efficitur dolor sit amet tortor blandit feugiat. In nec neque arcu. Etiam et convallis urna, non malesuada lorem. Integer semper felis vitae lorem tincidunt placerat. Mauris auctor est at egestas blandit.

	Nulla eget est tortor. Aliquam in tellus est. Fusce nec lectus ut eros vulputate faucibus. Donec eget aliquet erat. Ut nisi elit, volutpat quis enim ut, tristique vulputate ex. Nam nulla turpis, mollis sit amet dictum id, tempus eu est. Donec sed leo sit amet erat aliquam condimentum. Ut placerat ligula id turpis tristique, eget sagittis neque gravida.

	Donec laoreet, erat et sollicitudin fringilla, turpis est mollis ex, a pellentesque dolor turpis a orci. Etiam et malesuada ipsum. Proin rhoncus ante egestas, accumsan enim eu, blandit est. Donec et pulvinar nisl. Donec lacinia risus et lacinia efficitur. Aliquam urna purus, aliquam imperdiet vestibulum sit amet, sodales in urna. Pellentesque turpis nulla, fringilla a tortor vitae, accumsan condimentum ex. Praesent convallis elit a hendrerit finibus. Fusce mattis egestas tortor, ut vehicula lectus consequat eget. Nullam euismod ante metus, at iaculis elit commodo a. In hac habitasse platea dictumst. Donec tincidunt, tortor vel luctus euismod, nibh sem tempor nunc, et suscipit neque purus eget mi. Proin hendrerit tincidunt mollis. Phasellus eu risus eget libero dapibus venenatis nec ac nunc. Aenean libero nulla, fringilla in lectus vitae, auctor bibendum justo. Praesent non mauris eget ex elementum fringilla in in eros.

	Vestibulum tincidunt finibus faucibus. Nulla auctor dictum elementum. Cras fringilla ex auctor urna imperdiet, faucibus hendrerit mauris laoreet. Nulla cursus dolor sed nisl luctus sollicitudin. Vestibulum porta aliquet aliquam. Donec finibus felis hendrerit aliquet pellentesque. Maecenas sed ligula et arcu vestibulum consectetur ac ut neque. Maecenas malesuada, lectus vulputate egestas tempus, orci velit ullamcorper lorem, ut tincidunt enim libero non velit. Pellentesque bibendum vel sem quis tincidunt. Duis at quam sollicitudin, aliquet sem sit amet, blandit nibh. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur commodo a metus nec congue. Vivamus ac nibh ac sem volutpat malesuada vel vitae risus. Duis in pretium risus. Fusce a eleifend neque, et finibus purus.
	`, AuthorId: 1, CreatedAt: time.Now().UTC()},
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
