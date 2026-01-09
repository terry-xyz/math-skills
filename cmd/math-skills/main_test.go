package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"simple", []int{1, 2, 3, 4, 5}, 3},
		{"single", []int{10}, 10},
		{"two numbers", []int{1, 2}, 2},
		{"larger numbers", []int{100, 200, 300}, 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateAverage(tt.numbers)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestCalculateMedian(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"odd count", []int{1, 2, 3, 4, 5}, 3},
		{"even count", []int{1, 2, 3, 4}, 3},
		{"single", []int{10}, 10},
		{"unsorted", []int{5, 1, 3, 2, 4}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateMedian(tt.numbers)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestCalculateSampleVariance(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"simple", []int{2, 4, 4, 4, 5, 5, 7, 9}, 5},
		{"identical", []int{5, 5, 5, 5}, 0},
		{"two numbers", []int{0, 10}, 50},
		{"single", []int{5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateSampleVariance(tt.numbers)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestCalculateStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"simple", []int{2, 4, 4, 4, 5, 5, 7, 9}, 2},
		{"identical", []int{5, 5, 5, 5}, 0},
		{"two numbers", []int{0, 10}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateStandardDeviation(tt.numbers)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestReadContent(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []int
	}{
		{"simple", "1\n2\n3\n", []int{1, 2, 3}},
		{"with empty lines", "1\n\n2\n3\n", []int{1, 2, 3}},
		{"windows line endings", "1\r\n2\r\n3\r\n", []int{1, 2, 3}},
		{"single number", "42\n", []int{42}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create temp file
			tmpDir := t.TempDir()
			tmpFile := filepath.Join(tmpDir, "test.txt")
			err := os.WriteFile(tmpFile, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}

			result := readContent(tmpFile)

			if len(result) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, want %d", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("index %d: got %d, want %d", i, result[i], tt.expected[i])
				}
			}
		})
	}
}
