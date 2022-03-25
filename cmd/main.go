package main

import (
	"context"

	"github.com/orkungursel/hey-taxi-identity-api/config"
	auth "github.com/orkungursel/hey-taxi-identity-api/internal/api/http"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/mongo"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/server"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/server/swagger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := config.NewConfig()

	loggerConfig := &logger.Config{
		AppName:  config.App.Name,
		LogLevel: "debug",
		Encoder:  "console",
		DevMode:  true,
	}
	if config.IsProduction() {
		loggerConfig.LogLevel = "info"
		loggerConfig.Encoder = "json"
		loggerConfig.DevMode = false
	}

	// initialize logger
	logger := logger.New(loggerConfig)
	defer logger.Sync()

	logger.Infof("current profile: %s", config.GetProfile())

	// initilize mongo
	mongoClient, err := mongo.Connect(ctx, config)
	if err != nil {
		logger.Errorf("error while connecting to mongo: %s", err)
		return
	}
	defer mongoClient.Disconnect(ctx)

	s := server.New(ctx, config, logger)

	// initialize auth api
	auth.Api(config, s, mongoClient)

	// initialize swagger
	swagger.Api(s)

	s.Run()
}
