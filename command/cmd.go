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
}

var commands = make(map[string]Command)

// RegisterCommands loads all available commands into memory.
func RegisterCommands() {
	commands["help"] = Command{
		Name:        "help",
		Description: "List all available commands",
	}

	commands["query"] = Command{
		Name:        "query",
		Description: "Ask a plain question to the LLM. Usage: /query What is a goroutine? or simply input your question and it will be handled as a query",
	}

	commands["format"] = Command{
		Name:        "format",
		Description: "Format a file using the LLM. Usage: /format main.go 'make it idiomatic or /format file:main.go 'make it idiomatic'",
	}
	commands["clear"] = Command{
		Name:        "clear",
		Description: "Clear the terminal",
	}
	commands["exit"] = Command{
		Name:        "exit",
		Description: "Exit the terminal",
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
	fls := extractFilesFromInput(query, client.Cfg)

	if len(fls) == 0 {
		return errors.New("no file found in working directory")
	}

	// Assume first match
	var filePath string
	for _, path := range fls {
		filePath = path
		break
	}

	code, err := files.ReadFile(filePath)
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
	// TODO: write to file
	return nil
}

func processHelpCommand() error {
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("/%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func processClearCommand() error {
	utils.ClearScreen()
	return nil
}

func processExitCommand() error {
	fmt.Println("Exiting. Goodbye!")
	os.Exit(0)
	return nil
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
		filePath, err := files.FindFile(match[1], cfg.Session.GetWorkingDir())
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
			if path, err := files.FindFile(fields[0], cfg.Session.GetWorkingDir()); err == nil {
				filesMap[fields[0]] = path
			}
		}
	}

	return filesMap
}
