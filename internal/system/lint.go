package system

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunESLint(lintConfig string, filesToLint []string) error {
	// Create the command
	args := append([]string{"eslint", "-c", lintConfig}, filesToLint...)
	cmd := exec.Command("npx", args...)

	// Capture both stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("eslint failed: %w\nStderr: %s", err, stderr.String())
	}

	return nil
}
