package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func loginMiddleware(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, "/login") ||
		strings.HasPrefix(c.Request.URL.Path, "/public") {
		return
	}

	token := c.GetHeader("Authorization")
	token = strings.Split(token, "Bearer ")[1]

	claims := jwt.MapClaims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	if !tkn.Valid {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	data := claims["user_id"]

	loginToken, ok := loginTokens[data.(string)]

	if !ok {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	if loginToken.access_token != token {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.Next()
}

type loginToken struct {
	userId        string
	refresh_token string
	access_token  string
}
