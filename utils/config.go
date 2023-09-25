package utils

import "github.com/spf13/viper"

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	TokenLifespan  int    `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	SecretString   string `mapstructure:"SECRET_STRING"`
}

var config Config
var err error

var vpr = viper.ReadInConfig

func init() {
	configPaths := []string{".", "../", "../../"}

	for _, path := range configPaths {
		config, err = LoadConfig(path)
		if err == nil {
			break
		}
	}

}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = vpr()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func GetAppConfig() Config {
	return config
}
