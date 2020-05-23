package endpoints

import "github.com/gin-gonic/gin"

func PongEndpoint(c *gin.Context) {
	// Just return 200 to indicate the service is functional
	c.JSON(200, gin.H{"message": "pong"})
}
