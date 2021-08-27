package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
func CheckToken(token string, c *gin.Context) error {
	log.Println("here")
	log.Println(token)
	// // Initialize a new instance of `Claims`
	claims := &jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		var jwtKey = []byte(os.Getenv("ACCESS_SECRET"))
		return jwtKey, nil
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

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	log.Println(http.StatusOK, []byte(fmt.Sprintf("Welcome %s!", claims)))
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
	log.Println(result.Username, result.Password)
	log.Println(user.Username, user.Password)
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
	// lc := loginCookie{
	// 	value:      token,
	// 	expiration: time.Now().Add(24 * time.Hour),
	// 	origin:     c.Request.RemoteAddr,
	// }

	// loginCookies[lc.value] = &lc
	// c.SetCookie(loginCookieName, lc.value, 10*60, "", "localhost", false, true)
	c.Header("auth", token)
	log.Println(loginCookies)
	c.JSON(http.StatusOK, token)
}
