package routes

import (
	"log"
	"net/http"
	"ride-sharing/services/api-gateway/models"
	"ride-sharing/services/api-gateway/writer"
	"ride-sharing/shared/contracts"

	"github.com/gin-gonic/gin"
)

func TripPreviewHandler(c *gin.Context) {
	var reqBody models.PreviewTripRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error parsing the request",
		})
		log.Printf("error parsing the request %s", err)
		return
	}

	if reqBody.UserID == "" {
		writer.WriteError(c, http.StatusBadRequest, "UserID is required")
		return
	}
	obj := contracts.APIResponse{Data: "ok"}
	writer.WriteSuccess(c, 200, obj)
}

