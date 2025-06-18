package config

import "fmt"

var chatGPTAPIEndpoint = APIEndpoint{
	BaseURL:        "https://api.openai.com/v1",
	CompletionPath: "/chat/completions",
}

var providersEndpoints = map[string]APIEndpoint{
	"openai":    chatGPTAPIEndpoint,
	"anthropic": chatGPTAPIEndpoint,
}

func InitSession(name, model string) Session {
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
