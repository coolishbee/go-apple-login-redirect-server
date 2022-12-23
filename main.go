package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	router := gin.Default()
	ctx := context.Background()

	oauth2Conf := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8077/redirect",
		Scopes: []string{
			"https://www.googleapis.com/auth/admob.readonly",
		},
	}

	router.GET("/redirect", func(c *gin.Context) {
		code := c.Query("code")
		log.Println(code)
		tok, err := oauth2Conf.Exchange(ctx, code)
		if err != nil {
			c.JSON(500, err)
			return
		}
		log.Println(tok.RefreshToken)
	})

	router.GET("/apple/redirect", func(c *gin.Context) {
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

		code := c.Param("code")
		log.Println(code)
		log.Println(c.Params)
	})

	router.Run(":8077") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//router.RunTLS(":8080", "server.crt", "server.key")
}
