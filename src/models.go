package main

import "go.mongodb.org/mongo-driver/bson/primitive"

var users = map[string]User{
	"1": {
		// ID:        primitive.NewObjectID(),
		UserNo:    1,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Username:  "",
		Password:  "",
	},
	"2": {
		// ID:        primitive.NewObjectID(),
		UserNo:    2,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Username:  "",
		Password:  "",
	},
	"3": {
		// ID:        primitive.NewObjectID(),
		UserNo:    3,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Username:  "",
		Password:  "",
	},
	"4": {
		// ID:        primitive.NewObjectID(),
		UserNo:    4,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Username:  "",
		Password:  "",
	},
}
var test_user = User{
	// ID:        primitive.NewObjectID(),
	UserNo:    0,
	FirstName: "Omkar",
	LastName:  "Chavan",
	Username:  "omkar",
	Password:  "password",
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserNo    uint               `form:"userno"        json:"userno"`
	Email     string             `form:"email" json:"email" validate:"required,email"`
	FirstName string             `form:"firstName" json:"firstName" validate:"required"`
	LastName  string             `form:"lastName"  json:"lastName" validate:"required"`
	Username  string             `form:"username" json:"username" validate:"required"`
	Password  string             `form:"password" json:"password" validate:"required"`
}

// https://learn.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr/
