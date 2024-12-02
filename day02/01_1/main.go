package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// retrievePuzzleInput reads the puzzle input from a file at the specified path.
// It returns the content of the file as a string or an error if the file cannot be read.
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

// level represents a collection of integer data from the input puzzle.
type level struct {
	data []int // Data is a slice of integers representing the current level's values.
}

// checkIncr checks if the data in the level is strictly increasing with each value
// being greater than the previous one, and the difference between consecutive values is at most 3.
func (l level) checkIncr() bool {
	for i, v := range l.data {
		// Skip the first value as there is no previous value to compare.
		if i == 0 {
			continue
		}
		// If the current value is greater than the previous and the difference is <= 3, continue checking.
		if v > l.data[i-1] && v-l.data[i-1] <= 3 {
			continue
		} else {
			// Return false if the sequence is not increasing or the difference is greater than 3.
			return false
		}
	}
	// Return true if the entire sequence satisfies the conditions.
	return true
}

// checkDecr checks if the data in the level is strictly decreasing with each value
// being smaller than the previous one, and the difference between consecutive values is at most 3.
func (l level) checkDecr() bool {
	for i, v := range l.data {
		// Skip the first value as there is no previous value to compare.
		if i == 0 {
			continue
		}
		// If the current value is smaller than the previous and the difference is <= 3, continue checking.
		if v < l.data[i-1] && l.data[i-1]-v <= 3 {
			continue
		} else {
			// Return false if the sequence is not decreasing or the difference is greater than 3.
			return false
		}
	}
	// Return true if the entire sequence satisfies the conditions.
	return true
}

func main() {
	// Read the puzzle input from the file "input.txt".
	txt, err := retrievePuzzleInput("../input.txt")
	if err != nil {
		// Log a fatal error and terminate the program if the input file cannot be read.
		log.Fatalf("Error retrieving the puzzle input: %v", err)
	}

	// Trim any leading or trailing whitespace characters from the input to clean it.
	txt = strings.TrimSpace(txt)

	// Split the input string into lines, with each line representing a group of numbers.
	r := strings.Split(txt, "\n")

	// Initialize the 'valid' counter with the number of lines in the input.
	valid := len(r)

	// Iterate over each line in the input.
	for _, v := range r {
		// Remove any carriage return characters that may be present.
		v = strings.Replace(v, "\r", "", -1)

		// Trim any extra whitespace from the line.
		strings.TrimSpace(v)

		// Split the line into individual string values (representing numbers).
		ds := strings.Split(v, " ")

		// Create a new 'level' instance to store the parsed data.
		l := level{}

		// Convert each string value to an integer and append it to the 'level' data.
		for _, w := range ds {
			di, err := strconv.Atoi(w)
			if err != nil {
				// Log an error and terminate the program if a string cannot be converted to an integer.
				log.Fatalf("Error converting %s to int: %v", w, err)
			}
			l.data = append(l.data, di)
		}

		// Check if the data in the level is neither increasing nor decreasing in a valid manner.
		// If both checks fail, decrement the 'valid' counter.
		if !l.checkIncr() && !l.checkDecr() {
			valid--
		}
	}

	// Print the final count of valid levels.
	fmt.Println("The result 'valid' should be: ", valid)
}
