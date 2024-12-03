package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"regexp"
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
	// If there's an error retrieving the input, log the error and stop execution.
	txt, err := retrievePuzzleInput("../input.txt")
	if err != nil {
		// Log a fatal error and terminate the program if the input file cannot be read.
		log.Fatalf("Error retrieving the puzzle input: %v", err)
	}

	// Trim any leading or trailing whitespace characters from the input to clean it.
	txt = strings.TrimSpace(txt)

	// Split the input string by the keyword "mul", so that we can process each segment separately.
	e := strings.SplitN(txt, "mul", -1)

	// Initialize a slice to store valid "mul" instructions.
	var f []string

	// Loop through each segment and check if it contains a valid "mul" instruction.
	for _, v := range e {
		// Use a regular expression to check if the segment contains a valid "mul(X,Y)" pattern.
		match, _ := regexp.MatchString(`\(\d{1,4},\d{1,4}\)`, v)
		// If the pattern matches, add the segment to the list of valid instructions.
		if match {
			f = append(f, v)
		}
	}

	// Initialize a variable to hold the sum of all multiplication results.
	var r int

	// Loop through each valid "mul" instruction in the list.
	for _, v := range f {
		// Skip the instruction if it doesn't start with a parenthesis.
		if v[0] != '(' {
			continue
		}
		// Extract the string inside the parentheses by cutting off the part before "(".
		_, bf, _ := strings.Cut(v, "(")
		// Extract the content inside the parentheses, removing the closing ")".
		af, _, _ := strings.Cut(bf, ")")
		// Split the string inside the parentheses by the comma to separate the two numbers.
		sl := strings.Split(af, ",")
		// Convert the first number to an integer.
		n1, err := strconv.Atoi(sl[0])
		if err != nil {
			// Log an error if parsing the first number fails.
			slog.Error("Error parsing string to int", "v", v, "af", af, "sl[0]", sl[0])
			continue
		}
		// Convert the second number to an integer.
		n2, err := strconv.Atoi(sl[1])
		if err != nil {
			// Log an error if parsing the second number fails.
			slog.Error("Error parsing string to int", "v", v, "af", af, "sl[1]", sl[1])
			continue
		}
		// Multiply the two numbers and add the result to the total sum.
		r = r + (n1 * n2)
	}

	// Print the final sum of all valid multiplications.
	fmt.Println("The result 'r' should be: ", r)
}
