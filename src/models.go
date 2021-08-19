package main

import "go.mongodb.org/mongo-driver/bson/primitive"

var users = map[string]User{
	"1": {
		ID:        primitive.NewObjectID(),
		UserNo:    1,
		FirstName: "User1",
		LastName:  "lastname",
		Email:     "user1@mkcl.org",
		Username:  "username1",
		Password:  "password1",
	},
	"2": {
		ID:        primitive.NewObjectID(),
		UserNo:    2,
		FirstName: "User2",
		LastName:  "lastname",
		Email:     "user2@mkcl.org",
		Username:  "username2",
		Password:  "password2",
	},
	"3": {
		ID:        primitive.NewObjectID(),
		UserNo:    3,
		FirstName: "User3",
		LastName:  "lastname",
		Email:     "user3@mkcl.org",
		Username:  "username3",
		Password:  "password3",
	},
	"4": {
		ID:        primitive.NewObjectID(),
		UserNo:    4,
		FirstName: "User4",
		LastName:  "lastname",
		Email:     "user4@mkcl.org",
		Username:  "username4",
		Password:  "password4",
	},
}
var test_user = User{
	ID:        primitive.NewObjectID(),
	UserNo:    0,
	Email:     "omkarc@mkcl.org",
	FirstName: "Omkar",
	LastName:  "Chavan",
	Username:  "omkar",
	Password:  "password",
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserNo    int64              `form:"userno" json:"userno"`
	Email     string             `form:"email" json:"email" validate:"required,email"`
	FirstName string             `form:"firstName" json:"firstName" validate:"required"`
	LastName  string             `form:"lastName"  json:"lastName" validate:"required"`
	Username  string             `form:"username" json:"username" validate:"required"`
	Password  string             `form:"password" json:"password" validate:"required"`
}

// https://learn.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr/
