package textWriter

import (
	"flag"
	"fmt"
	"os"

	types "github.com/MichaelConlon/githubwriter/internal/types"
)

// Define alphabet
var (
	a = types.Letter{
		{0, 0, 1, 1, 1},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 1, 1},
	}

	b = types.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0},
	}

	c = types.Letter{
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
	}

	d = types.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}

	e = types.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
	}

	f = types.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 0, 0, 0},
	}

	g = types.Letter{
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}

	h = types.Letter{
		{1, 1, 1, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 1},
	}

	i = types.Letter{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
	}

	j = types.Letter{
		{1, 0, 0, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
	}

	k = types.Letter{
		{1, 1, 1, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
	}

	l = types.Letter{
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1},
	}

	m = types.Letter{
		{1, 1, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}

	n = types.Letter{
		{1, 1, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 1, 1, 1},
	}

	o = types.Letter{
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}

	p = types.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{0, 1, 0, 0, 0},
	}

	q = types.Letter{
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 1},
	}

	r = types.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1},
	}

	s = types.Letter{
		{0, 1, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 1, 0},
	}

	t = types.Letter{
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	}

	u = types.Letter{
		{1, 1, 1, 1, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
	}

	v = types.Letter{
		{1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
	}

	w = types.Letter{
		{1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 1, 1, 1},
	}

	x = types.Letter{
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
	}

	y = types.Letter{
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
	}

	z = types.Letter{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 0, 0, 1},
		{1, 0, 0, 0, 1},
	}

	// Create alphabet slice
	alphabet = [26]types.Letter{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z}
)

func writeSpace() {
	for i := 0; i < 7; i++ {
		fmt.Print("")
	}
	fmt.Println() // New line after each row
}

// Print the characters to the terminal
// Formatted to represent the git activity tracker
func writeWordDryRun(offset int, text string) {
	for depthIdx, _ := range alphabet[0] { //used to cycle through depth of character
		for i := 0; i < offset; i++ {
			fmt.Printf("%c", '\u25A1')
		}
		for _, char := range text { //used to cycle through laters

			idx := int(char - 'a')
			letterArray := alphabet[idx]
			for _, bit := range letterArray {
				if bit[depthIdx] == 1 {
					fmt.Printf("%c", '\u25A3')
				} else {
					fmt.Printf("%c", '\u25A1')
				}
			}
			fmt.Printf("%c", '\u25A1')
		}
		remainingActivityColumns := 53 - offset - len(text)*6 - 1
		for i := 0; i < remainingActivityColumns; i++ {
			fmt.Printf("%c", '\u25A1')
		}
		fmt.Println()
	}
}

func writeLetter(char byte) {
	idx := int(char - 'a')
	letterArray := alphabet[idx]

	for i := 0; i < len(letterArray); i++ {
		for j := 0; j < len(letterArray[i]); j++ {
			if j == 0 {
				fmt.Printf("%c", '\u25A1')
			}
			if letterArray[i][j] == 0 {
				fmt.Printf("%c", '\u25A1')
			} else {
				fmt.Printf("%c", '\u25A3')
			}
		}
		fmt.Printf("%c", '\u25A1')
		fmt.Println() // New line after each row
	}
}

func writeWord(offsetLines int, text string) {
	// whitespace lines
	for i := 0; i < offsetLines; i++ {
		writeSpace()
	}

	for i := 0; i < len(text); i++ {
		writeLetter(text[i])
		fmt.Println()
	}
}

// Entry Point
// Handle flags specific to the write mode
// set up commits
func Write(cfg *types.Config) {
	if cfg.Mode != "" {
		fmt.Printf("write mode: '%s'\n", cfg.Mode)
	}

	// Check if text exists and length
	if cfg.Text == "" {
		fmt.Println("Error: text is required")
		flag.Usage()
		os.Exit(1)
	}
	if len(cfg.Text) > 8 {
		fmt.Printf("Error: text must be 8 characters or less, got %d\n", len(cfg.Text))
		os.Exit(1)
	}

	if cfg.DryRun {
		writeWordDryRun(cfg.OffsetLines, cfg.Text)
	} else {
		writeWord(cfg.OffsetLines, cfg.Text)
	}
}
