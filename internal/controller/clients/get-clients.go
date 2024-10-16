package controller

import (
	"context"
	"net/http"

	"github.com/CriticalNoob02/store-control-be/internal/model"
	service "github.com/CriticalNoob02/store-control-be/internal/service/common"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetClients(request *gin.Context) {
	filter := bson.M{}
	queryParams := request.Request.URL.Query()
	ctx := context.TODO()

	for _, param := range model.ClientFilterList {
		value := queryParams.Get(param)
		if value != "" {
			filter[param] = value
		}
	}

	data, err := service.CommonGet("clients", ctx, filter)
	if err != nil {
		request.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request.JSON(http.StatusOK, data)
}
