package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/stclaird/questionGenerator/pkg/question"
)

func main() {
    //get the app configuration
    config := GetConfig()

    router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "status": "OK",
		})
	  })

    router.Use(cors.New(CORSConfig()))
    question.RegisterRoutes(router)
    router.Run(config.port)
}
