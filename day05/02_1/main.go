package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"reflect"
	"slices"
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

// Function to find the index of a specific []int inside a [][]int
func findIndex(matrix [][]int, target []int) int {
	for i, row := range matrix {
		if reflect.DeepEqual(row, target) {
			return i
		}
	}
	return -1 // Return -1 if the target slice is not found
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
	lns := strings.Split(txt, "\n")

	var pRules [][]int
	var pProduce [][]int

	// Parse the puzzle input into rules and pages to produce.
	for _, v := range lns {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		if len(v) == 5 {
			arr := strings.Split(v, "|")
			var i []int
			for _, w := range arr {
				d, err := strconv.Atoi(w)
				if err != nil {
					slog.Error("Error converting from string to int", "Value", w, "Error", err)
				}
				i = append(i, d)
			}
			pRules = append(pRules, i)
			continue
		}
		arr := strings.Split(v, ",")
		var i []int
		for _, w := range arr {
			d, err := strconv.Atoi(w)
			if err != nil {
				slog.Error("Error converting from string to int", "Value", w, "Error", err)
			}
			i = append(i, d)
		}
		pProduce = append(pProduce, i)
	}

	var r int

	// Loop over each pList
	for _, pList := range pProduce {
		// Initially check if the list is valid
		initiallyValid := true

		// List of checks to perform for every pList
		var toCheck [][]int
		for i := 0; i < len(pList); i++ {
			for j := 0; j < len(pList); j++ {
				if j <= i {
					continue
				}
				toCheck = append(toCheck, []int{pList[j], pList[i]})
			}
		}

		// Check if any of the pairs violate the order rules
		for _, c := range toCheck {
			if findIndex(pRules, c) < 0 {
				// If any pair do not match a rule, mark the update as correct.
				break
			}
			// Mark the list as invalid and perform the necessary updates
			initiallyValid = false
			// Swap the elements in the list to match the rule
			pListI1 := slices.Index(pList, c[1])
			pListI0 := slices.Index(pList, c[0])
			pList[pListI1] = c[0]
			pList[pListI0] = c[1]
		}

		// If the list is now valid and was modified, add the middle page number to `r`
		if !initiallyValid {
			// Only add to `r` if the pList was not valid initially and now is fixed.
			r += pList[(len(pList)-1)/2]
		}
	}

	// Print the final sum of all valid multiplications.
	fmt.Println("The result 'r' should be: ", r)
}
