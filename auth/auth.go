package auth

import (
	"github.com/BryanMwangi/go-agent/config"
	"github.com/BryanMwangi/go-agent/prompts"
)

func AuthenticateUser() config.Session {
	name := prompts.PromptUsername()
	model := prompts.PromptModel()
	apiKey := prompts.PromptAPIKey()

	return config.InitSession(name, model, apiKey)
}
