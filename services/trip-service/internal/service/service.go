package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {

	trip := &domain.TripModel{
		Id:       primitive.NewObjectID(),
		UserID:   "",
		Status:   "pending",
		RideFare: fare,
	}
	return s.repo.CreateTrip(ctx, trip)
}

func (s *service) GetRoute(ctx context.Context, pickup, destination *types.Coordinate) (*types.OsrmApiResponse, error) {
	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full&geometries=geojson", pickup.Longitude, pickup.Latitude, destination.Longitude, destination.Latitude)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching the OSRM api: %w", err)
	}
	defer resp.Body.Close()
	bytes, er := io.ReadAll(resp.Body)
	log.Println(string(bytes))
	if er != nil {
		return nil, fmt.Errorf("error reading the response: %w", er)
	}

	var response *types.OsrmApiResponse

	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling the response: %w", err)
	}
	return response, nil
}
