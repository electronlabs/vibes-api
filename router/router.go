package router

import "github.com/gin-gonic/gin"

// Start sets up the router on the specified port
func Start(port string) {
	router := gin.Default()
	router.Run(":" + port)
}
