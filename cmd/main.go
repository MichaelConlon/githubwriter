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
	flag.StringVar(&cfg.Mode, "mode", "", "Operation mode (required)")
	flag.StringVar(&cfg.Text, "text", "", "text (required if write mode selected)")
	flag.IntVar(&cfg.OffsetLines, "offset", 0, "number of lines to offset the text in your git activity tracker")
	flag.BoolVar(&cfg.DryRun, "dryrun", false, "print the resulting git activity as ascii in the terminal")

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
