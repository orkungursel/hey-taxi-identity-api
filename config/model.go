package config

type Config struct {
	App    App    `mapstructure:",squash"`
	Server Server `mapstructure:",squash"`
	Redis  Redis  `mapstructure:",squash"`
	Auth   Auth   `mapstructure:",squash"`
	Mongo  Mongo  `mapstructure:",squash"`
}

type (
	App struct {
		Name string `mapstructure:"app_name"`
	}

	Server struct {
		Host            string `mapstructure:"server_host"`
		Port            string `mapstructure:"server_port" default:"8080"`
		ShutdownTimeout int    `mapstructure:"server_shutdown_timeout"`
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
	}

	Auth struct {
		PrivateKeyFile  string `mapstructure:"auth_private_key_file"`
		PublicKeyFile   string `mapstructure:"auth_public_key_file"`
		AccessTokenExp  int    `mapstructure:"auth_access_token_exp"`
		RefreshTokenExp int    `mapstructure:"auth_refresh_token_exp"`
		DatabaseName    string `mapstructure:"auth_database_name"`
		CollectionName  string `mapstructure:"auth_collection_name"`
		Issuer          string `mapstructure:"auth_issuer"`
	}
)
