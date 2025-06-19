package command

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BryanMwangi/go-agent/config"
	"github.com/BryanMwangi/go-agent/files"
	"github.com/BryanMwangi/go-agent/llm"
	"github.com/BryanMwangi/go-agent/utils"
)

type Command struct {
	Name        string
	Description string
	Execute     func(args []string, client *llm.Client) error
}

var commands = make(map[string]Command)

// RegisterCommands loads all available commands into memory.
func RegisterCommands() {
	commands["help"] = Command{
		Name:        "help",
		Description: "List all available commands",
		Execute:     processHelpCommand,
	}

	commands["query"] = Command{
		Name:        "query",
		Description: "Ask a plain question to the LLM. Usage: /query What is a goroutine?",
		Execute:     processQueryCommand,
	}

	commands["format"] = Command{
		Name:        "format",
		Description: "Format a file using the LLM. Usage: /format main.go 'make it idiomatic'",
		Execute:     processFormatCommand,
	}
}

func processQueryCommand(args []string, client *llm.Client) error {
	if len(args) == 0 {
		return errors.New("usage: /query [your question]")
	}
	query := strings.Join(args, " ")
	resp, err := llm.Query(client, query)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func processFormatCommand(args []string, client *llm.Client) error {
	if len(args) < 1 {
		return errors.New("usage: /format [filename] [message]")
	}

	query := strings.Join(args, " ")
	files := extractFilesFromInput(query, client.Cfg)

	if len(files) == 0 {
		return errors.New("no file found in working directory")
	}

	// Assume first match
	var filePath string
	for _, path := range files {
		filePath = path
		break
	}

	code, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	message := strings.ReplaceAll(query, filepath.Base(filePath), "")
	lang := utils.LangFromExt(filePath)

	resp, err := llm.Format(client, message, string(code), lang)
	if err != nil {
		return err
	}

	fmt.Println("Formatted output:\n" + string(resp))
	return nil
}

func processHelpCommand(args []string, client *llm.Client) error {
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("/%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

// ParseAndExecute checks if input is a command and runs it.
func ParseAndExecute(input string, client *llm.Client) error {
	if !strings.HasPrefix(input, "/") {
		return handleUnknownCommand(input, client)
	}

	parts := strings.Fields(strings.TrimPrefix(input, "/"))
	if len(parts) == 0 {
		return errors.New("empty command")
	}

	cmdName := parts[0]
	args := parts[1:]

	cmd, exists := commands[cmdName]
	if !exists {
		return handleUnknownCommand(input, client)
	}

	err := cmd.Execute(args, client)
	return err
}

func handleUnknownCommand(input string, client *llm.Client) error {
	input = strings.TrimPrefix(input, "/")
	args := strings.Fields(input)
	return processQueryCommand(args, client)
}

func extractFilesFromInput(input string, cfg *config.Config) map[string]string {
	filesMap := make(map[string]string)
	fileRegex := regexp.MustCompile(`(?m)file:\s*([^\s]+)`)

	// Look for "file: filename.go" patterns
	matches := fileRegex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		filePath, err := files.FindFile(match[1], cfg.Session.WorkDir)
		if err != nil {
			fmt.Println("Error finding file:", err)
			continue
		}
		filesMap[match[1]] = filePath
	}

	// Fallback: if no "file:" was found, treat first word as filename
	if len(filesMap) == 0 {
		fields := strings.Fields(input)
		if len(fields) > 0 {
			if path, err := files.FindFile(fields[0], cfg.Session.WorkDir); err == nil {
				filesMap[fields[0]] = path
			}
		}
	}

	return filesMap
}
