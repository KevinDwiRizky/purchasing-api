package models

type Supplier struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Email       string `gorm:"not null;uniqueIndex"`
	Address     string
	Purchasings []Purchasing `gorm:"foreignKey:SupplierID"`
}
