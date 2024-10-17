package service

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EchoService returns all the request payload including body
func EchoService(c *gin.Context) {
	// Read the request body using a more efficient method
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error",
			"code":   http.StatusBadRequest,
			"error":  "Failed to read payload",
		})
		return
	}

	// Send response with the request body and status
	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"code":   http.StatusOK,
		"body":   string(body), // Body is returned as a string
	})
}
