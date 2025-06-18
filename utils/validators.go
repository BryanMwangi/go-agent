package utils

import (
	"fmt"
	"regexp"
)

func ValidateName(name string) (bool, string, error) {
	if name == "" {
		return false, "", fmt.Errorf("name is empty")
	}
	// Define the regex patterns
	nameRegex := `^[A-Za-z\s'-À-ÖØ-öø-ÿ]+$`
	sanitizerRegex := `[\/\?<>\\:\*\|":]`

	// Compile the regex for illegal characters
	sanitizerRe, err := regexp.Compile(sanitizerRegex)
	if err != nil {
		return false, "", fmt.Errorf("failed to compile sanitizer regex: %v", err)
	}

	// Remove illegal characters
	sanitizedString := sanitizerRe.ReplaceAllString(name, "")

	// Compile the regex for name validation
	re, err := regexp.Compile(nameRegex)
	if err != nil {
		return false, "", fmt.Errorf("failed to compile name regex: %v", err)
	}

	if re.MatchString(sanitizedString) {
		// Check if the sanitized string is empty
		if sanitizedString == "" {
			return false, "", fmt.Errorf("name contains only illegal characters")
		}
		return true, sanitizedString, nil
	} else {
		return false, "", fmt.Errorf("name contains only illegal characters")
	}
}
