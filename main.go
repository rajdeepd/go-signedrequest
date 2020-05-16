package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/signedrequest", func(c *gin.Context) {
		log.Printf("%v","Inside /signedrequest GET")
		c.HTML(http.StatusOK, "signed_request_get.tmpl.html", nil)
	})

	router.POST("/api", handleVerification)

	router.POST("/login", func(c *gin.Context) {
		//What do I need to put here?
		//formContent := c.PostForm("loginForm")
		emailValue := c.PostForm("email");
		passwordValue := c.PostForm("password");
		c.HTML(http.StatusOK, "result.tmpl.html",
			gin.H{
				"email": emailValue,
				"password":passwordValue,
			})
		/*
		c.JSON(200, gin.H{
			"status":  "posted to login",
			"message": "whoo",
			"email": emailValue,
		    "password": passwordValue})*/
	})

	router.POST("/signedrequest", func(c *gin.Context) {
		//What do I need to put here?
		//formContent := c.PostForm("loginForm")
		//emailValue := c.PostForm("email")
		//passwordValue := c.PostForm("password");
		log.Printf("%v","Inside /signedrequest POST") // ""
		signedRequest := c.PostForm("signed_request")
		log.Printf("%v",signedRequest)
		c.HTML(http.StatusOK, "result_sr.tmpl.html",
			gin.H{
				"signed_request": signedRequest,
			})
		/*
			c.JSON(200, gin.H{
				"status":  "posted to login",
				"message": "whoo",
				"email": emailValue,
			    "password": passwordValue})*/
	})
	router.Run(":" + port)
}
func handleVerification(c *gin.Context) {
	if c.Request.Method == "POST" {

		c.Request.ParseForm()
		param1 := c.Request.PostForm["param1"]
		log.Printf("%v",c.Request.PostForm["param1"]) // ""
		c.HTML(http.StatusOK, "index.tmpl.html",
			gin.H{
				"param1": param1,
			})
	}
}
