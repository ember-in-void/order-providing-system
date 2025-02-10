package model

type MenuItem struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Allergens   []string `json:"allergens"`
	Size        string   `json:"size"`
}
