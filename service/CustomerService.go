package service

import (
	"../model"
	"../repository"
)

func FindByEmail(Email *string) *model.Customer {
	return repository.FindByEmail(Email)
}
