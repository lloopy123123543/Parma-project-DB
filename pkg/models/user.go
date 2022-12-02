package models

type User struct {
	ID       uint   `json:"ID" gorm:"primaryKey"`
	Username string `json:"Name" gorm:"not null"`
	Password string `json:"Password" gorm:"not null"`
	Email    string `json:"Email" gorm:"not null"`
}
