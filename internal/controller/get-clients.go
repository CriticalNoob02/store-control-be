package controller

import (
	"context"
	"net/http"

	"github.com/CriticalNoob02/store-control-be/internal/validation"
	"github.com/CriticalNoob02/store-control-be/pkg/service"
	"github.com/CriticalNoob02/store-control-be/pkg/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetClients(request *gin.Context) {
	filter := bson.M{}
	queryParams := request.Request.URL.Query()
	var decodedClients []map[string]interface{}

	conn, err := service.GetDbConnection(context.Background())
	if err != nil {
		request.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userCollection := conn.Database("teste").Collection("clients")

	for _, param := range validation.ClientsFilterList {
		value := queryParams.Get(param)
		if value != "" {
			filter[param] = value
		}
	}

	cur, err := userCollection.Find(context.Background(), filter)
	if err != nil {
		util.Logger.Error("Erro ao buscar dados no banco", "error", err.Error())
		request.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var decodedClient map[string]interface{}
		if err := cur.Decode(&decodedClient); err != nil {
			util.Logger.Error("Erro ao decodificar o resultado", "error", err.Error())
			request.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar cliente"})
			return
		}
		decodedClients = append(decodedClients, decodedClient)
	}

	request.JSON(http.StatusOK, decodedClients)
}
