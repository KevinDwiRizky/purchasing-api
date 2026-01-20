package models

type PurchasingDetail struct {
	ID           uint `gorm:"primaryKey"`
	PurchasingID uint
	ItemID       uint
	Qty          int
	SubTotal     float64
}
