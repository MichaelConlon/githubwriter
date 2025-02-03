package system

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/MichaelConlon/githubwriter/internal/types"
)

func Commit(commit types.Commit) error {
	// Set Git environment variables
	env := append(os.Environ(),
		fmt.Sprintf("GIT_AUTHOR_DATE=%s 12:00:00", commit.Date),
		fmt.Sprintf("GIT_COMMITTER_DATE=%s 12:00:00", commit.Date),
	)

	// Add all changes
	for _, file := range commit.Files {
		addCmd := exec.Command("git", "add", file)
		if err := addCmd.Run(); err != nil {
			return fmt.Errorf("failed to git add file: %s:  Error: %w", file, err)
		}
	}

	// Make the commit
	commitCmd := exec.Command("git", "commit", "-m", commit.Message)
	commitCmd.Env = env
	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("failed to git commit: %w", err)
	}

	return nil
}
