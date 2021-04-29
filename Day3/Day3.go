package main
import (
	"fmt"
	//"strings"
	"advent-of-code-2020/utils"
	//"strconv"
)
func main() {
	grid := utils.GetInput("3")
	part1(grid)
	part2(grid)
}
func part1(grid []string) {
	deltaRow := 1
	deltaCol := 3
	trees := traverse(deltaRow, deltaCol, grid)
	fmt.Printf("Part 1 Total trees: %v\n", trees)
}
func part2(grid []string) {
	rowDeltas := []int {1, 1, 1, 1, 2}
	colDeltas := []int {1, 3, 5, 7, 1}
	i := 0
	totalTrees := 1
	for i < 5 {
		deltaRow := rowDeltas[i]
		deltaCol := colDeltas[i]
		trees := traverse(deltaRow, deltaCol, grid)
		totalTrees *= trees
		i += 1
	}
	fmt.Printf("Part 2 Total trees: %v", totalTrees)
}
func traverse(deltaRow int, deltaCol int, grid []string) int {
	row := 0
	col := 0
	trees := 0
	height := len(grid)
	width := len(grid[0])
	for {
		if (row >= height) {
			break
		}
		if (string(grid[row][col]) == "#") {
			trees += 1
		}
		col = (col + deltaCol) % width
		row = (row + deltaRow)
	}
	return trees
}