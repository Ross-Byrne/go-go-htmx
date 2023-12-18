package helpers

import (
	"log"

	"go-go-htmx/src/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB(c *fiber.Ctx) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
		if c != nil {
			return nil, c.SendStatus(500)
		}

		return nil, err
	}

	return db, nil
}

func InitialiseDatabase(seedData bool) {
	db, err := ConnectToDB(nil)

	if err != nil {
		log.Fatal("Failed to connect database")
		return
	}

	// Migrate the schema
	db.AutoMigrate(&models.Post{})

	// Create seed data
	if seedData {
		log.Println("Seeding Database...")

		db.Create(&models.Post{Title: "First Post", Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean sollicitudin, elit sed mollis hendrerit, turpis lectus pellentesque arcu, ut pharetra nibh risus bibendum urna. Cras bibendum gravida orci ut vehicula. Ut egestas neque sed imperdiet vulputate. Integer faucibus consequat ante ut posuere. Suspendisse quis diam quis eros dictum feugiat eu at dui. Donec id sollicitudin erat. Aliquam pulvinar purus eu venenatis posuere. Donec eleifend aliquam nunc, nec pellentesque ex sagittis non. Mauris luctus sodales mi vitae pretium. Sed volutpat metus eu justo iaculis, eget venenatis velit sagittis. Curabitur sem neque, euismod in lectus non, hendrerit ullamcorper leo. Aliquam tempor, mi sit amet sagittis mollis, lacus lectus convallis mi, ut tincidunt nibh purus vel lorem. Proin a tortor neque. Aenean magna nulla, vestibulum non mi et, rutrum condimentum velit.", AuthorID: 1})
		db.Create(&models.Post{Title: "Post 2", Text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean sollicitudin, elit sed mollis hendrerit, turpis lectus pellentesque arcu, ut pharetra nibh risus bibendum urna. Cras bibendum gravida orci ut vehicula. Ut egestas neque sed imperdiet vulputate. Integer faucibus consequat ante ut posuere. Suspendisse quis diam quis eros dictum feugiat eu at dui. Donec id sollicitudin erat. Aliquam pulvinar purus eu venenatis posuere. Donec eleifend aliquam nunc, nec pellentesque ex sagittis non. Mauris luctus sodales mi vitae pretium. Sed volutpat metus eu justo iaculis, eget venenatis velit sagittis. Curabitur sem neque, euismod in lectus non, hendrerit ullamcorper leo. Aliquam tempor, mi sit amet sagittis mollis, lacus lectus convallis mi, ut tincidunt nibh purus vel lorem. Proin a tortor neque. Aenean magna nulla, vestibulum non mi et, rutrum condimentum velit.

		Morbi sed tempus mi. Aenean eget lorem et neque rutrum blandit. Etiam ut mattis enim. Maecenas ornare sagittis malesuada. Aliquam tortor nibh, porttitor vitae libero ut, pharetra auctor diam. Duis a consequat massa. Morbi est dolor, consequat vel libero ultricies, cursus semper neque. Cras consectetur porttitor odio ut rutrum. Nulla sollicitudin vehicula viverra. Vestibulum efficitur dolor sit amet tortor blandit feugiat. In nec neque arcu. Etiam et convallis urna, non malesuada lorem. Integer semper felis vitae lorem tincidunt placerat. Mauris auctor est at egestas blandit.

		Nulla eget est tortor. Aliquam in tellus est. Fusce nec lectus ut eros vulputate faucibus. Donec eget aliquet erat. Ut nisi elit, volutpat quis enim ut, tristique vulputate ex. Nam nulla turpis, mollis sit amet dictum id, tempus eu est. Donec sed leo sit amet erat aliquam condimentum. Ut placerat ligula id turpis tristique, eget sagittis neque gravida.

		Donec laoreet, erat et sollicitudin fringilla, turpis est mollis ex, a pellentesque dolor turpis a orci. Etiam et malesuada ipsum. Proin rhoncus ante egestas, accumsan enim eu, blandit est. Donec et pulvinar nisl. Donec lacinia risus et lacinia efficitur. Aliquam urna purus, aliquam imperdiet vestibulum sit amet, sodales in urna. Pellentesque turpis nulla, fringilla a tortor vitae, accumsan condimentum ex. Praesent convallis elit a hendrerit finibus. Fusce mattis egestas tortor, ut vehicula lectus consequat eget. Nullam euismod ante metus, at iaculis elit commodo a. In hac habitasse platea dictumst. Donec tincidunt, tortor vel luctus euismod, nibh sem tempor nunc, et suscipit neque purus eget mi. Proin hendrerit tincidunt mollis. Phasellus eu risus eget libero dapibus venenatis nec ac nunc. Aenean libero nulla, fringilla in lectus vitae, auctor bibendum justo. Praesent non mauris eget ex elementum fringilla in in eros.

		Vestibulum tincidunt finibus faucibus. Nulla auctor dictum elementum. Cras fringilla ex auctor urna imperdiet, faucibus hendrerit mauris laoreet. Nulla cursus dolor sed nisl luctus sollicitudin. Vestibulum porta aliquet aliquam. Donec finibus felis hendrerit aliquet pellentesque. Maecenas sed ligula et arcu vestibulum consectetur ac ut neque. Maecenas malesuada, lectus vulputate egestas tempus, orci velit ullamcorper lorem, ut tincidunt enim libero non velit. Pellentesque bibendum vel sem quis tincidunt. Duis at quam sollicitudin, aliquet sem sit amet, blandit nibh. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur commodo a metus nec congue. Vivamus ac nibh ac sem volutpat malesuada vel vitae risus. Duis in pretium risus. Fusce a eleifend neque, et finibus purus.
		`, AuthorID: 1})
		db.Create(&models.Post{Title: "Post 3", Text: "This is the body of the post", AuthorID: 1})
		db.Create(&models.Post{Title: "Post 4", Text: "This is the body of the post", AuthorID: 1})
		db.Create(&models.Post{Title: "Post 5", Text: "This is the body of the post", AuthorID: 1})

		db.Create(&models.Post{Title: "P1", Text: "Hello, there!", AuthorID: 1})
		db.Create(&models.Post{Title: "P2", Text: "Hello, there!", AuthorID: 1})
		db.Create(&models.Post{Title: "P3", Text: "Hello, there!", AuthorID: 1})
		db.Create(&models.Post{Title: "P4", Text: "Hello, there!", AuthorID: 1})
	}

	// Read
	// var post Post
	// db.First(&post, 1)                 // find product with integer primary key
	// db.First(&post, "title = ?", "P1") // find product with code D42

	// // Delete - delete product
	// db.Delete(&post, 1)
}
