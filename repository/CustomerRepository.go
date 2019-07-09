package repository

import (
	"fmt"

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
	err := database.DB.Where(model.Customer{Email: *Email}).First(&customer)
	fmt.Printf("%+v\n", customer)
	fmt.Printf("%+v\n", err)
	if !err.RecordNotFound() {
		return &customer
	}

	return nil
}

func Create(Customer *model.Customer) bool {
	err := database.DB.Create(&Customer)
	if err != nil {
		return false
	}
	return true
}

// func getID() {
// 	conn := <-DBTemp
// 	defer DBTemp -< conn

//  // Do your fucking Database Manipulation
// }
