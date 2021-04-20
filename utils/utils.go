package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(day string) []string {
	filePath := fmt.Sprintf("../Day%v/Input.txt", day)
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
 
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
 
	file.Close()

	return lines
}