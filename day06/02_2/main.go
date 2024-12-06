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
//   - p (string): The path to the input file.
//
// Returns:
//   - string: The contents of the file as a string.
//   - error: An error if the file cannot be read.
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

// spotFinder finds the first Coord object in the provided slice that has a non-empty direction.
//
// Parameters:
//   - points ([]Coord): A slice of Coord objects to search through.
//
// Returns:
//   - Coord: The first Coord with a non-empty Dir field.
//   - error: An error if no such Coord is found.
func spotFinder(points []Coord) (point Coord, err error) {
	for _, p := range points {
		if p.Dir != "" {
			return p, nil
		}
	}
	return Coord{}, errors.New("spotFinder doesn't find any point")
}

// move performs a single move operation for a given starting point based on its direction.
// The function modifies the state of the points slice to update visited status.
//
// Parameters:
//   - startPt (*Coord): A pointer to the starting Coord object.
//
// Returns:
//   - Coord: The Coord object at the end of the move.
//   - error: An error if the move cannot be completed.
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

// Coord represents a point in the grid with its properties.
type Coord struct {
	X         int    // X-coordinate of the point.
	Y         int    // Y-coordinate of the point.
	Dir       string // Direction of movement ("^", "v", "<", ">", or empty if none).
	IsBlocked bool   // Indicates whether the point is blocked.
	Walked    bool   // Indicates whether the point has been walked on.
}

// switchDir changes the direction of the Coord to the next direction in clockwise order.
// Does nothing if the Coord is blocked or has no direction.
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

var points []Coord // Global slice holding all points in the grid.

// createPlan initializes the global points slice from a 2D array of string representations.
// Each string represents a cell, which may be blocked, empty, or contain a direction.
//
// Parameters:
//   - coords ([][]string): A 2D slice representing the grid.
func createPlan(coords [][]string) {
	points = []Coord{}
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
		}
	}
}

// main is the entry point of the program.
// It initializes the grid, finds the start point, and computes the result based on the movement rules.
func main() {
	// Read the puzzle input from the file "input.txt".
	// If there's an error retrieving the input, log the error and stop execution.
	txt, err := retrievePuzzleInput("../input.txt")
	if err != nil {
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

	createPlan(coords)

	startPoint, err := spotFinder(points)
	if err != nil {
		log.Fatalf("Error finding startPoint spot: %v", err)
	}
	fmt.Printf("startPoint data : x is %v, y is %v, dir is %v\n", startPoint.X, startPoint.Y, startPoint.Dir)

	// Starting the result
	r := 0

	// Loop over each point of the plan
	for i, p := range points {
		// If the point is blocked initially or if the point is the starting point, no need to loop.
		if !p.IsBlocked && !(p.X == startPoint.X && p.Y == startPoint.Y) {

			// Manually switch and block the point to see if this creates the guard to be in an endless path loop.
			points[i].IsBlocked = true

			fmt.Printf("Current value of r: %v | Checking for mutated blocked p with coodinates: %v %v\n", r, p.X, p.Y)

			// Initialize a slice to register the points walked on and their direction.
			var visited []string
			// Save the starting point. To be updated when a move is completed.
			currentPoint := startPoint
			// Start a while loop
			for {
				// Process one move
				endPoint, err := move(&currentPoint)
				// If move() returns an error, the guard moved out of bound, we break the loop to go to the next point
				if err != nil {
					break
				}

				// If not out of bound we store the visited point and the direction in visited
				visitedKey := fmt.Sprintf("%d,%d,%s", endPoint.X, endPoint.Y, endPoint.Dir)

				// If visited already contains the point and the direction, the guard already visited this point and is therefore blocked in a infinite path loop.
				if slices.Contains(visited, visitedKey) {
					// We increment the result and break the loop to move to the next point.
					r++
					break
				}

				// We add the point visited and its direction to the visited map
				visited = append(visited, visitedKey)

				// We set the current point to the endpoint from move()
				currentPoint = endPoint
			}

			// We reset the plan with the coordinates initially provided by the puzzle input
			createPlan(coords)
		}
	}

	// Print the final sum of all valid middle page numbers.
	fmt.Println("The result 'r' should be: ", r)

	fmt.Println("✨ Gracefully exiting the program...")
	os.Exit(0)
}
