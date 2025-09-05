package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewViper is a function to load config from .env file
func NewViper() *viper.Viper {
	config := viper.New()

	// Set config to read from .env file
	config.SetConfigName(".env")
	config.SetConfigType("env")
	config.AddConfigPath("./")
	
	// Enable reading from environment variables
	config.AutomaticEnv()

	err := config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return config
}
