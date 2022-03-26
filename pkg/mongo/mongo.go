package mongo

import (
	"context"
	"time"

	"github.com/orkungursel/hey-taxi-identity-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(ctx context.Context, config *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(config.Mongo.ConnectionTimeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.Uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
