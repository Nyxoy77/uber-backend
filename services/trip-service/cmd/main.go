package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"syscall"
	"time"

	stdhttp "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	inmem := repository.NewInmemRepository()
	svc := service.NewService(inmem)
	// fare := &domain.RideFareModel{
	// 	ID: primitive.NewObjectID(),
	// }

	router := gin.Default() // This is a handler which can be passed to an actual http server

	httpHandler := http.HttpHandler{Service: svc}
	router.POST("/preview", httpHandler.HandleTripPreview)

	server := &stdhttp.Server{
		Addr:    ":8083",
		Handler: router,
	}
	serverChannel := make(chan error, 1)

	go func() {
		serverChannel <- server.ListenAndServe()
	}()

	sigChannel := make(chan os.Signal, 1)

	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverChannel:
		log.Printf("Error starting the server: %v", err)

	case sig := <-sigChannel:
		log.Printf("Server gracefully shutting down because of signal %v ", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not stop the server gracefully %v ", err)
			server.Close()
		}

	}

	log.Println("gracefully shutting down the server")

}
