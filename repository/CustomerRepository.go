package repository

import (
	"../database"
	"../model"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

// func restoreConnection() {
// 	database.DB <- conn
// }

func FindByEmail(Email *string) *model.Customer {
	var customer model.Customer
	// conn = <-database.DB
	// defer func() {
	// 	go restoreConnection()
	// }()
	err := database.DB.Where(&model.Customer{Email: *Email}).First(&customer)
	if err == nil {
		return &customer
	}

	return nil
}

// func getID() {
// 	conn := <-DBTemp
// 	defer DBTemp -< conn

//  // Do your fucking Database Manipulation
// }
