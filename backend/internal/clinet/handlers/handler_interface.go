package handlers

import "github.com/gin-gonic/gin"

type ClinetHandlersInterface interface {
	CreateClient(*gin.Context)
}
