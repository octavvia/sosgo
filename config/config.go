package config

import "github.com/spf13/viper"

type config struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_URL      string
	DB_DATABASE string
}

var ENV *config

// menghubungkan dengan env
func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}

}
