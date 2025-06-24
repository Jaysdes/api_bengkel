package utils

import (
	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data, // <- pastikan data dibungkus di sini
	})
}

func ResponseError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": message,
	})
}
