package env

import (
	"github.com/spf13/viper"
)

type Config struct {
	GoogleApp string
}

var Conf *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.SetDefault("GOOGLE_APP", "")

	Conf = &Config{
		GoogleApp: viper.GetString("GOOGLE_APP"),
	}
}
