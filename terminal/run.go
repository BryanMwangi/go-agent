package terminal

import (
	"context"
	"fmt"

	"github.com/BryanMwangi/go-agent/command"
	"github.com/BryanMwangi/go-agent/config"
	"github.com/BryanMwangi/go-agent/llm"
	"github.com/BryanMwangi/go-agent/prompts"
)

// Our main app loop
func Run(cfg config.Config, ctx context.Context) {
	client := llm.NewClient(&cfg)

	// Ensure working directory before starting
	if cfg.Session.GetWorkingDir() == "" {
		dir := prompts.PromptWorkingDirectory()
		if dir != "" {
			cfg.Session.SetWorkingDir(dir)
			cfg.UpdateConfig()
		}
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nShutting down...")
			// TODO: save session or further tasks before exiting
			return
		default:
			input := prompts.PromptUserInput()
			if input != "" {
				if err := command.ProcessUserInput(input, client); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
