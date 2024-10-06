package route

import (
	controller "github.com/CriticalNoob02/store-control-be/internal/controller/products"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(api *gin.Engine) {
	api.POST("/products", controller.CreateProduct)
	api.GET("/products", controller.GetProducts)
}
