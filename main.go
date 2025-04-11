package main

import (
	"go-jwt/initializers"
	"go-jwt/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	// Inject Services
	routes.SetupRoutes(router)

	router.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.212"})
	router.Run()
}
