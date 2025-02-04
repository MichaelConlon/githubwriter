package system

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(filename string, content string) error {
	if err := os.WriteFile("contribution.txt", []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write contribution file: %w", err)
	}
	return nil
}

// Recursively find all files with the given extensions
func FindFilesByExtension(extensions string, dir string) ([]string, error) {
	extensionsList := []string{}
	if extensions != "" {
		extensionsList = strings.Split(extensions, ",")
	}

	// Grab all files in the current directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read current directory: %w", err)
	}

	// Recursively find all files with the given extensions
	matchedFiles := []string{}
	for _, file := range files {
		if file.IsDir() {
			// Recursively find all files with the given extensions
			dirFiles, err := FindFilesByExtension(extensions, filepath.Join(dir, file.Name())) // append found directory to path and recursively find
			if err != nil {
				return nil, fmt.Errorf("failed to find files in directory: %s\nError: %w", file.Name(), err)
			}
			pathedDirFiles := []string{}
			for _, dirFile := range dirFiles {
				pathedDirFiles = append(pathedDirFiles, filepath.Join(dir, dirFile))
			}
			matchedFiles = append(matchedFiles, pathedDirFiles...) // add all files from directory to matched files
		}
		// Get extention of file and check if it is in the list of allowed extensions
		parts := strings.Split(file.Name(), ".")
		if len(parts) > 1 {
			ext := parts[len(parts)-1]
			for _, allowedExt := range extensionsList {
				if ext == allowedExt {
					matchedFiles = append(matchedFiles, file.Name())
					break
				}
			}
		}
	}
	return matchedFiles, nil
}

func AddRemoveNewLine(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Add a new line to the end of the file if it doesn't already end with a newline
	if !strings.HasSuffix(string(content), "\n") {
		content = append(content, '\n')
		err = os.WriteFile(filename, content, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}

	// Remove a new line from the end of the file if it already ends with a newline
	if strings.HasSuffix(string(content), "\n") {
		content = content[:len(content)-1]
		err = os.WriteFile(filename, content, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}
	return nil
}
