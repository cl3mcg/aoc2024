package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
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

func spotFinder(points []Coord) (point Coord, err error) {
	for _, p := range points {
		if p.Dir != "" {
			return p, nil
		}
	}
	return Coord{}, errors.New("spotFinder doesn't find any point")
}

func move(startPt *Coord) (endPt Coord, err error) {
	if startPt.IsBlocked {
		return Coord{}, errors.New("move: The startPt is blocked")
	}
	if startPt.Dir == "" {
		return Coord{}, errors.New("move: No direction specified in startPt")
	}
	if startPt.Dir == "^" {
		for i, pt := range points {
			if pt.Y == startPt.Y+1 && pt.X == startPt.X {
				if pt.IsBlocked {
					startPt.switchDir()
					return *startPt, nil
				}
				pt.Dir = startPt.Dir
				startPt.Dir = ""
				points[i].Walked = true // Update the Walked field in the points slice
				return pt, nil
			}
		}
		return *startPt, errors.New("no end point found")
	}
	if startPt.Dir == "v" {
		for i, pt := range points {
			if pt.Y == startPt.Y-1 && pt.X == startPt.X {
				if pt.IsBlocked {
					startPt.switchDir()
					return *startPt, nil
				}
				pt.Dir = startPt.Dir
				startPt.Dir = ""
				points[i].Walked = true // Update the Walked field in the points slice
				return pt, nil
			}
		}
		return *startPt, errors.New("no end point found")
	}
	if startPt.Dir == "<" {
		for i, pt := range points {
			if pt.X == startPt.X-1 && pt.Y == startPt.Y {
				if pt.IsBlocked {
					startPt.switchDir()
					return *startPt, nil
				}
				pt.Dir = startPt.Dir
				startPt.Dir = ""
				points[i].Walked = true // Update the Walked field in the points slice
				return pt, nil
			}
		}
		return *startPt, errors.New("no end point found")
	}
	if startPt.Dir == ">" {
		for i, pt := range points {
			if pt.X == startPt.X+1 && pt.Y == startPt.Y {
				if pt.IsBlocked {
					startPt.switchDir()
					return *startPt, nil
				}
				pt.Dir = startPt.Dir
				startPt.Dir = ""
				points[i].Walked = true // Update the Walked field in the points slice
				return pt, nil
			}
		}
		return *startPt, errors.New("no end point found")
	}
	return Coord{}, errors.New("move: No valid direction specified in startPt")
}

type Coord struct {
	X         int
	Y         int
	Dir       string
	IsBlocked bool
	Walked    bool
}

func (c *Coord) switchDir() {
	if !c.IsBlocked && c.Dir != "" {
		switch c.Dir {
		case "^":
			c.Dir = ">"
		case ">":
			c.Dir = "v"
		case "v":
			c.Dir = "<"
		case "<":
			c.Dir = "^"
		}
	}
}

var points []Coord

// main reads the puzzle input, etc...
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

	var coords [][]string

	// Split the input text into lines based on newline characters.
	yRaw := strings.Split(txt, "\n")
	for _, v := range yRaw {
		v = strings.TrimSpace(v)
		p := strings.Split(v, "")
		coords = append(coords, p)
	}

	slices.Reverse(coords)

	for i, y := range coords {
		for j := 0; j < len(y); j++ {
			var pt Coord
			pt.X = j
			pt.Y = i
			if y[j] == "#" {
				pt.IsBlocked = true
			}
			if y[j] != "." && y[j] != "#" {
				pt.Dir = y[j]
			}
			points = append(points, pt)
			fmt.Printf("points is appended pt with values %v: pt X = %v, pt Y = %v\n", pt, pt.X, pt.Y)
		}
	}

	startPoint, err := spotFinder(points)
	if err != nil {
		log.Fatalf("Error finding startPoint spot: %v", err)
	}
	fmt.Printf("startPoint data : x is %v, y is %v, dir is %v\n", startPoint.X, startPoint.Y, startPoint.Dir)

	// Iterate through the moves until an error is thrown
	currentPoint := startPoint
	for {
		endPoint, err := move(&currentPoint) // Perform the move
		if err != nil {
			// If an error occurs, break the loop
			fmt.Printf("Error occurred: %v\n", err)
			break
		}

		// Log the successful move
		fmt.Printf("Moved to endPoint: x = %v, y = %v, dir = %v\n", endPoint.X, endPoint.Y, endPoint.Dir)
		// Update currentPoint with the new position after the move
		currentPoint = endPoint
	}

	// Starting the result at 1 because the starting spot is already walked on.
	r := 1

	for _, pt := range points {
		if pt.Walked {
			r++
		}
	}

	// Print the final sum of all valid middle page numbers.
	fmt.Println("The result 'r' should be: ", r)

	fmt.Println("✨ Gracefully exiting the program...")
	os.Exit(0)
}
