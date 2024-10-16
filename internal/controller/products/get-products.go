package controller

import (
	"context"
	"net/http"

	"github.com/CriticalNoob02/store-control-be/internal/model"
	service "github.com/CriticalNoob02/store-control-be/internal/service/common"
	"github.com/CriticalNoob02/store-control-be/pkg/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProducts(request *gin.Context) {
	filter := bson.M{}
	queryParams := request.Request.URL.Query()
	ctx := context.TODO()

	for _, param := range model.ProductFilterList {
		value := queryParams.Get(param)
		if value != "" {
			filter[param] = value
		}
	}

	data, err := service.CommonGet("products", ctx, filter)
	if err != nil {
		util.Logger.Error("Erro ao buscar dados no banco", "error", err.Error())
		request.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request.JSON(http.StatusOK, data)
}
