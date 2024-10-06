package controller

import (
	"context"
	"net/http"

	"github.com/CriticalNoob02/store-control-be/internal/model"
	"github.com/CriticalNoob02/store-control-be/internal/validation"
	"github.com/CriticalNoob02/store-control-be/pkg/service"
	"github.com/CriticalNoob02/store-control-be/pkg/util"
	"github.com/gin-gonic/gin"
)

func CreateProduct(request *gin.Context) {
	conn, err := service.GetDbConnection(context.Background())
	if err != nil {
		request.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userCollection := conn.Database("teste").Collection("products")

	var user model.Product

	if err := request.BindJSON(&user); err != nil {
		util.Logger.Error("Ocorreu um erro aqui!")
		request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validation.ProductValidation(user)
	if err != nil {
		request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		util.Logger.Error("Ocorreu um erro aqui!")
		request.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request.JSON(http.StatusCreated, gin.H{"message": "Produto criado com sucesso"})
}
