package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	numbers := readContent("../../data.txt")

	average := calculateAverage(numbers)
	fmt.Println("Average: ", average)

	median := calculateMedian(numbers)
	fmt.Println("Median: ", median)

	variance := calculateSampleVariance(numbers)
	fmt.Println("Variance: ", variance)
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

func calculateMedian(numbers []int) int {
	temp := make([]int, len(numbers)) // create a temporary slice to sort the numbers
	copy(temp, numbers)               // copy the numbers to the temporary slice
	slices.Sort(temp)                 // sort the numbers
	mid := len(temp) / 2              // get the middle index
	if len(temp)%2 == 0 {
		sum := temp[mid-1] + temp[mid] // get the sum of the two middle numbers
		return (sum + 1) / 2           // round up to the nearest integer
	}
	return temp[mid] // return the middle number
}

func calculateSampleVariance(numbers []int) int {

	if len(numbers) < 2 {
		return 0 // Variance requires at least 2 numbers
	}

	average := calculateAverage(numbers) // calculate the average of the numbers
	sumSquaredDiffs := 0
	for _, num := range numbers {
		diff := num - average  // calculate the difference between the number and the average
		squared := diff * diff // calculate the squared difference
		if sumSquaredDiffs > math.MaxInt64-squared {
			log.Fatalf("Data overflow: The sum of squares is too large for an integer.")
			os.Exit(1)
		}
		sumSquaredDiffs += squared // add the squared difference to the sum of squared differences
	}

	divisor := len(numbers) - 1
	// Integer rounding logic: (sum + (divisor/2)) / divisor
	return (sumSquaredDiffs + (divisor / 2)) / divisor
}
