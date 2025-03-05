package types

// CMD args
type Config struct {
	Mode     string
	DryRun   bool
	TextArgs TextArgs
	WorkArgs WorkArgs
}

type TextArgs struct {
	Year        int
	Text        string
	OffsetLines int
}

type WorkArgs struct {
	Date       string
	Ticket     string
	Message    string
	Extensions string
	NumFiles   int
	Files      string
	LintConfig string
}

type Commit struct {
	Date    string
	Ticket  string
	Message string
	Files   []string
}

type Letter [5][5]int
