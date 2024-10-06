package route

import (
	"github.com/CriticalNoob02/store-control-be/internal/controller"

	"github.com/gin-gonic/gin"
)

func ClientRoutes(api *gin.Engine) {
	api.POST("/clients", controller.CreateClient)
	api.GET("/clients", controller.GetClients)
}
