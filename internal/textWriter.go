package textWriter

import (
	"flag"
	"fmt"
	"os"
	"time"

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

func commit(commitDate time.Time) {
	// stub to be implemented later
	fmt.Println("commit called")
}

// Print the characters to the terminal
// Formatted to represent the git activity tracker
func drawActivity(offset int, text string) {
	// print a leading space and a fully blank top row
	fmt.Println()
	for i := 0; i < 52; i++ {
		fmt.Printf("%c", '\u25A1')
	}
	fmt.Println()

	for depthIdx := range alphabet[0] { // used to cycle through depth of character
		// leading blank columns
		for i := 0; i < offset; i++ {
			fmt.Printf("%c", '\u25A1')
		}
		for _, char := range text { // used to cycle through letters

			idx := int(char - 'a')
			letterArray := alphabet[idx]
			for _, bit := range letterArray {
				if bit[depthIdx] == 1 {
					fmt.Printf("%c", '\u25A3') // filled box
				} else {
					fmt.Printf("%c", '\u25A1') // empty box
				}
			}
			fmt.Printf("%c", '\u25A1')
		}
		// fill in remaining columns after letters with blank pixels
		remainingActivityColumns := 53 - offset - len(text)*6 - 1
		for i := 0; i < remainingActivityColumns; i++ {
			fmt.Printf("%c", '\u25A1')
		}
		fmt.Println()
	}
	// print fully empty bottom row
	for i := 0; i < 52; i++ {
		fmt.Printf("%c", '\u25A1')
	}
	fmt.Println()
	fmt.Println()
}

// print empty line on the activity tracker
// really just increment date by 7 for 7 empty pixels
func writeSpace(commitDate time.Time) time.Time {
	for i := 0; i < 7; i++ {
		commitDate = commitDate.AddDate(0, 0, 1)
	}
	return commitDate
}

// take a character and convert it to index in the alphabet array
// print/commit the corresponding pixels
func writeLetter(char byte, commitDate time.Time, dryrun bool) time.Time {
	idx := int(char - 'a')
	letterArray := alphabet[idx]

	for outerIdx := range letterArray { // loop through outer array
		for innerIdx, bit := range letterArray[outerIdx] { // loop through inner array
			if innerIdx == 0 { // space at top = increment date
				commitDate = commitDate.AddDate(0, 0, 1)
			}
			if bit == 1 { // commit date == true
				if dryrun {
					fmt.Printf("%c: %s\n", char, commitDate.Format("2006-01-02"))
				} else {
					commit(commitDate)
				}
				commitDate = commitDate.AddDate(0, 0, 1) // move date forward
			}
		}
		commitDate = commitDate.AddDate(0, 0, 1) // space at bottom = increment date
	}
	return commitDate
}

func writeWord(offsetLines int, text string, commitDate time.Time, dryrun bool) {
	// whitespace lines
	// dryrun doesn't matter, this only increments date
	for i := 0; i < offsetLines; i++ {
		commitDate = writeSpace(commitDate)
	}

	// write each letter
	for i := 0; i < len(text); i++ {
		commitDate = writeLetter(text[i], commitDate, dryrun)
		commitDate = writeSpace(commitDate)
		fmt.Println()
	}
}

// Entry Point
// Handle flags specific to the write mode
func Write(cfg *types.Config) {
	if cfg.Mode != "" {
		fmt.Printf("write mode: '%s'\n", cfg.Mode)
	}

	// Check flags for valid values
	// refactor with error object later
	if cfg.Text == "" {
		fmt.Println("Error: text is required")
		flag.Usage()
		os.Exit(1)
	}
	if len(cfg.Text) > 8 {
		fmt.Printf("Error: text must be 8 characters or less, got %d\n", len(cfg.Text))
		os.Exit(1)
	}
	if cfg.Year < 0 {
		fmt.Printf("Error: year must be a positive integer, got %d\n", cfg.Year)
		os.Exit(1)
	}

	var commitDate time.Time
	if cfg.Year != 0 { // get year from args
		commitDate = time.Date(cfg.Year, 1, 1, 0, 0, 0, 0, time.Local)

	} else { //get begining of default tracker range
		today := time.Now()
		commitDate = today.AddDate(0, 0, -371) // -371 days (53 weeks * 7 days)
	}

	// Get next Sunday - first pixel of first full row of the tracker
	daysUntilSunday := (7 - int(commitDate.Weekday())) % 7
	commitDate = commitDate.AddDate(0, 0, daysUntilSunday)

	// Visually represent what will be printed on the git activity tracker
	if cfg.DryRun {
		drawActivity(cfg.OffsetLines, cfg.Text)
	}
	// Always run this, if its a dry run then we're printing out the commit dates
	writeWord(cfg.OffsetLines, cfg.Text, commitDate, cfg.DryRun)
}
