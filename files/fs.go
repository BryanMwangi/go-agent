package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

func FindFile(fileName, workingDir string) (string, error) {
	var foundPath string

	err := filepath.WalkDir(workingDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err // stop if a path can't be read
		}
		if !d.IsDir() && strings.EqualFold(filepath.Base(path), fileName) {
			foundPath = path
			return filepath.SkipDir // stop walking once we find it
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	if foundPath == "" {
		return "", errors.New("file not found: " + fileName)
	}

	return foundPath, nil
}

func ReadFile(filePath string) ([]byte, error) {
	code, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return code, nil
}

func WriteFile(filePath string, code []byte) error {
	err := os.WriteFile(filePath, code, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
