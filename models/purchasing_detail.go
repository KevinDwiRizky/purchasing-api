package models

type PurchasingDetail struct {
	ID           uint `gorm:"primaryKey"`
	PurchasingID uint `gorm:"not null;index"`
	ItemID       uint `gorm:"not null;index"`
	Qty          int  `gorm:"not null"`
	SubTotal     float64
	Purchasing   Purchasing `gorm:"foreignKey:PurchasingID;constraint:OnDelete:CASCADE"`
	Item         Item       `gorm:"foreignKey:ItemID"`
}
