package config

type Config struct {
	Version           string      `json:"version"`
	Model             string      `json:"model"`
	APITimeout        int         `json:"api_timeout"`
	MaxContextSize    int         `json:"max_context_size"`
	MaxRequestRetries int         `json:"max_request_retries"`
	Temperature       float64     `json:"temperature"`
	EnableHistory     bool        `json:"enable_history"`
	ConfigPaths       ConfigPaths `json:"-"`
}

type ConfigPaths struct {
	ConfigDir  string
	ConfigFile string
}

type APIEndpoint struct {
	BaseURL        string
	CompletionPath string
}
