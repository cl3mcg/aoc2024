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

// removeElementAtIndex removes an element from a slice at a specified index.
// It returns a new slice with the element removed.
func removeElementAtIndex(slice []int, index int) []int {
	// If the index is out of range, return the original slice
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Remove the element at the given index by appending the part before and after it
	return append(slice[:index], slice[index+1:]...)
}

// level represents a collection of integer data from the input puzzle.
type level struct {
	data          []int   // Data is a slice of integers representing the current level's values.
	dataTolerance [][]int // DataTolerance is a slice of slices representing the current level's values with one element removed.
}

// initDataTolerance initializes the dataTolerance slice by removing one element at a time from data.
func (l *level) initDataTolerance() {
	// Iterate over the original slice and remove one element at each index
	for i := 0; i < len(l.data); i++ {
		// Create a new slice with the element at index i removed
		modifiedSlice := removeElementAtIndex(append([]int(nil), l.data...), i)
		// Append the modified slice to the result slice
		l.dataTolerance = append(l.dataTolerance, modifiedSlice)
	}
}

// checkIncr checks if the data in the level is strictly increasing with each value
// being greater than the previous one, and the difference between consecutive values is at most 3.
func (l *level) checkIncr() bool {
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
func (l *level) checkDecr() bool {
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

// checkWithTolerance checks if any subsequence of the data is valid by removing one element at a time.
func (l *level) checkWithTolerance() bool {
	// Generate all possible subsequences by removing one level
	l.initDataTolerance()

	// Check each subsequence
	for _, s := range l.dataTolerance {
		// Create a new level object for each subsequence and apply normal checks
		tempLevel := level{data: s}
		if tempLevel.checkIncr() || tempLevel.checkDecr() {
			return true // Safe if any subsequence is valid
		}
	}

	return false // No valid subsequence found
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

	valid := 0 // Start with a count of safe reports

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

		// Check if the report is safe either directly or via tolerance
		if l.checkIncr() || l.checkDecr() {
			valid++
		} else if l.checkWithTolerance() {
			valid++
		}
	}

	// Print the final count of valid levels.
	fmt.Println("The result 'valid' should be: ", valid)
}
