package utils

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
)

var extToLang = map[string]string{
	".go":    "go",
	".js":    "javascript",
	".py":    "python",
	".ts":    "typescript",
	".jsx":   "javascript",
	".tsx":   "typescript",
	".java":  "java",
	".kt":    "kotlin",
	".rb":    "ruby",
	".php":   "php",
	".cs":    "csharp",
	".cpp":   "cpp",
	".c":     "c",
	".h":     "c",
	".hpp":   "c",
	".cc":    "c",
	".hh":    "c",
	".m":     "objective-c",
	".mm":    "objective-c",
	".swift": "swift",
	".rs":    "rust",
	".exs":   "elixir",
	".ex":    "elixir",
	".ex1":   "elixir",
}

func LangFromExt(filename string) string {
	return extToLang[filepath.Ext(filename)]
}

func ParseCode(code, lang string) string {
	newCode := fmt.Sprintf("```%s\n%s\n```", lang, code)
	return newCode
}

func ParseCodeRequest(req, query, code string) string {
	newReq := fmt.Sprintf("%s\n%s\n%s\n", req, query, code)
	return newReq
}

func ExtractFirstCodeBlock(response string) (string, error) {
	re := regexp.MustCompile("(?s)```[a-zA-Z]*\\n(.*?)```")
	matches := re.FindStringSubmatch(response)
	if len(matches) < 2 {
		return "", errors.New("no code block found")
	}
	return matches[1], nil
}
