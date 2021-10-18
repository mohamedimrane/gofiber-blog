package data

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Tag struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (t *Tag) Validate() error {
	err := validator.New().Struct(t)
	return err
}
