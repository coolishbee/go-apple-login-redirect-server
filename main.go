package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/callback?"+c.Request.RequestURI)
	})

	router.GET("/callback", func(c *gin.Context) {
		log.Println(c.Request.RequestURI)
	})

	router.GET("/test", func(c *gin.Context) {
		example := c.Accepted
		// it would print: "12345"
		log.Println(example)
		log.Println("test log")
		log.Println(c.Params)
	})

	router.GET("/auth/callback", func(c *gin.Context) {
		code := c.Param("code")
		log.Println(code)
		log.Println(c.Params)
	})

	router.Run("192.168.7.2:3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//router.RunTLS(":8080", "server.crt", "server.key")
}
