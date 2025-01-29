package config

// CMD args
type Config struct {
	Mode        string
	Text        string
	OffsetLines int
	DryRun      bool
}

type Letter [5][5]int
