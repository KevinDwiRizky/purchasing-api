package models

type Supplier struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Email   string
	Address string
}
