package config

import (
	"os"

	"github.com/spf13/viper"
)

func NewConfig() *Config {
	return NewConfigWithFile("")
}

func NewConfigWithFile(file string) *Config {
	cm := NewManager()
	cm.file = file
	cm.v = viper.New()
	cm.c = defaults()

	if err := cm.load(); err != nil {
		panic(err)
	}

	if err := cm.unmarshal(); err != nil {
		panic(err)
	}

	return cm.c
}

// GetProfile returns current mode by reading environment variable.
func (c *Config) GetProfile() string {
	profile, ok := os.LookupEnv("ACTIVE_PROFILE")
	if !ok {
		profile = "local"
	}

	return profile
}

func (c *Config) IsLocal() bool {
	return c.GetProfile() == "local"
}

func (c *Config) IsProduction() bool {
	return c.GetProfile() == "production"
}
