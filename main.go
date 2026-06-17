package main

import (
	"api-booktracker/config"
	"api-booktracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")

}
