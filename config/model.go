package config

type Config struct {
	App    App
	Server Server
	Redis  Redis
	Mongo  Mongo
	Auth   Auth
	Jwt    Jwt
}

type (
	App struct {
		Name string `default:"HeyTaxi Identity API"`
	}

	Server struct {
		Http ServerHttp
		Grpc ServerGrpc
	}

	ServerHttp struct {
		Host            string `default:""`
		Port            string `default:"8080"`
		BodyLimit       string `default:"1M"`
		RequestTimeout  int    `default:"60"`
		ShutdownTimeout int    `default:"5"`
	}

	ServerGrpc struct {
		Host                  string `default:""`
		Port                  string `default:"50051"`
		MaxConnectionIdle     int    `default:"10"`
		Timeout               int    `default:"10"`
		MaxConnectionAge      int    `default:"10"`
		MaxConnectionAgeGrace int    `default:"10"`
		Time                  int    `default:"20"`
	}

	Redis struct {
		Addr         string `default:"localhost:6379"`
		Password     string `default:""`
		DB           int    `default:""`
		DefaultDb    string `default:""`
		MinIdleConns int    `default:""`
		PoolSize     int    `default:""`
		PoolTimeout  int    `default:""`
		MaxRetries   int    `default:"3"`
	}

	Mongo struct {
		Uri               string `default:"mongodb://root:root@localhost:27017"`
		ConnectionTimeout int    `default:"3"`
		SocketTimeout     int    `default:"3"`
	}

	Auth struct {
		DatabaseName   string `default:"auth"`
		CollectionName string `default:"users"`
	}

	Jwt struct {
		AccessTokenExp             int    `default:"3600"`
		RefreshTokenExp            int    `default:"1296000"`
		Issuer                     string `default:"hey-taxi-identity-api"`
		AccessTokenPrivateKeyFile  string `default:"/etc/certs/access-token-private-key.pem"`
		AccessTokenPublicKeyFile   string `default:"/etc/certs/access-token-public-key.pem"`
		RefreshTokenPrivateKeyFile string `default:"/etc/certs/refresh-token-private-key.pem"`
		RefreshTokenPublicKeyFile  string `default:"/etc/certs/refresh-token-public-key.pem"`
	}
)
