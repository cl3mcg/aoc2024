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

// check checks if the word "MAS" is in the form of an X (both forwards and backwards) around the given position.
// The check looks at the surrounding characters to determine if the current position forms part of an "X-MAS".
//
// input: The entire puzzle content.
// line: The current line being processed.
// indexL: The index of the current line in the puzzle.
// indexC: The index of the current character in the line.
// Returns: true if the surrounding characters form a valid "X-MAS", false otherwise.
func check(input Input, line Line, indexL, indexC int) bool {
	// Ensure indexL is not at the first or last position, and indexC is not at the boundaries of the puzzle.
	if indexL == 0 || indexL == len(line.Content)-1 || indexC == 0 || indexC == len(line.Content)-1 {
		return false
	}

	// Check if the surrounding characters are valid candidates for an X-MAS pattern (either 'M' or 'S').
	if input.Content[indexL-1].Content[indexC-1] != "M" && input.Content[indexL-1].Content[indexC-1] != "S" {
		return false
	}
	if input.Content[indexL+1].Content[indexC-1] != "M" && input.Content[indexL+1].Content[indexC-1] != "S" {
		return false
	}
	if input.Content[indexL-1].Content[indexC+1] != "M" && input.Content[indexL-1].Content[indexC+1] != "S" {
		return false
	}
	if input.Content[indexL+1].Content[indexC+1] != "M" && input.Content[indexL+1].Content[indexC+1] != "S" {
		return false
	}

	// Check if the specific arrangement of 'M' and 'S' form the correct pattern for an X-MAS.
	if input.Content[indexL-1].Content[indexC-1] == "M" && input.Content[indexL+1].Content[indexC+1] != "S" {
		return false
	}
	if input.Content[indexL-1].Content[indexC-1] == "S" && input.Content[indexL+1].Content[indexC+1] != "M" {
		return false
	}
	if input.Content[indexL-1].Content[indexC+1] == "M" && input.Content[indexL+1].Content[indexC-1] != "S" {
		return false
	}
	if input.Content[indexL-1].Content[indexC+1] == "S" && input.Content[indexL+1].Content[indexC-1] != "M" {
		return false
	}

	return true
}

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

	// Split the input into lines based on newline characters.
	ls := strings.Split(txt, "\n")

	// Initialize an Input object to store the puzzle content as lines.
	var inputLines Input

	// Process each line of the input text.
	// Each line is split into individual characters and added to the content of the puzzle.
	for i, v := range ls {
		v = strings.TrimSpace(v)
		v = strings.Trim(v, "\n")
		v = strings.Trim(v, "\r")
		var l Line
		// Split each line into individual characters and store it in the Line object.
		l.Content = strings.SplitN(v, "", -1)
		// Store the line index and the content.
		l.Index = i
		inputLines.Content = append(inputLines.Content, l)
	}

	// Initialize a variable to store the count of valid "X-MAS" patterns.
	var r int

	// Iterate over each line and each character in the line to check for "X-MAS" patterns.
	for i, v := range inputLines.Content {
		for j, w := range v.Content {
			// Check if the character is 'A' (since "MAS" starts with 'A').
			if strings.ToUpper(w) == "A" {
				// If the character is 'A', check if it forms part of a valid "X-MAS".
				if check(inputLines, v, i, j) {
					// Increment the count if the check is successful.
					r++
				}
			}
		}
	}

	// Print the final count of valid "X-MAS" patterns found in the puzzle.
	fmt.Println("The result 'r' should be: ", r)
}
