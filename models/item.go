package models

type Item struct {
	ID                uint               `gorm:"primaryKey"`
	Name              string             `gorm:"not null;uniqueIndex"`
	Stock             int                `gorm:"not null"`
	Price             float64            `gorm:"not null"`
	PurchasingDetails []PurchasingDetail `gorm:"foreignKey:ItemID"`
}
