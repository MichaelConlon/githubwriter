// cmd/main.go
package main

import (
	"flag"
	"fmt"
	"os"

	textWriter "github.com/MichaelConlon/githubwriter/internal"
	types "github.com/MichaelConlon/githubwriter/internal/types"
)

func main() {
	// Create config instance
	cfg := &types.Config{}

	// Define the mode flag
	flag.StringVar(&cfg.Mode, "mode", "", "Operation mode: text, work")
	flag.StringVar(&cfg.Text, "text", "", "text mode only - text to be printed in the activity tracker (8 letter max)")
	flag.IntVar(&cfg.OffsetLines, "offset", 0, "text mode only - number of lines to offset the text in your git activity tracker")
	flag.BoolVar(&cfg.DryRun, "dryrun", false, "run the code without creating any commits. This will only print result that would have been commited")

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
	case "write":
		fmt.Println("Write mode selected")
		textWriter.Write(cfg)
	case "delete":
		fmt.Println("Delete mode selected")
	default:
		fmt.Printf("Error: unknown mode '%s'\n", cfg.Mode)
		os.Exit(1)
	}
}
