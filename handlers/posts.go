// package handlers contains all the route handlers
package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// PostsHandler is a struct managing posts related handlers dependencies ex: database
type PostsHandler struct {
}

// NewPostsHandler is the constuctor of handlers.PostsHandler
func NewPostsHandler() PostsHandler {
	return PostsHandler{}
}

// GetPosts returns to the client all the posts stored in the database
func (p *PostsHandler) GetPosts(c *fiber.Ctx) error {
	c.SendString("Getting posts")

	return nil
}

// GetPost returns to the client a specific post based on its given id in the route
func (p *PostsHandler) GetPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrNotFound
	}
	c.SendString("Getting post " + strconv.FormatInt(int64(id), 32))

	return nil
}

// AddPost adds a post populated by the data in the request body to the database
func (p *PostsHandler) CreatePost(c *fiber.Ctx) error {
	c.SendString("Creating post")

	return nil
}

// AddPost updates a post in the database with the data in the request body and given id in its route
func (p *PostsHandler) UpdatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrNotFound
	}
	c.SendString("Updating post " + strconv.FormatInt(int64(id), 32))

	return nil
}

// AddPost detetes the post with the id given in its route
func (p *PostsHandler) DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrNotFound
	}
	c.SendString("Deleting post " + strconv.FormatInt(int64(id), 32))

	return nil
}
