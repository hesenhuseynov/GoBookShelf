package models

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title" validate:"required"`
	AuthorName  string  `json:"author_name"`
	ISBN        string  `json:"isbn" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gt=0"`
	Description string  `json:"description,omitempty"`
	PublisherID *int    `json:"publisher_id,omitempty"`
	CategoryID  *int    `json:"category_id,omitempty"`
}
