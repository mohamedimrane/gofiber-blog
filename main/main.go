// package main is reserved to server and database setup
package main

import (
	"os"
	"os/signal"

	"github.com/mohamedimrane/gofiber-blog/data"
	"github.com/mohamedimrane/gofiber-blog/handlers"
	"github.com/mohamedimrane/gofiber-blog/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Creates new fiber app
	var app *fiber.App = fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	// Creates database if it does not exist and opens a connection with it
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		utils.PrintRed(" Failed to connect to database ")
		os.Exit(1)
	} else {
		utils.PrintGreen(" Connect to database successfuly ")
	}

	// Migartes the database with models
	err = db.AutoMigrate(&data.Post{}, &data.Tag{})
	if err != nil {
		utils.PrintRed(" Failed to migrate database ")
		os.Exit(1)
	} else {
		utils.PrintGreen(" Migrated database successfuly ")
	}

	// Registers a middleware that changes Content-Type header to application/json
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")

		return c.Next()
	})

	// Creates a handler struct for posts
	var ph *handlers.PostsHandler = handlers.NewPostsHandler(db)
	// Creates a handler struct for tags
	var th *handlers.TagsHandler = handlers.NewTagsHandler(db)

	// Creates a router group for posts
	var posts fiber.Router = app.Group("/posts/")
	// Creates a router group for posts
	var tags fiber.Router = app.Group("/tags/")

	// Registers routes for posts
	posts.Get("/", ph.GetPosts)
	posts.Get("/:id", ph.GetPost)
	posts.Post("/", ph.CreatePost)
	posts.Patch("/:id", ph.UpdatePost)
	posts.Delete("/:id", ph.DeletePost)

	// Registers routes for tags
	tags.Get("/", th.GetTags)

	// Launches the server
	go func(app *fiber.App) {
		app.Listen(":8080")
	}(app)

	// Gracefully shutdown when ordered to
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)

	sig := <-sigChannel

	utils.PrintYellow("\n Gracefuly shutting down (" + sig.String() + ") ")
	app.Shutdown()
}
