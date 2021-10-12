// package main is reserved to server and database setup
package main

import (
	"blog/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Creates new fiber app
	var app *fiber.App = fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	// Creates a handler struct for posts
	var ph handlers.PostsHandler = handlers.NewPostsHandler()
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
