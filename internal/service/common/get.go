package service

import (
	"context"

	"github.com/CriticalNoob02/store-control-be/internal/config"
)

func CommonGet(collection string, ctx context.Context, filter interface{}) ([]interface{}, error) {
	var datas []interface{}

	conn, err := config.GetDbConnection(ctx)
	if err != nil {
		return nil, err
	}
	c := conn.Database("teste").Collection(collection)

	cur, err := c.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var data interface{}
		if err := cur.Decode(&data); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	return datas, nil
}
