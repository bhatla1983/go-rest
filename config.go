package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBServerName string `mapstructure:"DB_SERVER_NAME"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
}

//var AppConfig *Config

func LoadAppConfig(path string) (config Config, err error) {
	log.Println("Loading DB Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// func LoadAppConfig() {
// 	log.Println("Loading DB Server Configurations...")
// 	viper.AddConfigPath(".")
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("json")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = viper.Unmarshal(&AppConfig)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
