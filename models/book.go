package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        uuid.UUID `json:"id" gorm:"<-:create;primarykey;not null;unique"`
	Title     string    `json:"title" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	Year      int       `json:"year" gorm:"not null"`
	Publisher string    `json:"publisher" gorm:"not null"`
	PageCount int       `json:"page_count" gorm:"not null"`
	IsReading bool      `json:"is_reading" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}
