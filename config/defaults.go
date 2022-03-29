package config

// defaults returns config with default values.
func defaults() *Config {
	return &Config{
		App:  App{Name: "HeyTaxi Identity API"},
		Auth: Auth{DatabaseName: "auth", CollectionName: "users"},

		Server: Server{
			Http: ServerHttp{
				Host:            "",
				Port:            "8080",
				BodyLimit:       "1M",
				RequestTimeout:  10,
				ShutdownTimeout: 5,
			},
			Grpc: ServerGrpc{
				Port:                  "50051",
				MaxConnectionIdle:     10,
				Timeout:               10,
				MaxConnectionAge:      10,
				MaxConnectionAgeGrace: 10,
				Time:                  20,
			},
		},

		Jwt: Jwt{
			AccessTokenExp:             60 * 60,
			RefreshTokenExp:            60 * 60 * 24 * 15,
			Issuer:                     "hey-taxi-identity-api",
			AccessTokenPrivateKeyFile:  "/etc/certs/access-token-private-key.pem",
			AccessTokenPublicKeyFile:   "/etc/certs/access-token-public-key.pem",
			RefreshTokenPrivateKeyFile: "/etc/certs/refresh-token-private-key.pem",
			RefreshTokenPublicKeyFile:  "/etc/certs/refresh-token-public-key.pem",
		},

		Mongo: Mongo{
			Uri:               "mongodb://root:root@localhost:27017",
			ConnectionTimeout: 3,
			SocketTimeout:     3,
		},
	}
}
