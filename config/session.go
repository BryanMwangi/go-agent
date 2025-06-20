package config

import (
	"fmt"
	"os"
)

var chatGPTAPIEndpoint = APIEndpoint{
	BaseURL:        "https://api.openai.com/v1",
	CompletionPath: "/chat/completions",
}

var providersEndpoints = map[string]APIEndpoint{
	"openai":    chatGPTAPIEndpoint,
	"anthropic": chatGPTAPIEndpoint,
}

func InitSession(name, model string) Session {
	if name == "" || model == "" {
		return Session{}
	}
	provider := getProvider(model)
	session := Session{
		Model:    model,
		Provider: provider,
		API:      providersEndpoints[provider],
	}
	return session
}

func getProvider(model string) string {
	for provider, models := range modelProviders {
		for _, m := range models {
			if m == model {
				return provider
			}
		}
	}
	fmt.Println("Model not found. Defaulting to OpenAI")
	return "openai"
}

func (S *Session) SetWorkingDir(dir string) {
	S.WorkDir = dir
}

func (S *Session) GetWorkingDir() string {
	return S.WorkDir
}

func InvalidateSession(cfg Config) error {
	err := os.Remove(cfg.ConfigPaths.ConfigFile)
	if err != nil {
		return err
	}
	fmt.Println("Restart the app to login again")
	os.Exit(0)
	return nil
}
