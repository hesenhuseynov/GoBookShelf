package models

type Publisher struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
}
