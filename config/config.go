package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	MySql *MySqlConfiguration
}

type MySqlConfiguration struct {
	Database string
	Username string
	Password string
	Host     string
	Port     int
}

func Initialize() *Configuration {
	// Initialize Viper
	setUpViper()
	configuration := &Configuration{}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("unable to decode configuration, %v", err)
	}
	// fmt.Print(configuration)
	return configuration
}

// func initMySQL("mysql")
func setUpViper() {

	// Set config file type
	viper.SetConfigType("toml")

	viper.SetConfigFile("/home/kunal/tweeter_api/config/config.toml")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to read configuration file: %s", err))
	}
}
