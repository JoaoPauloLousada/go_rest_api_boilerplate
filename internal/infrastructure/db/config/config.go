package configuration

import "github.com/spf13/viper"

type Config struct {
	ApiPort int    `mapstructure:"API_PORT"`
	DBHost  string `mapstructure:"DB_HOST"`
	DBPort  int    `mapstructure:"DB_PORT"`
	DBUser  string `mapstructure:"DB_USER"`
	DBPass  string `mapstructure:"DB_PASS"`
	DBName  string `mapstructure:"DB_NAME"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
