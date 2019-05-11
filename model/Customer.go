package model

type Customer struct {
	Id       int    `gorm:"primary_key;column:id"`
	Uuid     string `gorm:"column:customer_uuid"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
}
