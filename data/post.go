package data

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `json:"title" validate:"required,gte=3,lte=300"`
	Content   string         `json:"content" validate:"required,gte=255,lte=10000"`
	Tags      []*Tag         `json:"-" gorm:"many2many:posts_tags;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Post) Validate() error {
	err := validator.New().Struct(p)
	return err
}
