package models

import "time"

type Purchasing struct {
	ID         uint `gorm:"primaryKey"`
	Date       time.Time
	SupplierID uint `gorm:"not null;index"`
	UserID     uint `gorm:"not null;index"`
	GrandTotal float64
	Supplier   Supplier           `gorm:"foreignKey:SupplierID"`
	User       User               `gorm:"foreignKey:UserID"`
	Details    []PurchasingDetail `gorm:"foreignKey:PurchasingID;constraint:OnDelete:CASCADE"`
}
