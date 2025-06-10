package utils

import (
	"github.com/spf13/viper"
)

// Config contains the configuration for the application.
// the values are read by viper from a config file or environment variables.
type Config struct {
	DSN                 string `mapstructure:"DSN"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

// LoadConfig loads the configuration from a config file or environment variables.
// func LoadConfig(path string) (config Config, err error) { // <--- HAPUS PATH DARI ARGUMEN
func LoadConfig() (config Config, err error) { // <--- JADI SEPERTI INI
	// viper.AddConfigPath(path) // <--- HAPUS BARIS INI
	viper.AddConfigPath(".") // <--- Tambahkan ini untuk selalu mencari di direktori saat ini
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