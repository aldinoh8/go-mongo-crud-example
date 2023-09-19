package repository

import (
	"example/dto"
	"example/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	Create(*dto.RegisterBody) (model.User, error)
	FindByEmail(string) (model.User, error)
	FindById(primitive.ObjectID) (model.User, error)
	UpdateAmount(primitive.ObjectID, int) error
	DeleteAccount(primitive.ObjectID) error
}
