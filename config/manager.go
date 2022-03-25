package config

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ConfigManager struct {
	v    *viper.Viper
	c    *Config
	file string
}

func NewManager() *ConfigManager {
	m := &ConfigManager{
		c: defaults(),
		v: viper.New(),
	}

	return m
}

// load reads in config file and ENV variables if set.
func (cm *ConfigManager) load() error {
	rpl := strings.NewReplacer(".", "_")

	if cm.file == "" {
		cm.file = ".env"
	}

	cm.v.SetTypeByDefaultValue(true) // set default value if no value is provided
	cm.v.SetEnvKeyReplacer(rpl)      // replace dots with underscores
	cm.v.SetConfigType("env")        // REQUIRED if the config file does not have the extension in the name
	cm.v.SetConfigName(cm.file)      // name of config file (without extension)
	cm.v.AddConfigPath("..")         // optionally look for config in the working directory
	cm.v.AddConfigPath(".")          // optionally look for config in the working directory
	cm.v.AutomaticEnv()              // read in environment variables that match

	cm.bindEnvs(cm.c)

	if err := cm.v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return errors.Wrap(err, "failed to read config file")
		}
	}

	return nil
}

// bindEnvs binds ENV variables for Viper package.
func (cm *ConfigManager) bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	if ift.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
		ift = ift.Elem()
	}

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")

		if !ok {
			tv = strings.ToLower(t.Name)
		}

		skip := tv == "-" || tv == ",squash"

		p := parts

		if !skip {
			p = append(p, tv)
		}

		switch v.Kind() {
		case reflect.Struct:
			cm.bindEnvs(v.Interface(), parts...)
		default:
			if ok {
				bn := strings.ToUpper(strings.Join(p, "_"))
				cm.v.BindEnv(bn)
			}
		}
	}
}

// unmarshal reads in config file and ENV variables if set.
func (cm *ConfigManager) unmarshal() error {
	err := cm.v.Unmarshal(&cm.c)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}

	return nil
}
