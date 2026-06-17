package routes

import (
	"booktracker/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/books", handlers.GetBooks)
}
