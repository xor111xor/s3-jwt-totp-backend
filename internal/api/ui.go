package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// user inteface for interaction api
func startUI(service *gin.Engine, ui_group string) {
	// load static html
	service.LoadHTMLFiles("templates/ui/signup.html", "templates/ui/signin.html", "templates/ui/logout.html",
		"templates/ui/list.html", "templates/ui/upload.html")

	// create group
	router := service.Group(ui_group)

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"title": "Sign Up",
		})
	})
	router.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.html", gin.H{
			"title": "Sign In",
		})
	})
	router.GET("/logout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "logout.html", gin.H{
			"title": "Logout",
		})
	})

	router.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{
			"title": "List files",
		})
	})

	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"title": "Upload",
		})
	})
}
