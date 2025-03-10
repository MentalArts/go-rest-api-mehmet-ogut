package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	BookID     uint   `json:"book_id"`
	Rating     int    `json:"rating" gorm:"check:rating>=1 AND rating<=5"`
	Comment    string `json:"comment"`
	DatePosted string `json:"date_posted"`
}
