package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(loginMiddleware)
	r.LoadHTMLGlob("E:\\Omkar\\Coding\\go\\src\\gin_tutorial\\templates\\*html")
	// r.LoadHTMLGlob("D:\\GO_Workspace\\src\\day7\\gin\\hello_world\\templates\\*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	test := r.Group("/test")
	test.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "group-test.html", nil)
	})

	admin := r.Group("/admin")
	admin.GET("/add/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			c.HTML(http.StatusOK, "admin-add.html", nil)
		}
		user, ok := users[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - Not Found")
			return
		}
		c.HTML(http.StatusOK, "admin-add.html", user)
		// c.IndentedJSON(http.StatusOK, user)
	})
	admin.POST("/addUser", func(c *gin.Context) {
		var user User
		err := c.Bind(&user)
		users[strconv.FormatUint(uint64(user.UserNo), 10)] = user
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(
			http.StatusOK,
			users,
		)
	})
	r.POST("/signup", func(c *gin.Context) {
		var user User

		v := validator.New()

		err := c.Bind(&user)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		err = v.Struct(user)

		for _, e := range err.(validator.ValidationErrors) {
			c.JSON(
				http.StatusOK,
				gin.H{"error": "invaid data" + e.Field()},
			)
			return
		}
		// user.UserNo = counter
		// counter += 1

		result := create_user(user)

		c.JSON(
			http.StatusOK,
			result.InsertedID,
		)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", Login)
	r.Static("/public", "D:\\GO_Workspace\\src\\day7\\gin\\hello_world\\public")
	return r
}
