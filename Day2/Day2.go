package main

import (
	"fmt"
	"strings"
	"advent-of-code-2020/utils"
	"strconv"

)

func parseLine(line string) (int, int, string, string) {
	
	splitLine := strings.Split(line, ":") 
	password := strings.TrimSpace(splitLine[1])
	occurrencesAndSearch := strings.Split(splitLine[0], " ")
	
	toSearch := strings.TrimSpace(occurrencesAndSearch[1])
	
	occurrences := strings.Split(occurrencesAndSearch[0], "-")

	min, err := strconv.Atoi(occurrences[0])
	max, err2 := strconv.Atoi(occurrences[1])

	if err != nil || err2 != nil {
		panic("AHHHH")
	}
	
	return min, max, toSearch, password
}

func part1(lines []string) {
	validPasswords := []string{}
	for _, line := range lines {
		min, max, toSearch, password := parseLine(line)

		//fmt.Println(password)

		count :=  strings.Count(password, toSearch)

		if count >= min && count <= max {
			validPasswords = append(validPasswords, strings.TrimSpace(password))
		}
	}

	fmt.Println(len(validPasswords))

}

func part2(lines []string) {
	validPasswords := []string{}
	for _, line := range lines {
		min, max, toSearch, password := parseLine(line)

		//fmt.Println(password)

		minChar := string(password[min-1])
		maxChar := string(password[max-1])

		if (minChar == toSearch || maxChar == toSearch) && !(minChar == toSearch && maxChar == toSearch){
			validPasswords = append(validPasswords, strings.TrimSpace(password))
		}
	}

	fmt.Println(len(validPasswords))

}

func main() {
	lines := utils.GetInput("2")
	
	part1(lines)
	part2(lines)
}