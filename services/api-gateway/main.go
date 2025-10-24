package main

import (
	"fmt"
	"log"
	"net/http"

	"ride-sharing/services/api-gateway/routes"
	"ride-sharing/shared/env"

	"github.com/gin-gonic/gin"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")
	router := gin.Default()
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "method not allowed",
		})
	})
	router.POST("/trip/preview", routes.TripPreviewHandler)
	// router.GET()
	fmt.Println("Hello")
	if err := router.Run(httpAddr); err != nil {
		log.Fatalf("error starting the server %v", err)
	}
}
