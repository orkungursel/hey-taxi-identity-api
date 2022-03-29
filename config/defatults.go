package config

// defaults returns config with default values.
func defaults() *Config {
	return &Config{
		App: App{Name: "HeyTaxi Identity API"},
		Mongo: Mongo{Uri: "mongodb://root:root@localhost:27017", ConnectionTimeout: 3, SocketTimeout: 3},
		Server: Server{
			Host: "", Port: "8080", ShutdownTimeout: 5,
			Grpc: Grpc{Port: "50051", MaxConnectionIdle: 10, Timeout: 10, MaxConnectionAge: 10, MaxConnectionAgeGrace: 10, Time: 20},
		},
		Redis: Redis{Addr: "localhost:6379", MaxRetries: 3},
		Mongo: Mongo{Uri: "mongodb://root:root@localhost:27017", ConnectionTimeout: 10},
		Auth:  Auth{DatabaseName: "auth", CollectionName: "users"},
		Jwt: Jwt{Issuer: "hey-taxi-identity-api", AccessTokenExp: 20, RefreshTokenExp: 86400,
			AccessTokenPrivateKeyFile:  "/etc/certs/access-token-private-key.pem",
			AccessTokenPublicKeyFile:   "/etc/certs/access-token-public-key.pem",
			RefreshTokenPrivateKeyFile: "/etc/certs/refresh-token-private-key.pem",
			RefreshTokenPublicKeyFile:  "/etc/certs/refresh-token-public-key.pem",
		},
	}
}
