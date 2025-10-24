package writer

import (
	"github.com/gin-gonic/gin"
)

func WriteError(c *gin.Context, status int, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, gin.H{
		"error": message,
	})
}

func WriteSuccess(c *gin.Context, status int, data any) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, data)
}
