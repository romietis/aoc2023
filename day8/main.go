package main

import (
	"fmt"
	"strings"

	"github.com/romietis/aocutil"
)

func main() {

	// Load directions
	file, _ := aocutil.ReadFile("input.txt")
	directions := file[0]

	// Load destinations
	destinations := make(map[string]destination)
	currentNodesEndWithA := make(map[string]destination)

	for i := 2; i < len(file); i++ {
		coordinatePart := strings.Split(file[i], " = ")[0]
		leftRightPart := strings.Trim(strings.Split(file[i], " = ")[1], "()")
		left := strings.Split(leftRightPart, ", ")[0]
		right := strings.Split(leftRightPart, ", ")[1]

		if strings.HasSuffix(coordinatePart, "A") {
			a := destination{coordinatePart, right, left}
			currentNodesEndWithA[coordinatePart] = a
		}

		d := destination{coordinatePart, right, left}
		destinations[coordinatePart] = d
	}

	// Part 1
	var counter int
	currentNode := destinations["AAA"]
	for currentNode != destinations["ZZZ"] {
		for _, direction := range directions {
			if string(direction) == "R" {
				currentNode = destinations[currentNode.right]
				counter++
			}
			if string(direction) == "L" {
				currentNode = destinations[currentNode.left]
				counter++
			}
		}
	}
	fmt.Println(counter)
}

type destination struct {
	coordinate string
	right      string
	left       string
}
