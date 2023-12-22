package models

type OrderDetail struct {
	OrderID  int     `json:"order_id"`
	BookID   int     `json:"book_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
