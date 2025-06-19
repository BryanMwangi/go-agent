package auth

import (
	"fmt"

	"github.com/BryanMwangi/go-agent/config"
	"github.com/BryanMwangi/go-agent/llm"
	"github.com/BryanMwangi/go-agent/utils"
)

func validateUser(cfg config.Config) error {
	// create a new client only for the purpose of validating the API key
	client := llm.NewClient(&cfg)
	r, err := llm.OnStartQuery(client)
	if err != nil {
		fmt.Printf("Invalid API key. Please try again. Error: %s\n", err)

		return config.InvalidateSession(cfg)
	}
	fmt.Println(string(r))

	utils.Welcome(cfg.Name)
	return nil
}
