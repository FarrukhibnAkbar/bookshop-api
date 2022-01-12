package models

type Book struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	AuthorId   Author     `json:"author_id"`
	CategoryId []string   `json:"category_id"`
	Categories []Category `json:"categories"`
	CreatedAt  string     `json:"created_at"`
	UpdateAt   string     `json:"update_at"`
}

type ListBooks []Book

type BookCU struct {
	Name       string   `json:"name"`
	AuthorId   string   `json:"author_id"`
	CategoryId []string `json:"category_id"`
}
