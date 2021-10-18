// package handlers contains all the route handlers
package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/gofiber-blog/data"
	"gorm.io/gorm"
)

// PostsHandler is a struct managing posts related handlers dependencies ex: database
type PostsHandler struct {
	DB *gorm.DB
}

// NewPostsHandler is the constuctor of handlers.PostsHandler
func NewPostsHandler(db *gorm.DB) PostsHandler {
	return PostsHandler{
		DB: db,
	}
}

// GetPosts returns to the client all the posts stored in the database
func (p *PostsHandler) GetPosts(c *fiber.Ctx) error {
	var posts []data.Post

	// Getting all the posts from database
	p.DB.Order("created_at desc").Find(&posts)

	// Returning the posts to the client as JSON
	c.JSON(posts)

	return nil
}

// GetPost returns to the client a specific post based on its given id in the route
func (p *PostsHandler) GetPost(c *fiber.Ctx) error {
	// Retrive the id from route as an int
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrNotFound
	}

	// Getting the post from database
	var post data.Post
	err = p.DB.First(&post, id).Error

	// Return 404 if the post is not found
	if err != nil {
		return fiber.ErrNotFound
	}

	// Returning the post to the client as JSON
	c.JSON(post)

	return nil
}

// AddPost adds a post populated by the data in the request body to the database
func (p *PostsHandler) CreatePost(c *fiber.Ctx) error {
	// Retrieve post data from request body
	var postData data.Post
	err := c.BodyParser(&postData)
	if err != nil {
		return err
	}

	// Validate request body
	err = postData.Validate()
	if err != nil {
		return fiber.NewError(406, err.Error()) // 406 means not acceptable
	}

	// Create post with the given data in the database
	var post data.Post = data.Post{
		Title:     postData.Title,
		Content:   postData.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	p.DB.Create(&post)

	// Return to the client the id of the post
	c.JSON(post.ID)

	return nil
}

// AddPost updates a post in the database with the data in the request body and given id in its route
func (p *PostsHandler) UpdatePost(c *fiber.Ctx) error {
	// Retrive the id from route as an int
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrNotFound
	}

	// Retrieve post data from request body
	var postData data.Post
	err = c.BodyParser(&postData)
	if err != nil {
		return err
	}

	// Validate request body
	err = postData.Validate()
	if err != nil {
		return fiber.NewError(http.StatusNotAcceptable, err.Error())
	}

	// Update post with the given data in the database
	var post data.Post
	p.DB.First(&post, id)
	post.Title = postData.Title
	post.Content = postData.Content
	post.UpdatedAt = time.Now()
	p.DB.Save(&post)

	// Return to the client the id of the post
	c.JSON(post.ID)

	return nil
}

// AddPost detetes the post with the id given in its route
func (p *PostsHandler) DeletePost(c *fiber.Ctx) error {
	// Retrive the id from route as an int
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrNotFound
	}

	// Delete the post with the given id
	var post data.Post
	err = p.DB.First(&post, id).Error
	p.DB.Where("id = ?", id).Delete(&post)

	// Return 404 if the post is not found
	if err != nil {
		return fiber.ErrNotFound
	}

	return nil
}
