package models

type Author struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string ` json:"updated_at"`
}

type ListAuthors []Author

type AuthorCU struct {
	Name string `json:"name"`
}
