package config

// defaults returns config with default values.
func defaults() *Config {
	return &Config{
		App:    App{Name: "HeyTaxi Identity API"},
		Server: Server{Host: "", Port: "80", ShutdownTimeout: 5},
		Redis:  Redis{Addr: "localhost:6379", MaxRetries: 3},
		Mongo:  Mongo{Uri: "mongodb://root:root@localhost:27017", ConnectionTimeout: 10},
		Auth: Auth{
			AccessTokenExp:  3600,
			RefreshTokenExp: 86400,
			PrivateKeyFile:  "../certs/private.pem",
			PublicKeyFile:   "../certs/public.pem",
			DatabaseName:    "users",
		},
	}
}
