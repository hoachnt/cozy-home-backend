package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}
type statusResopnse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Warn(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
