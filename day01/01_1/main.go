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
		// Return an error if the file could not be read.
		return "", err
	}

	// Convert the byte slice to a string and return it.
	return string(d), nil
}

func main() {
	// Read the puzzle input from the file "input.txt".
	txt, err := retrievePuzzleInput("../input.txt")
	if err != nil {
		// Log a fatal error and terminate the program if the input file cannot be read.
		log.Fatalf("Error retrieving the puzzle input: %v", err)
	}

	// Trim any leading or trailing whitespace characters from the input.
	txt = strings.TrimSpace(txt)

	// Split the input string into lines, each representing a pair of numbers.
	l := strings.Split(txt, "\n")

	// Initialize slices to store the left and right integers from each line.
	var cl []int // cl will hold the left values.
	var cr []int // cr will hold the right values.

	// Iterate over each line to process the pairs of numbers.
	for i, v := range l {
		// Split the current line into two parts: left and right, using a single space as a delimiter.
		codes := strings.SplitN(v, " ", 2)

		// Trim spaces from the left and right parts to ensure clean inputs.
		ls := strings.TrimSpace(codes[0]) // Left side as string.
		rs := strings.TrimSpace(codes[1]) // Right side as string.

		// Convert the left string to an integer.
		li, err := strconv.Atoi(ls)
		if err != nil {
			// Log a fatal error if conversion fails, including contextual information for debugging.
			log.Fatalf("Error converting the input to int on the left side. Index %v, Line Value %v, Left Value %v, Error: %v\n", i, v, ls, err)
		}

		// Convert the right string to an integer.
		ri, err := strconv.Atoi(rs)
		if err != nil {
			// Log a fatal error if conversion fails, including contextual information for debugging.
			log.Fatalf("Error converting the input to int on the right side. Index %v, Line Value %v, Right Value %v, Error: %v\n", i, v, rs, err)
		}

		// Append the converted integers to their respective slices.
		cl = append(cl, li) // Left integer added to cl.
		cr = append(cr, ri) // Right integer added to cr.
	}

	// Sort both slices in ascending order to prepare for the next step of comparison.
	slices.Sort(cl)
	slices.Sort(cr)

	// Initialize a variable to accumulate the total difference between corresponding elements.
	var r int

	// Iterate over the sorted left slice and calculate the absolute difference with the corresponding right slice.
	for i, v := range cl {
		// Calculate the difference between the current left and right values.
		d := v - cr[i]
		if d < 0 {
			// If the difference is negative, convert it to positive (absolute value).
			d = -d
		}
		// Accumulate the difference.
		r += d
	}

	// Print the final result, which is the sum of all absolute differences.
	fmt.Println("The result 'r' should be: ", r)
}
