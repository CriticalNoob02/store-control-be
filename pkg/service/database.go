package database

import (
	"context"

	"github.com/brenofelips/api-meu-deputado/internal/config"
	"github.com/brenofelips/api-meu-deputado/pkg/util"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetDbConnection(ctx context.Context) (*mongo.Client, error) {
	uri := config.GetDbDsn()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		util.Logger.Error(err, "database", uri)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		util.Logger.Error(err, "database", uri)
		return nil, err
	}

	util.Logger.Info("Conectado!", "database", uri)
	return client, nil
}
