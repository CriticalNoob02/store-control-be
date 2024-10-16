package service

import (
	"context"

	"github.com/CriticalNoob02/store-control-be/internal/config"
)

func CommonInsert(collection string, ctx context.Context, model interface{}) error {
	conn, err := config.GetDbConnection(ctx)
	if err != nil {
		return err
	}

	c := conn.Database("teste").Collection(collection)

	_, err = c.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
