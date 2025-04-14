package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseOptions struct {
	Status  int
	Message string
}

// SuccessResponse is a helper function to standardize API responses
func SuccessResponse(c *gin.Context, data interface{}, opts ...ResponseOptions) {
	status, message := applyResponseOptionsIfExists(
		http.StatusOK,
		"success",
		opts...,
	)

	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse is a helper function for error responses
func ErrorResponse(c *gin.Context, err error, opts ...ResponseOptions) {
	status, message := applyResponseOptionsIfExists(
		http.StatusInternalServerError,
		"error",
		opts...,
	)

	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"error":   err.Error(),
	})
}

// applyResponseOptionsIfExists applies the provided response options or defaults to the given values
func applyResponseOptionsIfExists(defaultStatus int, defaultMessage string, opts ...ResponseOptions) (int, string) {
	if len(opts) <= 0 {
		return defaultStatus, defaultMessage
	}

	status := defaultStatus
	message := defaultMessage

	if opts[0].Status != 0 {
		status = opts[0].Status
	}

	if opts[0].Message != "" {
		message = opts[0].Message
	}

	return status, message
}
