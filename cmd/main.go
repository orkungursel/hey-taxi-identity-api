package main

import (
	"context"
	"os"

	"github.com/orkungursel/hey-taxi-identity-api/config"
	_ "github.com/orkungursel/hey-taxi-identity-api/internal/api"
	"github.com/orkungursel/hey-taxi-identity-api/internal/server"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	_ "github.com/orkungursel/hey-taxi-identity-api/pkg/swagger"
)

// @title     Hey Taxi Identity API
// @version   1.0
// @BasePath  /api/v1

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @desc                        Add Bearer token to the request header
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := config.New()
	logger := logger.New(config)
	defer logger.Sync()

	logger.Infof("current profile: %s", config.GetProfile())

	if err := server.New(ctx, config, logger).Run(); err != nil {
		logger.Errorf("error while starting server: %s", err)
		os.Exit(1)
	}
}
