package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FullName      string             `bson:"full_name,omitempty" json:"full_name"`
	Email         string             `bson:"email,omitempty" json:"email"`
	Password      string             `bson:"password,omitempty" json:"-"`
	DepositAmount int                `bson:"deposit_amount" json:"deposit_amount"`
}
