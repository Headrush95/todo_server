package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type respError struct {
	Message string `json:"message"`
}

// newErrorResponse помещает ошибку в ответ и ...
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, respError{message})
}
