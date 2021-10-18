package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// TagsHandler is a struct managing tags related handlers dependencies ex: database
type TagsHandler struct {
	DB *gorm.DB
}

// NewPostsHandler is the constuctor of handlers.PostsHandler
func NewTagsHandler(db *gorm.DB) *TagsHandler {
	return &TagsHandler{
		DB: db,
	}
}

func (t *TagsHandler) GetTags(c *fiber.Ctx) error {
	return nil
}
