package models

import (
	"time"
)

// @Book Model
type Book struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	Title           string     `json:"title"`
	AuthorID        uint       `json:"author_id"`
	Author          Author     `json:"author" gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE;"`
	ISBN            string     `json:"isbn" gorm:"unique"`
	PublicationYear int        `json:"publication_year"`
	Description     string     `json:"description"`
	Reviews         []Review   `json:"reviews" gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}
