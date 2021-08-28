package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var loginTokens = map[string]*loginToken{}

func CreateToken(user User) (map[string]string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = strconv.FormatUint(uint64(user.UserNo), 10)
	atClaims["username"] = user.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	// refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  token,
		"refresh_token": rt,
	}, nil

}
func CheckToken(token string, c *gin.Context) error {
	// // Initialize a new instance of `Claims`
	claims := &jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// c.JSON(http.StatusUnauthorized, "StatusUnauthorized")
			return errors.New("StatusUnauthorized")
		}
		// c.JSON(http.StatusBadRequest, "StatusBadRequest")
		return errors.New("StatusUnauthorized")
	}
	if !tkn.Valid {
		// c.JSON(http.StatusUnauthorized, "StatusUnauthorized")
		return errors.New("StatusUnauthorized")
	}

	return nil
}
func Login(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//TODO
	// validate username password
	//https://github.com/go-validator/validator
	// v := validator.New()
	// err := v.Struct(user)

	// for _, e := range err.(validator.ValidationErrors) {
	// 	c.JSON(
	// 		http.StatusOK,
	// 		gin.H{"error": "Invaid data" + e.Field()},
	// 	)
	// 	return
	// }

	result, err := find_user(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No user found")
		return
	}

	if result.Password != user.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(result)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	lt := loginToken{
		access_token:  token["access_token"],
		refresh_token: token["refresh_token"],

		userId: strconv.FormatUint(uint64(result.UserNo), 10),
	}
	loginTokens[strconv.FormatUint(uint64(result.UserNo), 10)] = &lt
	c.Header("Authorization", token["access_token"])
	c.Header("refresh_token", token["refresh_token"])
	c.JSON(http.StatusOK, token)
}
