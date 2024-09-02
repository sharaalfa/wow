package config

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"runtime"
)

type Config struct {
	Quotes     QuotesConfig
	Server     ServerConfig
	Difficulty int
	Timeout    int
}

type QuotesConfig struct {
	EntryPoint string
	MaxRetries int
}

type ServerConfig struct {
	Host string
	Port string
}

func InitConfig() {
	basePath := getBasePath()
	configPath := filepath.Join(basePath, "..", "..", "config")
	setConfig(configPath)
}

func GetConfig() Config {
	return Config{
		Quotes: QuotesConfig{
			EntryPoint: viper.GetString("quotes.entry.point"),
			MaxRetries: viper.GetInt("quotes.maxRetries"),
		},
		Server: ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
		Difficulty: viper.GetInt("difficulty"),
		Timeout:    viper.GetInt("timeout"),
	}
}
func getBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
func setConfig(configPath string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

}
