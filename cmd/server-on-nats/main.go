package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/Jeffail/gabs/v2"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

type ActorId int

const (
	NoSuchActor       ActorId = 0
	Producer          ActorId = 1
	LevelOneDispather ActorId = 2
)

func main() {
	router := gin.New()

	router.Use(GinMiddleware("http://localhost:3000"))
	router.StaticFS("/public", http.Dir("../../roots"))

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
