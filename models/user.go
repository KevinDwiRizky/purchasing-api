package models

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	Role        string
	Purchasings []Purchasing `gorm:"foreignKey:UserID"`
}
