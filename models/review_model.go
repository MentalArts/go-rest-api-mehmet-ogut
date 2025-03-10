package models

import (
	"time"
)

// @Review Model
type Review struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	BookID     uint       `json:"book_id"`
	Rating     int        `json:"rating" gorm:"check:rating>=1 AND rating<=5"`
	Comment    string     `json:"comment"`
	DatePosted time.Time  `json:"date_posted"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
