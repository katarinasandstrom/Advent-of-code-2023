package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var grid [][]rune = readFile()
	totSum := makeNums(grid)
	fmt.Printf("Total sum is %d\n", totSum)

}

func readFile() [][]rune {
	var grid [][]rune

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return nil
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return grid
}

func makeNums(grid [][]rune) int {

	var firstNum int
	var lastNum int
	var totSum int

	for i := 0; i < len(grid); i++ {
		isFirst := false
		isLast := false
		line := grid[i]
		firstNum = 0
		lastNum = 0

		for j := 0; j < len(line); j++ {
			if line[j] >= '0' && line[j] <= '9' {
				num := int(line[j] - '0')
				if !isFirst {
					firstNum = num
					isFirst = true
				} else {
					lastNum = num
					isLast = true
				}
				lastNum = num
			}
		}

		if !isLast {
			lastNum = firstNum
		}

		if isFirst {
			partSum := firstNum*10 + lastNum
			totSum += partSum
			fmt.Println(firstNum, lastNum, "->", partSum)
		}
	}
	return totSum
}
