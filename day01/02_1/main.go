package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// retrievePuzzleInput reads the puzzle input from a file at the specified path.
// It returns the content of the file as a string or an error if the file cannot be read.
//
// p: The path to the file containing the puzzle input.
// Returns: A string containing the file content, or an error if the file cannot be read.
func retrievePuzzleInput(p string) (string, error) {
	// Read the entire content of the file into a byte slice.
	d, err := os.ReadFile(p)
	if err != nil {
		// Return an error if reading the file fails.
		return "", err
	}

	// Convert the byte slice to a string and return it.
	return string(d), nil
}

// countOccurrences counts how many times a specific value appears in a slice of integers.
// It returns the count of occurrences.
func countOccurrences(slice []int, value int) int {
	var count int
	// Iterate through the slice and increment the count for each match.
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

func main() {
	// Read the puzzle input from the file "input.txt".
	txt, err := retrievePuzzleInput("../input.txt")
	if err != nil {
		// Log a fatal error and terminate the program if the input cannot be retrieved.
		log.Fatalf("Error retrieving the puzzle input: %v", err)
	}

	// Trim any leading or trailing whitespace characters from the input.
	txt = strings.TrimSpace(txt)

	// Split the input string into lines, each representing a pair of numbers.
	l := strings.Split(txt, "\n")

	// Initialize slices to store the left and right integers from each line.
	var cl []int // cl will hold the left values.
	var cr []int // cr will hold the right values.

	// Process each line to extract and convert the pair of numbers.
	for i, v := range l {
		// Split the current line into two parts: left and right, using a single space as a delimiter.
		codes := strings.SplitN(v, " ", 2)

		// Trim spaces from the left and right parts to ensure clean input.
		ls := strings.TrimSpace(codes[0]) // Left side as string.
		rs := strings.TrimSpace(codes[1]) // Right side as string.

		// Convert the left string to an integer.
		li, err := strconv.Atoi(ls)
		if err != nil {
			// Log a fatal error if conversion fails, providing context for debugging.
			log.Fatalf("Error converting the input to int on the left side. Index %v, Line Value %v, Left Value %v, Error: %v\n", i, v, ls, err)
		}

		// Convert the right string to an integer.
		ri, err := strconv.Atoi(rs)
		if err != nil {
			// Log a fatal error if conversion fails, providing context for debugging.
			log.Fatalf("Error converting the input to int on the right side. Index %v, Line Value %v, Right Value %v, Error: %v\n", i, v, rs, err)
		}

		// Append the converted integers to their respective slices.
		cl = append(cl, li) // Left integer added to cl.
		cr = append(cr, ri) // Right integer added to cr.
	}

	// Sort both slices in ascending order for consistent comparison.
	slices.Sort(cl)
	slices.Sort(cr)

	// Initialize a variable to accumulate the total weighted sum of matches.
	var r int

	// Iterate over the sorted left slice.
	for _, v := range cl {
		// Check if the current left value exists in the right slice.
		if !slices.Contains(cr, v) {
			// If not, skip this value.
			continue
		}

		// Count how many times the current value appears in the right slice.
		t := countOccurrences(cr, v)

		// Accumulate the product of the value and its count in the result.
		r += t * v
	}

	// Print the final result, which represents the weighted sum of matching values.
	fmt.Println("The result 'r' should be: ", r)
}
