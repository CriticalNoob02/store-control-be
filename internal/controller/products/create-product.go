package controller

import (
	"context"
	"net/http"

	"github.com/CriticalNoob02/store-control-be/internal/model"
	service "github.com/CriticalNoob02/store-control-be/internal/service/common"
	"github.com/CriticalNoob02/store-control-be/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProduct(request *gin.Context) {
	var product model.Product
	ctx := context.TODO()

	if err := request.BindJSON(&product); err != nil {
		request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.IdProduct = uuid.New().String()

	if err := util.ValidateStruct(product); err != nil {
		request.JSON(http.StatusBadRequest, gin.H{"validation": err.Error()})
		return
	}

	err := service.CommonInsert("products", ctx, product)
	if err != nil {
		request.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request.JSON(http.StatusCreated, gin.H{"message": "Produto criado com sucesso"})
}
