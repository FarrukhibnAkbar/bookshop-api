package models

type Category struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ParentId  string ` json:"parent_id"`
	CreatedAt string ` json:"created_at"`
	UpdatedAt string ` json:"updated_at"`
}

type ListCategories []Category

type CategoryCU struct {
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}
