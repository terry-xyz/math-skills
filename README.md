 <h1 align="center">Math Skills</h1>

A Go program that reads numbers from a file and parses them into a slice for processing.

## Prerequisites

- Go 1.25+
- Make (optional)

## Usage

1. Create a file with numbers (one per line):
   ```
   189
   113
   121
   ```

2. Run the program with the file path as argument:
   ```bash
   go run ./cmd/math-skills <filename>
   ```

   Example:
   ```bash
   go run ./cmd/math-skills data.txt
   ```

## Make Commands

```bash
make build  # Compile binary to bin/
make run    # Build and run with data.txt
make test   # Run all tests
make clean  # Remove build artifacts
```

## Project Structure

```
math-skills/
├── cmd/math-skills/
│   ├── main.go       # Main program
│   └── main_test.go  # Tests
├── data.txt          # Input file
├── Makefile          # Build automation
└── go.mod
```
