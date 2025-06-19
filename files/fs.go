package files

import (
	"errors"
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
