package system

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Commit(commitDate time.Time) error {
	// Write date to contribution file
	dateStr := commitDate.Format("2006-01-02")
	if err := os.WriteFile("contribution.txt", []byte(dateStr), 0644); err != nil {
		return fmt.Errorf("failed to write contribution file: %w", err)
	}

	// Set Git environment variables
	env := append(os.Environ(),
		fmt.Sprintf("GIT_AUTHOR_DATE=%s 12:00:00", dateStr),
		fmt.Sprintf("GIT_COMMITTER_DATE=%s 12:00:00", dateStr),
	)

	// Add all changes
	addCmd := exec.Command("git", "add", "contribution.txt")
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("failed to git add: %w", err)
	}

	// Make the commit
	commitCmd := exec.Command("git", "commit", "-m", fmt.Sprintf("contribution: %s", dateStr))
	commitCmd.Env = env
	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("failed to git commit: %w", err)
	}

	return nil
}
