package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sharjil07/service-request/controller"
	"github.com/sharjil07/service-request/database"
)

func main() {
	router := gin.Default()
	database.ConnectDB()

	router.POST("/identify", controllers.Identify)

	router.Run(":8080")
}
