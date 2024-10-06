package service

import (
	"context"

	"github.com/CriticalNoob02/store-control-be/internal/config"
	"github.com/CriticalNoob02/store-control-be/pkg/util"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetDbConnection(ctx context.Context) (*mongo.Client, error) {
	uri := config.GetDbUri()

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
	return client, nil
}
