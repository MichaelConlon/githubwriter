# GitHub Writer

A tool for creating custom GitHub contribution patterns by automating git commits. This tool allows you to write text messages in your GitHub activity graph or create work-simulating commit patterns.

## Features

- **Text Mode**: Write messages in your GitHub contribution graph using ASCII-style letters
  - Supports lowercase a-z characters
  - Customizable horizontal offset
  - Can target specific years
  - Preview functionality with dry-run option

- **Work Mode**: Simulate regular work patterns with automated commits
  - Filter files by extension
  - Specify number of files to include
  - Random file selection
  - Customizable commit messages

## Installation
```bash
go install github.com/MichaelConlon/githubwriter@latest
```

## Usage

### Text Mode

Write a message in your GitHub contribution graph:

Options:
- `-text`: The message to write (lowercase a-z only, max 8 characters)
- `-year`: Target year for the contribution pattern
- `-offset`: Number of empty lines before the text (vertical offset)
- `-dry-run`: Preview the pattern without making commits

#### Examples

Write "hello" in your 2024 contribution graph:
```bash
githubwriter -mode text -text hello -year 2024
```

### Work Mode

Create commit patterns that simulate regular work:

Options:
- `-extensions`: Comma-separated list of file extensions to include
- `-num-files`: Number of files to include in each commit
- `-message`: Custom commit message template
- `-ticket`: Add ticket/issue reference to commits

#### Examples

Create a work pattern with 3 files:
```bash
githubwriter -mode work -extensions go -num-files 3
```