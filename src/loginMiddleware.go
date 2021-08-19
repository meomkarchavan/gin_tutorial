package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var loginCookies = map[string]*loginCookie{}

const loginCookieName = "token"

func loginMiddleware(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, "/login") {
		return
	}

	cookieValue, err := c.Cookie(loginCookieName)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	cookie, ok := loginCookies[cookieValue]

	if !ok ||
		cookie.expiration.Unix() < time.Now().Unix() ||
		cookie.origin != c.Request.RemoteAddr {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	c.Next()
}

type loginCookie struct {
	value      string
	expiration time.Time
	origin     string
	token      string
}
