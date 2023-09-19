package repository

import (
	"context"
	"example/dto"
	"example/helper"
	"example/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserImplementation struct {
	DB *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserImplementation {
	return UserImplementation{DB: db}
}

func (u UserImplementation) Create(newUser *dto.RegisterBody) (model.User, error) {
	doc := model.User{
		Email:         newUser.Email,
		Password:      helper.HashPassword(newUser.Password),
		FullName:      newUser.FullName,
		DepositAmount: 0,
	}

	coll := u.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return doc, err
	}

	doc.ID = result.InsertedID.(primitive.ObjectID)
	return doc, nil
}

func (u UserImplementation) FindByEmail(email string) (model.User, error) {
	coll := u.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := model.User{}

	filter := bson.M{"email": email}
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserImplementation) FindById(id primitive.ObjectID) (model.User, error) {
	coll := u.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := model.User{}
	filter := bson.M{"_id": id}

	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserImplementation) UpdateAmount(uid primitive.ObjectID, newAmount int) error {
	coll := u.DB.Collection("users")

	filter := bson.M{"_id": uid}
	update := bson.M{"$set": bson.M{"deposit_amount": newAmount}}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := coll.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (u UserImplementation) DeleteAccount(uid primitive.ObjectID) error {
	coll := u.DB.Collection("users")
	filter := bson.M{"_id": uid}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
