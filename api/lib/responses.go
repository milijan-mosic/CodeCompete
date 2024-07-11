package lib

import (
	gin "github.com/gin-gonic/gin"
)

func Response(data interface{}, statusCode int, errorCode, message string) gin.H {
	return gin.H{
		"status":      "SUCCESS",
		"status_code": statusCode,
		"error_code":  errorCode,
		"message":     message,
		"data":        data,
	}
}

func IntegrityResponse(message string, statusCode int, errorCode string) gin.H {
	return gin.H{
		"status":      "WARNING",
		"status_code": statusCode,
		"error_code":  errorCode,
		"message":     message,
		"data":        gin.H{},
	}
}
