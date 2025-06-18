package config

import (
	"encoding/json"
	"os"
	"path/filepath"
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

func InitConfig(model string) Config {
	paths := buildPaths()

	config, err := LoadConfig(paths)
	if err != nil {
		config = defaultConfig(model, paths)
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

	_ = os.MkdirAll(configDir, 0755)

	return ConfigPaths{
		ConfigDir:  configDir,
		ConfigFile: configFile,
	}
}

func defaultConfig(model string, paths ConfigPaths) Config {
	selectedModel := AvailableModels[0]
	for _, m := range AvailableModels {
		if m == model {
			selectedModel = m
			break
		}
	}

	return Config{
		Version:           version,
		Model:             selectedModel,
		APITimeout:        10,
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
