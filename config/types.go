package config

import "time"

type Config struct {
	Name              string        `json:"name"`
	Version           string        `json:"version"`
	APIKey            string        `json:"api_key"`
	APITimeout        time.Duration `json:"api_timeout"`
	MaxContextSize    int           `json:"max_context_size"`
	MaxRequestRetries int           `json:"max_request_retries"`
	Temperature       float64       `json:"temperature"`
	EnableHistory     bool          `json:"enable_history"`
	ConfigPaths       ConfigPaths   `json:"config_paths"`
	Session           Session       `json:"session"`
}

type ConfigPaths struct {
	ConfigDir  string
	ConfigFile string
}

type APIEndpoint struct {
	BaseURL        string
	CompletionPath string
}

type Session struct {
	Model   string `json:"model"`
	WorkDir string `json:"work_dir"`
	// "openai" or "anthropic". For now, we will only support OpenAI.
	Provider string `json:"provider"`
	// The API endpoint to use. Obtained from the provider after validating
	// the model chosen
	API APIEndpoint `json:"api"`
}
