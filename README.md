 <h1 align="center">Math Skills</h1>

[![Go](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![CI](https://github.com/terry-xyz/math-skills/workflows/CI/badge.svg)](https://github.com/terry-xyz/math-skills/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/terry-xyz/math-skills?v=1)](https://goreportcard.com/report/github.com/terry-xyz/math-skills)
[![Tests](https://img.shields.io/badge/Tests-Passing-green.svg)](#testing)
[![Makefile](https://img.shields.io/badge/Build-Makefile-orange.svg)](#testing)

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
