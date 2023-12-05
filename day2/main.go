package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/romietis/aocutil"
)

func main() {
	games, err := aocutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	total := 0
	for _, game := range games {
		_, sets := parseGame(game)
		// if isPossible(sets) {
		// 	total += id
		// }

		temp := maxValues(sets)
		total += multiplyMaxValues(temp)

	}

	fmt.Println(total)
}

func parseGame(game string) (int, []string) {
	parts := strings.Split(game, ": ")
	idPart := parts[0]
	id, _ := strconv.Atoi(strings.Split(idPart, " ")[1])
	sets := strings.Split(parts[1], "; ")

	return id, sets
}

func isPossible(sets []string) bool {
	for _, set := range sets {
		cubes := strings.Split(set, ", ")
		for _, cube := range cubes {
			colourCount := strings.Split(cube, " ")
			count, _ := strconv.Atoi(colourCount[0])
			colour := colourCount[1]

			if (colour == "red" && count > 12) || (colour == "green" && count > 13) || (colour == "blue" && count > 14) {
				return false
			}
		}
	}

	return true
}

func maxValues(sets []string) map[string]int {
	max_values := make(map[string]int)
	for _, set := range sets {
		cubes := strings.Split(set, ", ")
		for _, cube := range cubes {
			colourCount := strings.Split(cube, " ")
			count, _ := strconv.Atoi(colourCount[0])
			colour := colourCount[1]

			if value, key := max_values[colour]; key {
				if count > value {
					max_values[colour] = count
				}
			} else {
				max_values[colour] = count
			}
		}
	}
	return max_values
}

func multiplyMaxValues(max_values map[string]int) int {
	multiplication_result := 1
	for _, value := range max_values {
		multiplication_result *= value
	}
	return multiplication_result
}
