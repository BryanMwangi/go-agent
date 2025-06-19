package command

import (
	"strings"

	"github.com/BryanMwangi/go-agent/llm"
)

var cmdAliases = map[string]string{
	"h":   "help",
	"q":   "query",
	"fmt": "format",
	"cls": "clear",
}

func normalizeCmd(cmd string) string {
	if canonical, ok := cmdAliases[cmd]; ok {
		return canonical
	}
	return cmd
}

func ProcessUserInput(input string, client *llm.Client) error {
	isCommand, cmd, newInput := extractCommandFromInput(input)
	if isCommand {
		return parseAndExecute(cmd, newInput, client)
	}
	return processQueryCommand([]string{newInput}, client)
}

// ParseAndExecute checks if input is a command and runs it.
func parseAndExecute(cmd, input string, client *llm.Client) error {
	switch normalizeCmd(cmd) {
	case "help":
		return processHelpCommand([]string{}, client)
	case "query":
		return processQueryCommand([]string{input}, client)
	case "format":
		return processFormatCommand([]string{input}, client)
	case "clear":
		return processClearCommand([]string{}, client)
	default:
		return handleUnknownCommand(input, client)
	}
}

func extractCommandFromInput(input string) (bool, string, string) {
	if !strings.HasPrefix(input, "/") {
		return false, "", input
	}
	parts := strings.Fields(strings.TrimPrefix(input, "/"))
	if len(parts) == 0 {
		return false, "", input
	}
	newInput := strings.Join(parts[1:], " ")
	return true, parts[0], newInput
}
