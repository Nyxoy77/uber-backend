package main

import (
	"context"
	"fmt"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	inmem := repository.NewInmemRepository()
	svc := service.NewService(inmem)
	fare := &domain.RideFareModel{
		ID: primitive.NewObjectID(),
	}
	_, err := svc.CreateTrip(context.TODO(), fare)
	if err != nil {
		fmt.Println("An error occured ", err)
		return
	}
	fmt.Println("Trip created with ID : ", fare.ID)

	for {
		// fmt.Println("Listening to the server ")
		time.Sleep(5 * time.Minute)
	}
}
