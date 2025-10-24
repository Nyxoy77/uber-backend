package main

import (
	"log"
	"ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	inmem := repository.NewInmemRepository()
	svc := service.NewService(inmem)
	// fare := &domain.RideFareModel{
	// 	ID: primitive.NewObjectID(),
	// }

	router := gin.Default()

	httpHandler := http.HttpHandler{Service: svc}
	router.POST("/preview", httpHandler.HandleTripPreview)

	if err := router.Run(); err != nil {
		log.Println("error running the http server")
		return
	}
	// router.Run()

}
