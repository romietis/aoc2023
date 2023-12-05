package main

import (
	"fmt"

	// "strings"
	"github.com/romietis/aocutil"
)

func main() {

	lines, _ := aocutil.ReadFile("example1.txt")

	grid := loadIntoGrid(lines)

	// for _, row := range grid {
	// 	fmt.Println(string(row))
	// }

	// fmt.Println(find_symbol(grid))

	// for x := 0; x<len(grid); x++{
	// 	for y := 0; y<(len(grid[0])); y++{
	// 		fmt.Print(string(grid[x][y]), " ")
	// 	}
	// 	fmt.Println()
	// }

	findSymbol(grid)

}

func findSymbol(grid [][]rune) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < (len(grid[0])); y++ {
			// fmt.Print(string(grid[x][y]), " ")
			if string(grid[x][y]) != "." {
				fmt.Println(string(grid[x][y]), x, y)
			}
		}
	}
}

func loadIntoGrid(lines []string) (grid [][]rune) {
	for _, line := range lines {
		runes := []rune(line)
		grid = append(grid, runes)
	}
	return
}
