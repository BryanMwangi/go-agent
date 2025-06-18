package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const version = "1.0.0"

var openAIModels = []string{
	"o4-mini",
	"o3",
	"o3-mini",
	"o1",
}

var anthropicModels = []string{
	"claude-2",
}

var modelProviders = map[string][]string{
	"openai":    openAIModels,
	"anthropic": anthropicModels,
}

var AvailableModels = append(openAIModels, anthropicModels...)

func InitConfig() Config {
	paths := buildPaths()

	config, err := LoadConfig(paths)
	if err != nil {
		config = defaultConfig(paths)
		err := SaveConfig(config)
		if err != nil {
			panic("Failed to save config: " + err.Error())
		}
	}

	return config
}

func buildPaths() ConfigPaths {
	configDir := filepath.Join(os.Getenv("HOME"), ".go-agent")
	configFile := filepath.Join(configDir, "config.json")

	fmt.Println(configDir)

	_ = os.MkdirAll(configDir, 0755)

	return ConfigPaths{
		ConfigDir:  configDir,
		ConfigFile: configFile,
	}
}

func defaultConfig(paths ConfigPaths) Config {
	return Config{
		Version:           version,
		APITimeout:        5 * time.Minute,
		MaxContextSize:    4000,
		MaxRequestRetries: 3,
		Temperature:       0.7,
		EnableHistory:     true,
		ConfigPaths:       paths,
	}
}

func SaveConfig(config Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(config.ConfigPaths.ConfigFile, data, 0644)
}

func LoadConfig(paths ConfigPaths) (Config, error) {
	data, err := os.ReadFile(paths.ConfigFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	config.ConfigPaths = paths
	return config, nil
}

func (C *Config) UpdateConfig() error {
	return SaveConfig(*C)
}
