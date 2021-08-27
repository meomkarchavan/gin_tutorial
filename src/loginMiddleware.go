package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var loginCookies = map[string]*loginCookie{}

const loginCookieName = "token"

func loginMiddleware(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, "/login") ||
		strings.HasPrefix(c.Request.URL.Path, "/public") {
		return
	}

	// token, err := c.Cookie(loginCookieName)
	// if err != nil {
	// 	log.Println("No Cookie")
	// 	c.Redirect(http.StatusTemporaryRedirect, "/login")
	// 	return
	// }

	// cookie, ok := loginCookies[token]

	// if !ok ||
	// 	cookie.expiration.Unix() < time.Now().Unix() ||
	// 	cookie.origin != c.Request.RemoteAddr {
	// 	log.Println("not ok")
	// 	log.Println(cookie)
	// 	c.Redirect(http.StatusTemporaryRedirect, "/login")
	// }
	// // Initialize a new instance of `Claims`
	token := c.GetHeader("auth")
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
	// log.Println("**********************")
	// log.Println(claims)
	c.Next()
}

type loginCookie struct {
	value      string
	expiration time.Time
	origin     string
	token      string
}
