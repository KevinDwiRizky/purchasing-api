package models

type Item struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Stock int
	Price float64
}
