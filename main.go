package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LOGIN struct {
	USER     string `json:"user" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.String(200, "on")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.POST("/foo", func(c *gin.Context) {
		var login LOGIN
		c.BindJSON(&login)
		c.JSON(200, gin.H{"status": login.USER}) // Your custom response here
	})
	r.Run(":8080") // listen for incoming connections
}
