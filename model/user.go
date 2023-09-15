package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FullName      string             `bson:"full_name,omitempty" json:"full_name"`
	Email         string             `bson:"email,omitempty" json:"email"`
	Password      string             `bson:"password,omitempty" json:"-"`
	DepositAmount float64            `bson:"deposit_amount,omitempty" json:"deposit_amount"`
	Transactions  []Transaction      `bson:"transactions,omitempty" json:"transactions"`
}

type Transaction struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	ReceiverID primitive.ObjectID `bson:"receiver_id" json:"receiver_id"`
	Amount     float64            `bson:"amount" json:"amount"`
}
