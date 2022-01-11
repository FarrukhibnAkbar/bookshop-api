package models

type Order struct {
	Id                   string   `json:"id"`
	BookId               string   `json:"book_id"`
	Description          string   `json:"description"`
	CreatedAt            string   `json:"created_at"`
	UpdatedAt            string   `json:"updated_at"`
}

type ListOrders []Order


type UpdateOrder struct {
	Id                   string   `json:"id"`
	BookId               string   `json:"book_id"`
	Description          string   `json:"description"`
}
