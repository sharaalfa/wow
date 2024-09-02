package config

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestGetConfig(t *testing.T) {
	err := os.Setenv("CONFIG_FILE", "config_test.yaml")
	if err != nil {
		log.Fatalf("Error setting CONFIG_FILE environment, %s", err)
		return
	}

	InitConfig()

	expectedEntryPoint := "https://api.quotable.io/random"
	actualEntryPoint := GetConfig().Quotes.EntryPoint
	if actualEntryPoint != expectedEntryPoint {
		t.Errorf("Expected ENTRY POINT %s, but got %s", expectedEntryPoint, actualEntryPoint)
	}
}

func TestInitConfig(t *testing.T) {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.yaml"
	}
	basePath := getBasePath()
	configPath := filepath.Join(basePath, "..", "..", "config")
	setConfig(configPath)
}
