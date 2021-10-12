// package main is reserved to server and database setup
package main

import (
	"blog/data"
	"blog/handlers"
	"os"

	"github.com/fatih/color"
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
		color.New(color.BgRed).Add(color.FgWhite).Add(color.Bold).Println(" Failed to connect to database ")
		os.Exit(1)
	} else {
		color.New(color.BgGreen).Add(color.FgWhite).Add(color.Bold).Println(" Connect to database successfuly ")
	}

	// Migartes the database with models
	err = db.AutoMigrate(&data.Post{})
	if err != nil {
		color.New(color.BgRed).Add(color.FgWhite).Add(color.Bold).Println(" Failed to migrate database ")
		os.Exit(1)
	} else {
		color.New(color.BgGreen).Add(color.FgWhite).Add(color.Bold).Println(" Migrated database successfuly ")
	}

	// Creates a handler struct for posts
	var ph handlers.PostsHandler = handlers.NewPostsHandler(db)
	// Creates a router group for posts
	var posts fiber.Router = app.Group("/posts/")

	// Registers routes for posts
	posts.Get("/", ph.GetPosts)
	posts.Get("/:id", ph.GetPost)
	posts.Post("/", ph.CreatePost)
	posts.Patch("/:id", ph.UpdatePost)
	posts.Delete("/:id", ph.DeletePost)

	// Launches the server
	app.Listen(":8080")
}
