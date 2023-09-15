package repository

import (
	"example/dto"
	"example/model"
)

type UserInterface interface {
	Create(dto.RegisterBody) (model.User, error)
	FindByEmail(string) (model.User, error)
	FindById(string) (model.User, error)
	Transfer(*model.User, *model.User, float64) error
	DeleteAccount(*model.User) error
}
