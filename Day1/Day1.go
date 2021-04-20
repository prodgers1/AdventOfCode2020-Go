package main

import (
	"fmt"
	"strconv"
	"advent-of-code-2020/utils"
)


func part1(lines []string, target int) int {
	for _, line := range lines {
		for _, line2 := range lines {
			if line != line2 {
				value1, err := strconv.Atoi(line)
				value2, err2:= strconv.Atoi(line2)

				if err != nil || err2 != nil {
					fmt.Println(err)
					fmt.Println(err2)
				}

				if value1 + value2 == target {
					
					//fmt.Printf("The answer is: %v", value1 * value2)
					return value1 * value2
				}
			}
		}
	}
	return 0
}

func part2(lines []string, target int) int {
	for _, line := range lines {
		for _, line2 := range lines {
			for _, line3 := range lines {
				value1, err := strconv.Atoi(line)
				value2, err2:= strconv.Atoi(line2)
				value3, err3:= strconv.Atoi(line3)

				if err != nil || err2 != nil || err3 != nil {
					fmt.Println(err)
					fmt.Println(err2)
					fmt.Println(err3)
				}

				if value1 + value2 + value3 == target {
					//fmt.Printf("The answer is: %v", value1 * value2 * value3)
					return value1 * value2 * value3
				}
			}
		}
	}
	return 0
}


func main() {
	lines := utils.GetInput("1")

	part1Ans := part1(lines, 2020)
	part2Ans := part2(lines, 2020)

	fmt.Printf("Part 1 Answer: %v \nPart 2 Answer: %v", part1Ans, part2Ans)
	
}