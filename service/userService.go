package service

import (
	"context"
	"time"

	"github.com/darias-developer/challenge-go/config"
	"github.com/darias-developer/challenge-go/middleware"
	"github.com/darias-developer/challenge-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* FindUserByEmail busca usuario en la db por medio del email */
func FindUserByEmail(email string) (model.UserModel, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := config.ConectDB().Database("curso_go")
	collection := db.Collection("user")

	condition := bson.M{"email": email}

	var userModel model.UserModel

	err := collection.FindOne(ctx, condition).Decode(&userModel)
	ID := userModel.ID.Hex()

	if err != nil {
		return userModel, false, ID
	}

	return userModel, true, ID
}

/* CreateUser crea usuario en la db */
func CreateUser(userModel model.UserModel) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := config.ConectDB().Database("curso_go")
	collection := db.Collection("user")

	encryptedPassword, err := middleware.EncryptPassword(userModel.Password)

	if err != nil {
		return "", false, err
	}

	userModel.Password = encryptedPassword

	result, err := collection.InsertOne(ctx, userModel)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
