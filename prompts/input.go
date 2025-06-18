package prompts

import (
	"fmt"
	"os"

	"github.com/BryanMwangi/go-agent/config"
	"github.com/BryanMwangi/go-agent/utils"
	"github.com/manifoldco/promptui"
)

func PromptUsername() string {
	validate := func(input string) error {
		pass, _, err := utils.ValidateName(input)
		if err != nil {
			return err
		}
		if !pass {
			return fmt.Errorf("name contains illegal characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter your name",
		Validate: validate,
	}
	name, err := prompt.Run()
	if err != nil {
		return ""
	}
	return name
}

func PromptAPIKey() string {
	prompt := promptui.Prompt{
		Label: "Enter your OpenAI API Key",
		Mask:  '*',
	}
	apiKey, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		os.Exit(1)
	}
	return apiKey
}

func PromptModel() string {
	prompt := promptui.Select{
		Label: "Select a model",
		Items: config.AvailableModels,
	}
	_, model, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		os.Exit(1)
	}
	return model
}
