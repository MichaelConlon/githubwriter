package mode

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	system "github.com/MichaelConlon/githubwriter/internal/system"
	types "github.com/MichaelConlon/githubwriter/internal/types"
)

func handleWorkFlags(args types.WorkArgs) bool {
	valid := true

	return valid
}

func craftCommit(args types.WorkArgs, files []string) types.Commit {
	commit := types.Commit{
		Date:    args.Date,
		Ticket:  args.Ticket,
		Message: args.Message,
	}

	// Set date if provided
	if args.Date != "" {
		commitDate, err := time.Parse("2006-01-02", args.Date)
		if err != nil {
			fmt.Println("Error: invalid date format")
			flag.Usage()
			os.Exit(1)
		}
		commit.Date = commitDate.Format("2006-01-02")
	} else {
		commit.Date = time.Now().Format("2006-01-02")
	}

	// Set ticket if provided
	if args.Ticket != "" {
		commit.Ticket = args.Ticket
	}

	// Set message if provided
	if args.Message != "" {
		commit.Message = args.Message
	}

	commit.Files = files

	return commit
}

func Work(args types.WorkArgs) {
	if !handleWorkFlags(args) {
		flag.Usage()
		os.Exit(1)
	}

	files, err := system.FindFilesByExtension(args.Extensions, ".")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// Randomly select NumFiles if specified
	if args.NumFiles > 0 && args.NumFiles < len(files) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(files), func(i, j int) {
			files[i], files[j] = files[j], files[i]
		})
		files = files[:args.NumFiles]
	}

	commit := craftCommit(args, files)

	fmt.Printf("Commit: %+v\n", commit)
}
