package utils

import (
	"github.com/spf13/viper"
	"path/filepath" // Import ini
)

// Config ...
type Config struct {
	DSN                 string `mapstructure:"DSN"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

// LoadConfig loads the configuration explicitly from app.env in the current directory.
func LoadConfig() (config Config, err error) {
	// Dapatkan path absolut ke file app.env di direktori saat ini
	currentDir, err := os.Getwd() // Anda perlu import "os"
	if err != nil {
		return config, err
	}
	configFilePath := filepath.Join(currentDir, "app.env")

	viper.SetConfigFile(configFilePath) // Mengatur path dan nama file secara eksplisit

	viper.AutomaticEnv() // Masih penting untuk variabel lingkungan

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}