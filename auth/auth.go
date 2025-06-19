package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/BryanMwangi/go-agent/config"
	"github.com/BryanMwangi/go-agent/prompts"
	"github.com/BryanMwangi/go-agent/utils"
)

func AuthenticateUser(cfg config.Config) {
	name := prompts.PromptUsername()
	model := prompts.PromptModel()
	apiKey := prompts.PromptAPIKey()

	session := config.InitSession(name, model)
	cfg.Name = name
	// update config
	cfg.APIKey = apiKey
	cfg.Session = session

	// show loader
	utils.ShowLoader("Verifying API Key...")
	err := validateUser(cfg)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	utils.StopLoader(3 * time.Second)

	cfg.UpdateConfig()
}
