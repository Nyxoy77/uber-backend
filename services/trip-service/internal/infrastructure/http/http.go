package http

import (
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	Service domain.TripService
}

type PreviewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (s *HttpHandler) HandleTripPreview(c *gin.Context) {
	var reqBody PreviewTripRequest
	if err := c.ShouldBindBodyWithJSON(&reqBody); err != nil {
		writeError(c, http.StatusBadRequest, "error parsing the request body")
		return
	}
	ctx := c.Request.Context()
	resp, err := s.Service.GetRoute(ctx, &reqBody.Pickup, &reqBody.Destination)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, resp)
}

func writeError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}
