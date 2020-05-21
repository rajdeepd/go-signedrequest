package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")
	secret := os.Getenv("SECRET")

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
		emailValue := c.PostForm("email");
		passwordValue := c.PostForm("password");
		c.HTML(http.StatusOK, "result.tmpl.html",
			gin.H{
				"email": emailValue,
				"password":passwordValue,
			})
	})

	router.POST("/signedrequest", func(c *gin.Context) {

		log.Printf("%v","Inside /signedrequest POST") // ""
		signedRequest := c.PostForm("signed_request")
		log.Printf("%v",signedRequest)
		splitSR := strings.Split(signedRequest, ".")
		encodedSig := splitSR[0]

		encodedEnvelope := splitSR[1]

		log.Printf("%v","Encoded Signature:" + encodedSig)
		log.Printf("%v","Encoded Envelope:" + splitSR[1])

		validDigest := validateDigest(secret, encodedSig, encodedEnvelope)

		jsonEnvelope, _ := base64.StdEncoding.DecodeString(encodedEnvelope)
		log.Printf("base64: %s\n", jsonEnvelope)
		log.Printf("Valid Digest: %s\n", validDigest)
		signedRequestStruct := SignedRequestStruct{}

		json.Unmarshal([]byte(jsonEnvelope), &signedRequestStruct)
		algo := signedRequestStruct.Algorithm
		oauthToken :=  signedRequestStruct.Client.OauthToken
		log.Printf("Algorithm:" + algo)
		log.Printf("OAuth Token:" + oauthToken)
		c.HTML(http.StatusOK, "result_sr.tmpl.html",
			gin.H{
				//"signed_request": signedRequest,
			    "algo": algo,
			    "oauth_token" : oauthToken,
			    "valid_digest" : validDigest,
			})
	})
	router.Run(":" + port)
}

func validateDigest(secret string, encodedSign string, encodedEnvelope string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	h.Write([]byte(encodedEnvelope))
	fmt.Println(len([]byte(encodedEnvelope)))

	sha := base64.StdEncoding.EncodeToString(h.Sum(nil))
	log.Printf("Expected signature:", encodedSign)
	log.Printf("Generated signature:", sha)
	return sha == encodedSign

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
