package data

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Tag is the tags model
type Tag struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Posts     []*Post        `json:"-" gorm:"many2many:posts_tags;"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Validate the tag
func (t *Tag) Validate() error {
	err := validator.New().Struct(t)
	return err
}
