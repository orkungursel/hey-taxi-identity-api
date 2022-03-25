package config

import (
	"reflect"
	"testing"
)

func TestOverrideEnvironments_WhenEnvFileExists(t *testing.T) {
	wantAppName := "Foobar"
	wantRedisAddr := "localhost2:6379"
	wantAccessTokenPath := "/foo/bar/access_token.pem"

	t.Setenv("APP_NAME", wantAppName)
	t.Setenv("REDIS_ADDR", wantRedisAddr)
	t.Setenv("JWT_ACCESS_TOKEN_PRIVATE_KEY_FILE", wantAccessTokenPath)

	c := NewConfigWithFile("")

	t.Run("override app_name", func(t *testing.T) {
		if c.App.Name != wantAppName {
			t.Errorf("c.App.Name = %v, want %v", c.App.Name, wantAppName)
		}
	})

	t.Run("override redis_addr", func(t *testing.T) {
		if c.Redis.Addr != wantRedisAddr {
			t.Errorf("c.Redis.Addr = %v, want %v", c.Redis.Addr, wantRedisAddr)
		}
	})

	t.Run("override jwt", func(t *testing.T) {
		if c.Jwt.AccessTokenPrivateKeyFile != wantAccessTokenPath {
			t.Errorf("c.Redis.Addr = %v, want %v", c.Jwt.AccessTokenPrivateKeyFile, wantAccessTokenPath)
		}
	})
}

func TestOverrideEnvironments_WhenEnvFileNotExists(t *testing.T) {
	wantAppName := "Foobar"
	wantRedisAddr := "http://foo.bar"
	wantAccessTokenPath := "/foo/bar/access_token.pem"

	t.Setenv("APP_NAME", wantAppName)
	t.Setenv("REDIS_ADDR", wantRedisAddr)
	t.Setenv("JWT_ACCESS_TOKEN_PRIVATE_KEY_FILE", wantAccessTokenPath)

	c := NewConfigWithFile("-")

	t.Run("override app_name", func(t *testing.T) {
		if c.App.Name != wantAppName {
			t.Errorf("c.App.Name = %v, want %v", c.App.Name, wantAppName)
		}
	})

	t.Run("override redis_addr", func(t *testing.T) {
		if c.Redis.Addr != wantRedisAddr {
			t.Errorf("c.Redis.Addr = %v, want %v", c.Redis.Addr, wantRedisAddr)
		}
	})

	t.Run("override jwt", func(t *testing.T) {
		if c.Jwt.AccessTokenPrivateKeyFile != wantAccessTokenPath {
			t.Errorf("c.Redis.Addr = %v, want %v", c.Jwt.AccessTokenPrivateKeyFile, wantAccessTokenPath)
		}
	})
}

func TestConfig_defaults(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		envs    map[string]string
		file    string
		wantErr bool
	}{
		{
			name:    "defaults",
			want:    defaults(),
			file:    "not-exists",
			wantErr: false,
		},
		{
			name:    "with custom envs",
			file:    ".env-sample",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr {
				if got := NewConfigWithFile(tt.file); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}
			} else {
				if got := NewConfigWithFile(tt.file); reflect.DeepEqual(got, tt.want) {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestConfig_GetMode(t *testing.T) {
	tests := []struct {
		name string
		c    *Config
		env  string
		want string
	}{
		{
			name: "should return development mode",
			c:    NewConfigWithFile(""),
			want: "local",
		},
		{
			name: "should return production mode",
			c:    NewConfigWithFile(""),
			want: "production",
			env:  "production",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != "" {
				t.Setenv("ACTIVE_PROFILE", tt.env)
			}

			if got := tt.c.GetProfile(); got != tt.want {
				t.Errorf("Config.GetMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
