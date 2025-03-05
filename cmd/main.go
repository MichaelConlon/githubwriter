// cmd/main.go
package main

import (
	"flag"
	"fmt"
	"os"

	mode "github.com/MichaelConlon/githubwriter/internal/mode"
	types "github.com/MichaelConlon/githubwriter/internal/types"
)

func main() {
	// Create config instance
	cfg := &types.Config{}

	// Define the mode flag
	flag.StringVar(&cfg.Mode, "mode", "", "Operation mode: text, work")
	flag.StringVar(&cfg.TextArgs.Text, "text", "", "text mode only - text to be printed in the activity tracker (8 letter max)")
	flag.IntVar(&cfg.TextArgs.OffsetLines, "offset", 0, "text mode only - number of columns to offset the text in your git activity tracker")
	flag.BoolVar(&cfg.DryRun, "dryrun", false, "run the code without creating any commits. This will only print result that would have been commited")
	flag.IntVar(&cfg.TextArgs.Year, "year", 0, "print text to the activity tracker for a specific year")
	flag.StringVar(&cfg.WorkArgs.Date, "date", "", "work mode only - date to be used for the commit - format YYYY-MM-DD")
	flag.StringVar(&cfg.WorkArgs.Ticket, "ticket", "", "work mode only - ticket number to be used for the commit")
	flag.StringVar(&cfg.WorkArgs.Message, "message", "", "work mode only - message to be used for the commit")
	flag.StringVar(&cfg.WorkArgs.Extensions, "extensions", "", "work mode only - file extensions to be used for the commit")
	flag.IntVar(&cfg.WorkArgs.NumFiles, "num-files", 0, "work mode only - number of files to be used for the commit - 0 or not provided will use all files")
	flag.StringVar(&cfg.WorkArgs.Files, "files", "", "work mode only - comma separated list of files to be used for the commit")
	flag.StringVar(&cfg.WorkArgs.LintConfig, "lint-config", "", "work mode only - path to the eslint config file")

	// Parse command line arguments
	flag.Parse()

	// Check if mode is provided
	if cfg.Mode == "" {
		fmt.Println("Error: mode is required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Starting application in %s mode...\n", cfg.Mode)

	// Use the mode value
	switch cfg.Mode {
	case "text":
		fmt.Println("text mode selected")
		mode.Text(cfg.TextArgs, cfg.DryRun)
	case "work":
		fmt.Println("Work mode selected")
		mode.Work(cfg.WorkArgs, cfg.DryRun)
	default:
		fmt.Printf("Error: unknown mode '%s'\n", cfg.Mode)
		os.Exit(1)
	}
}
