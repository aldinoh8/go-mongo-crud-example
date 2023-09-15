package repository

import (
	"context"
	"example/dto"
	"example/helpers"
	"example/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	DB         *mongo.Database
	Collection string
}

func NewUserRepository(db *mongo.Database) User {
	return User{DB: db, Collection: "users"}
}

func (u User) Register(newUser dto.RegisterBody) (model.User, error) {
	doc := model.User{
		Email:         newUser.Email,
		FullName:      newUser.FullName,
		Password:      helpers.HashPassword(newUser.Password),
		DepositAmount: 0,
		Transactions:  []model.Transaction{},
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := u.DB.Collection(u.Collection).InsertOne(ctx, doc)
	if err != nil {
		return doc, err
	}

	doc.ID = result.InsertedID.(primitive.ObjectID)
	return doc, nil
}

func (u User) FindByEmail(email string) (model.User, error) {
	doc := model.User{}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := u.DB.Collection(u.Collection)
	filter := bson.D{{"email", email}}

	err := coll.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		return doc, err
	}

	return doc, nil
}

func (u User) FindById(id string) (model.User, error) {
	doc := model.User{}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := u.DB.Collection(u.Collection)
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	err := coll.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		return doc, err
	}

	return doc, nil
}

func (u User) Transfer(sender, receiver *model.User, amount float64) error {
	coll := u.DB.Collection(u.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	senderFilter := bson.M{"_id": bson.M{"$eq": sender.ID}}
	senderChanges := bson.M{"$set": model.User{DepositAmount: sender.DepositAmount - amount}}
	newTrans := model.Transaction{ID: primitive.NewObjectID(), ReceiverID: receiver.ID, Amount: amount}
	senderTransferItem := bson.D{
		{"$push", bson.D{
			{"transactions", newTrans},
		}}}

	_, err := coll.UpdateOne(ctx, senderFilter, senderChanges)
	if err != nil {
		return err
	}

	_, err = coll.UpdateOne(ctx, senderFilter, senderTransferItem)
	if err != nil {
		return err
	}

	receiverFilter := bson.M{"_id": bson.M{"$eq": receiver.ID}}
	receiverChanges := bson.M{"$set": model.User{DepositAmount: receiver.DepositAmount + amount}}
	_, err = coll.UpdateOne(ctx, receiverFilter, receiverChanges)
	if err != nil {
		return err
	}

	return nil
}

func (u User) DeleteAccount(user *model.User) error {
	coll := u.DB.Collection(u.Collection)
	userFilter := bson.M{"_id": bson.M{"$eq": user.ID}}

	_, err := coll.DeleteOne(context.TODO(), userFilter)

	return err
}
