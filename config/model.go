package config

type Config struct {
	App      App      `mapstructure:",squash"`
	Server   Server   `mapstructure:",squash"`
	Redis    Redis    `mapstructure:",squash"`
	Mongo    Mongo    `mapstructure:",squash"`
	Auth     Auth     `mapstructure:",squash"`
	Jwt      Jwt      `mapstructure:",squash"`
	Password Password `mapstructure:",squash"`
}

type (
	App struct {
		Name string `mapstructure:"app_name"`
	}

	Server struct {
		Host            string `mapstructure:"server_host"`
		Port            string `mapstructure:"server_port" default:"8080"`
		RequestTimeout  int    `mapstructure:"server_request_timeout"`
		ShutdownTimeout int    `mapstructure:"server_shutdown_timeout"`
		Grpc            Grpc   `mapstructure:"-"`
	}

	Grpc struct {
		Host                  string `mapstructure:"server_grpc_host"`
		Port                  string `mapstructure:"server_grpc_port" default:"50051"`
		MaxConnectionIdle     int    `mapstructure:"server_grpc_max_connection_idle"`
		Timeout               int    `mapstructure:"server_grpc_timeout"`
		MaxConnectionAge      int    `mapstructure:"server_grpc_max_connection_age"`
		MaxConnectionAgeGrace int    `mapstructure:"server_grpc_max_connection_age_grace"`
		Time                  int    `mapstructure:"server_grpc_time"`
	}

	Redis struct {
		Addr         string `mapstructure:"redis_addr"`
		Password     string `mapstructure:"redis_password"`
		DB           int    `mapstructure:"redis_db"`
		DefaultDb    string `mapstructure:"redis_defaultdb"`
		MinIdleConns int    `mapstructure:"redis_min_idle_conns"`
		PoolSize     int    `mapstructure:"redis_pool_size"`
		PoolTimeout  int    `mapstructure:"redis_pool_timeout"`
		MaxRetries   int    `mapstructure:"redis_max_retries"`
	}

	Mongo struct {
		Uri               string `mapstructure:"mongo_uri"`
		ConnectionTimeout int    `mapstructure:"mongo_connection_timeout"`
		SocketTimeout     int    `mapstructure:"mongo_socket_timeout"`
	}

	Auth struct {
		DatabaseName   string `mapstructure:"auth_database_name"`
		CollectionName string `mapstructure:"auth_collection_name"`
	}

	Jwt struct {
		AccessTokenPrivateKeyFile  string `mapstructure:"jwt_access_token_private_key_file"`
		AccessTokenPublicKeyFile   string `mapstructure:"jwt_access_token_public_key_file"`
		AccessTokenExp             int    `mapstructure:"jwt_access_token_exp"`
		RefreshTokenPrivateKeyFile string `mapstructure:"jwt_refresh_token_private_key_file"`
		RefreshTokenPublicKeyFile  string `mapstructure:"jwt_refresh_token_public_key_file"`
		RefreshTokenExp            int    `mapstructure:"jwt_refresh_token_exp"`
		Issuer                     string `mapstructure:"jwt_issuer"`
	}

	Password struct {
		HashSecret string `mapstructure:"password_hash_secret"`
	}
)
