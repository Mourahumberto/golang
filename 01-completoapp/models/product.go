package models

type Product struct {
	Id    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
