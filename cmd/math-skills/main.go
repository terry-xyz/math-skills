package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := readContent("../../data.txt")

	average := calculateAverage(numbers)
	fmt.Println("Average: ", average)
}

func readContent(path string) []int {
	// read the file content
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
		os.Exit(1)
	}
	// split the content into lines
	lines := strings.Split(string(content), "\n")
	// create a slice of integers
	numbers := []int{}
	for _, line := range lines {
		line = strings.TrimSpace(line) // trim the space from the line
		if line == "" {
			continue // skip the empty line
		}
		num, err := strconv.Atoi(line) // convert the line to an integer
		if err != nil {
			log.Fatalf("failed to convert string to int: %v", err)
			os.Exit(1)
		}
		numbers = append(numbers, num)
	}
	// return the numbers as a slice of integers
	return numbers
}

func calculateAverage(numbers []int) int {
	// split the content into lines
	total := 0
	for _, num := range numbers {
		total += num
	}
	return (total + (len(numbers) / 2)) / len(numbers) // round up to the nearest integer
}
