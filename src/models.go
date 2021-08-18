package main

var users = map[string]User{
	"0": {
		ID:        "1",
		FirstName: "Jennifer",
		LastName:  "Watson",
	},
	"1": {
		ID:        "1",
		FirstName: "Jennifer",
		LastName:  "Watson",
	},
	"2": {
		ID:        "2",
		FirstName: "Jennifer",
		LastName:  "Watson",
	},
	"3": {
		ID:        "3",
		FirstName: "Jennifer",
		LastName:  "Watson",
	},
	"4": {
		ID:        "4",
		FirstName: "Jennifer",
		LastName:  "Watson",
	},
}

type User struct {
	ID        string `form:"id"        json:"id"`
	FirstName string `form:"firstName" json:"firstName"`
	LastName  string `form:"lastName"  json:"lastName"`
}

type AuthUser struct {
	ID uint32 `json:"id"`
}

// https://learn.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr/
