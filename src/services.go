package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func CreateToken(user User) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "omkar")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.UserNo
	atClaims["username"] = user.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//TODO
	// validate username password
	//https://github.com/go-validator/validator
	v := validator.New()
	err := v.Struct(user)

	for _, e := range err.(validator.ValidationErrors) {
		c.JSON(
			http.StatusOK,
			gin.H{"error": "Invaid data" + e.Field()},
		)
		return
	}

	result, err := find_user(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(result)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	lc := loginCookie{
		value:      strconv.FormatUint(uint64(result.UserNo), 10),
		expiration: time.Now().Add(24 * time.Hour),
		origin:     c.Request.RemoteAddr,
		token:      token,
	}

	loginCookies[lc.value] = &lc
	c.SetCookie(loginCookieName, lc.token, 10*60, "", "localhost", false, true)

	c.JSON(http.StatusOK, token)
}
