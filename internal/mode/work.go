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

var commitMessages = []string{
	"feat(auth): implement JWT authentication middleware",
	"feat(database): add user preferences table migration",
	"feat(api): create endpoint for batch processing orders",
	"fix(cache): resolve Redis connection timeout issue",
	"perf(query): optimize user search with database indexing",
	"refactor(middleware): simplify error handling logic",
	"test(auth): add unit tests for password hashing",
	"security(api): implement rate limiting for public endpoints",
	"fix(logging): handle concurrent write operations to log files",
	"feat(queue): implement message broker for async tasks",
	"feat(ui): add responsive navigation drawer component",
	"fix(layout): resolve overflow issues in mobile view",
	"style(theme): update primary color palette",
	"perf(rendering): implement virtual scrolling for large lists",
	"feat(auth): add biometric authentication option",
	"fix(forms): validate phone numbers before submission",
	"a11y(buttons): improve screen reader compatibility",
	"feat(state): implement Redux toolkit for user settings",
	"perf(bundle): reduce main bundle size by code splitting",
	"style(components): standardize button styling across app",
	"ci(pipeline): configure GitHub Actions workflow",
	"build(docker): optimize container layer caching",
	"docs(api): update OpenAPI documentation",
	"chore(deps): update dependencies to latest versions",
	"feat(monitoring): integrate Prometheus metrics",
	"test(e2e): add Cypress tests for checkout flow",
	"security(deps): patch vulnerable dependencies",
	"refactor(config): centralize environment variables",
	"perf(cdn): implement asset caching strategy",
	"docs(readme): update deployment instructions",
}

func handleWorkFlags() bool {
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
	} else {
		commit.Ticket = fmt.Sprintf("[%d]", rand.Intn(10000))
	}

	// Set message if provided
	if args.Message != "" {
		commit.Message = args.Message
	} else {
		commit.Message = commitMessages[rand.Intn(len(commitMessages))]
	}

	commit.Files = files

	return commit
}

func Work(args types.WorkArgs, dryrun bool) {
	if !handleWorkFlags() {
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

	if !dryrun {
		// Change file whitespace to create a change
		for _, file := range files {
			err := system.AddRemoveNewLine(file)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
		}

		// Craft commit
		commit := craftCommit(args, files)

		for _, file := range commit.Files {
			fmt.Printf("    - %s\n", file)
		}

		if err := system.Commit(commit); err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	} else {
		commit := craftCommit(args, files)

		fmt.Printf("Commit Details:\n"+
			"  Date: %s\n"+
			"  Ticket: %s\n"+
			"  Message: %s\n"+
			"  Files (%d):\n",
			commit.Date,
			commit.Ticket,
			commit.Message,
			len(commit.Files))

		for _, file := range commit.Files {
			fmt.Printf("    - %s\n", file)
		}
	}
}
