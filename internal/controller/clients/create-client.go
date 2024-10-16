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

func CreateClient(request *gin.Context) {
	ctx := context.TODO()
	var user model.Client

	if err := request.BindJSON(&user); err != nil {
		util.Logger.Error("Ocorreu um erro aqui!")
		request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.IdClient = uuid.New().String()

	if err := util.ValidateStruct(user); err != nil {
		request.JSON(http.StatusBadRequest, gin.H{"validation": err.Error()})
		return
	}

	err := service.CommonInsert("clients", ctx, user)
	if err != nil {
		request.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}
