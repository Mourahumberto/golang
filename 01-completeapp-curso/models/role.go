package models

type Role struct {
	Id          uint        `json:"id"`
	Name        string      `json:"name"`
	Permissions []Permision `json:"permisions" gorm:"many2many:role_permissions"`
}
