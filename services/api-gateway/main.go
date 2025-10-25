package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ride-sharing/services/api-gateway/routes"
	ws "ride-sharing/services/api-gateway/websockets"
	"ride-sharing/shared/env"

	"github.com/gin-contrib/cors"
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
	// Cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Content-Type", "Authorization")
	router.Use(cors.New(config))

	router.POST("/trip/preview", routes.TripPreviewHandler)
	router.GET("/ws/riders", ws.HandleRiderWebSocket)
	router.GET("/ws/drivers", ws.HandleDriverWebSocket)
	// router.GET()

	server := &http.Server{
		Addr:    httpAddr,
		Handler: router,
	}

	serverErrorChan := make(chan error, 1)
	signalChan := make(chan os.Signal, 1)
	go func() {
		serverErrorChan <- server.ListenAndServe()
	}()

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrorChan:
		log.Printf("error occured starting the server %v :", err)
	case sig := <-signalChan:
		log.Printf("recieved signal : %v ", sig)
		log.Print("shutting down gracefully")

		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("cannot gracefull shutdown error : %v ", err)
			server.Close()
		}
	}
	log.Println("gracefull shutdown completed")
}
