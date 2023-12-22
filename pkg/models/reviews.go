package models

import "time"

type Review struct {
	ID         int       `json:"id"`
	BookID     int       `json:"book_id"`
	UserID     int       `json:"user_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment,omitempty"`
	ReviewDate time.Time `json:"review_date"`
}
