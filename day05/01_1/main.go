package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// retrievePuzzleInput reads the content of a file and returns it as a string.
// It returns the file content as a string and an error (if any) that occurred during the file read.
//
// Parameters:
//
//	p (string): The path to the input file.
//
// Returns:
//
//	string: The contents of the file as a string.
//	error: An error if the file cannot be read.
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

// findIndex searches for a target slice inside a matrix of slices and returns its index.
// If the target slice is not found, it returns -1.
//
// Parameters:
//
//	matrix ([][]int): A two-dimensional slice to search within.
//	target ([]int): A slice to find inside the matrix.
//
// Returns:
//
//	int: The index of the target slice in the matrix, or -1 if the target is not found.
func findIndex(matrix [][]int, target []int) int {
	for i, row := range matrix {
		// Use reflect.DeepEqual to compare slices, as slices cannot be directly compared.
		if reflect.DeepEqual(row, target) {
			return i
		}
	}
	// Return -1 if the target slice is not found.
	return -1
}

// main reads the puzzle input, processes it to determine if updates are in the correct order,
// and calculates the sum of middle page numbers for the correct updates.
//
// It processes the page ordering rules and the updates provided, checks whether each update is in the correct order
// based on the rules, and if an update is correct, it sums up the middle page numbers from those updates.
//
// Parameters:
//
//	None
//
// Returns:
//
//	None
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
	lns := strings.Split(txt, "\n")

	// Initialize slices to hold page rules and pages to produce.
	var pRules [][]int
	var pProduce [][]int

	// Process each line from the puzzle input.
	for _, v := range lns {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue // Skip empty lines.
		}
		// If the line contains a rule (e.g., "47|53"), process it.
		if len(v) == 5 {
			arr := strings.Split(v, "|")
			var i []int
			// Convert the split parts of the rule into integers and append them.
			for _, w := range arr {
				d, err := strconv.Atoi(w)
				if err != nil {
					// Log the error if conversion fails.
					log.Fatal("Error converting from string to int", "Value", w, "Error", err)
				}
				i = append(i, d)
			}
			// Add the rule to the pRules slice.
			pRules = append(pRules, i)
			continue
		}
		// If the line contains an update (e.g., "75,47,61,53,29"), process it.
		arr := strings.Split(v, ",")
		var i []int
		// Convert the split parts of the update into integers and append them.
		for _, w := range arr {
			d, err := strconv.Atoi(w)
			if err != nil {
				// Log the error if conversion fails.
				log.Fatal("Error converting from string to int", "Value", w, "Error", err)
			}
			i = append(i, d)
		}
		// Add the update to the pProduce slice.
		pProduce = append(pProduce, i)
	}

	// Initialize a variable to hold the result of the sum of middle page numbers.
	var r int

	// Process each update in pProduce.
	for _, pList := range pProduce {
		var toCheck [][]int
		// Generate all pairs of pages to check the order for.
		for i := 0; i < len(pList); i++ {
			for j := 0; j < len(pList); j++ {
				// Skip pairs where the second index is less than or equal to the first index.
				if j <= i {
					continue
				}
				toCheck = append(toCheck, []int{pList[j], pList[i]})
			}
		}

		// Flag to track if the update passes the ordering check.
		var continuer bool
		// Check if any of the pairs violate the order rules.
		for _, c := range toCheck {
			if findIndex(pRules, c) >= 0 {
				// If any pair matches a rule, mark the update as incorrect.
				continuer = true
				break
			}
		}
		// If any check failed, continue to the next update.
		if continuer {
			continue
		}

		// If the update is in correct order, add the middle page number to the result.
		r = r + pList[(len(pList)-1)/2]
	}

	// Print the final sum of all valid middle page numbers.
	fmt.Println("The result 'r' should be: ", r)
}
