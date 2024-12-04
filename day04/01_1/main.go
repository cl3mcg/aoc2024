package main

import (
	"fmt"
	"log"
	"os"
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

// Input represents the overall structure of the puzzle, which contains multiple lines of content.
type Input struct {
	Content []Line // A slice of Line objects, each representing a single line of the puzzle.
}

// Line represents a single line in the puzzle, containing its index and content as a slice of strings.
type Line struct {
	Index   int      // The index of the line in the puzzle grid.
	Content []string // A slice of strings representing individual characters in the line.
}

// check0000 checks for a specific pattern.
// If the coordinates of the character is the center of a clock, check0000 checks the values at "00:00"
//
// input: The puzzle input.
// indexL: The line index to start checking from.
// indexC: The column index to start checking from.
// Returns: True if the pattern is found, false otherwise.
func check0000(input Input, indexL, indexC int) bool {
	// Ensure indexL is at least 3 and indexC is within bounds
	if indexL < 3 {
		return false
	}

	if input.Content[indexL-1].Content[indexC] != "M" {
		return false
	}
	if input.Content[indexL-2].Content[indexC] != "A" {
		return false
	}
	if input.Content[indexL-3].Content[indexC] != "S" {
		return false
	}

	return true
}

// check0130 checks for a specific pattern in the input starting from the given indices.
// If the coordinates of the character is the center of a clock, check0130 checks the values at "01:30"
//
// input: The puzzle input.
// line: The current line being checked.
// indexL: The line index to start checking from.
// indexC: The column index to start checking from.
// Returns: True if the pattern is found, false otherwise.
func check0130(input Input, line Line, indexL, indexC int) bool {
	// Ensure indexL is at least 3 and indexC is within bounds
	if indexL < 3 || indexC+3 >= len(line.Content) {
		return false
	}

	if input.Content[indexL-1].Content[indexC+1] != "M" {
		return false
	}
	if input.Content[indexL-2].Content[indexC+2] != "A" {
		return false
	}
	if input.Content[indexL-3].Content[indexC+3] != "S" {
		return false
	}

	return true
}

// check0300 checks for a specific pattern in the input starting from the given indices.
// If the coordinates of the character is the center of a clock, check0300 checks the values at "03:00"
//
// input: The puzzle input.
// line: The current line being checked.
// indexL: The line index to start checking from.
// indexC: The column index to start checking from.
// Returns: True if the pattern is found, false otherwise.
func check0300(input Input, line Line, indexL, indexC int) bool {
	if indexC+3 >= len(line.Content) {
		return false
	}

	if input.Content[indexL].Content[indexC+1] != "M" {
		return false
	}
	if input.Content[indexL].Content[indexC+2] != "A" {
		return false
	}
	if input.Content[indexL].Content[indexC+3] != "S" {
		return false
	}

	return true
}

// check0430 checks for a specific pattern in the input starting from the given indices.
// If the coordinates of the character is the center of a clock, check0300 checks the values at "04:30"
//
// input: The puzzle input.
// line: The current line being checked.
// indexL: The line index to start checking from.
// indexC: The column index to start checking from.
// Returns: True if the pattern is found, false otherwise.
func check0430(input Input, line Line, indexL, indexC int) bool {
	if indexL+3 >= len(input.Content) || indexC+3 >= len(line.Content) {
		return false
	}

	if input.Content[indexL+1].Content[indexC+1] != "M" {
		return false
	}
	if input.Content[indexL+2].Content[indexC+2] != "A" {
		return false
	}
	if input.Content[indexL+3].Content[indexC+3] != "S" {
		return false
	}

	return true
}

// check0600 checks for a specific pattern in the input starting from the given indices.
// If the coordinates of the character is the center of a clock, check0600 checks the values at "06:00"
//
// input: The puzzle input.
// indexL: The line index to start checking from.
// indexC: The column index to start checking from.
// Returns: True if the pattern is found, false otherwise.
func check0600(input Input, indexL, indexC int) bool {
	if indexL+3 >= len(input.Content) {
		return false
	}

	if input.Content[indexL+1].Content[indexC] != "M" {
		return false
	}
	if input.Content[indexL+2].Content[indexC] != "A" {
		return false
	}
	if input.Content[indexL+3].Content[indexC] != "S" {
		return false
	}

	return true
}

// check0730 checks for a specific pattern in the input starting from the given indices.
// If the coordinates of the character is the center of a clock, check0730 checks the values at "07:30"
//
// input: The puzzle input.
// indexL: The line index to start checking from.
// indexC: The column index to start checking from.
// Returns: True if the pattern is found, false otherwise.
func check0730(input Input, indexL, indexC int) bool {
	if indexL+3 >= len(input.Content) || indexC < 3 {
		return false
	}

	if input.Content[indexL+1].Content[indexC-1] != "M" {
		return false
	}
	if input.Content[indexL+2].Content[indexC-2] != "A" {
		return false
	}
	if input.Content[indexL+3].Content[indexC-3] != "S" {
		return false
	}

	return true
}

// check0900 checks for a specific pattern in the input starting from the given indices.
// If the coordinates of the character is the center of a clock, check0900 checks the values at "09:00"
//
// input: The puzzle input.
// indexL: The line index to start checking from.
// indexC: The column index to start checking from.
// Returns: True if the pattern is found, false otherwise.
func check0900(input Input, indexL, indexC int) bool {
	if indexC < 3 {
		return false
	}

	if input.Content[indexL].Content[indexC-1] != "M" {
		return false
	}
	if input.Content[indexL].Content[indexC-2] != "A" {
		return false
	}
	if input.Content[indexL].Content[indexC-3] != "S" {
		return false
	}

	return true
}

// check1030 checks for a specific pattern in the input starting from the given indices.
// If the coordinates of the character is the center of a clock, check1030 checks the values at "10:30"
//
// input: The puzzle input.
// indexL: The line index to start checking from.
// indexC: The column index to start checking from.
// Returns: True if the pattern is found, false otherwise.
func check1030(input Input, indexL, indexC int) bool {
	if indexL < 3 || indexC < 3 {
		return false
	}

	if input.Content[indexL-1].Content[indexC-1] != "M" {
		return false
	}
	if input.Content[indexL-2].Content[indexC-2] != "A" {
		return false
	}
	if input.Content[indexL-3].Content[indexC-3] != "S" {
		return false
	}

	return true
}

// main reads the puzzle input from a file, processes it, and counts occurrences of the word "XMAS"
// in all 8 possible directions in the puzzle grid. It prints the final count.
func main() {
	// Read the puzzle input from the file "input.txt".
	// If there's an error retrieving the input, log the error and stop execution.
	txt, err := retrievePuzzleInput("../input.txt")
	if err != nil {
		// Log a fatal error and terminate the program if the input file cannot be read.
		log.Fatalf("Error retrieving the puzzle input: %v", err)
	}

	// Trim any leading or trailing whitespace characters from the input to clean it.
	txt = strings.TrimSpace(txt)

	// Split the input text into lines based on newline characters.
	ls := strings.Split(txt, "\n")

	// Initialize an Input object to store the puzzle content.
	var inputLines Input

	// Process each line of the puzzle and add it to the inputLines.Content slice.
	for i, v := range ls {
		v = strings.TrimSpace(v)
		v = strings.Trim(v, "\n")
		v = strings.Trim(v, "\r")
		var l Line
		l.Content = strings.SplitN(v, "", -1)
		l.Index = i
		inputLines.Content = append(inputLines.Content, l)
	}

	// Initialize a variable to count the occurrences of "XMAS".
	var r int

	// Iterate through each line and each character in the line to check for "XMAS".
	for i, v := range inputLines.Content {
		for j, w := range v.Content {
			if strings.ToUpper(w) == "X" {
				// Check if the "X" character can form the word "XMAS" in any of the 8 directions.
				if check0000(inputLines, i, j) {
					r++
				}
				if check0130(inputLines, v, i, j) {
					r++
				}
				if check0300(inputLines, v, i, j) {
					r++
				}
				if check0430(inputLines, v, i, j) {
					r++
				}
				if check0600(inputLines, i, j) {
					r++
				}
				if check0730(inputLines, i, j) {
					r++
				}
				if check0900(inputLines, i, j) {
					r++
				}
				if check1030(inputLines, i, j) {
					r++
				}
			}
		}
	}

	// Print the final sum of all valid multiplications.
	fmt.Println("The result 'r' should be: ", r)
}
