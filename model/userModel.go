package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* estructura para el registro de usuarios */
type UserModel struct {
	ID        primitive.ObjectID `bson: "_id,omitempty" json:"id"`
	Name      string             `bson: "name" json:"name,omitempty"`
	Lastname  string             `bson: "lastname" json:"lastname,omitempty"`
	BirthDate time.Time          `bson: "birthDate" json:"birthDate,omitempty"`
	Email     string             `bson: "email" json:"email"`
	Password  string             `bson: "password" json:"password,omitempty"`
	Avatar    string             `bson: "avatar" json:"avatar,omitempty"`
	Banner    string             `bson: "banner" json:"banner,omitempty"`
	Biography string             `bson: "biography" json:"biography,omitempty"`
	Location  string             `bson: "location" json:"location,omitempty"`
	WebPage   string             `bson: "webPage" json:"webPage,omitempty"`
}
