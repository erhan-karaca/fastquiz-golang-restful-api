package entities

type Type struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
