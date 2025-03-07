package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var input []string = readFile("data.txt")
	var cleanedInput [][]int = cleanInput(input)
	var totalSum int = calculate(cleanedInput)
	fmt.Println("Total sum is ", totalSum)

}

func readFile(inputData string) []string {

	var input []string

	file, err := os.Open(inputData)
	if err != nil {
		fmt.Println("Error opening file", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		return nil
	}
	return input
}

func cleanInput(lines []string) [][]int {
	var gameGrid [][]int
	gameRe := regexp.MustCompile(`Game (\d+)`)
	colorRe := regexp.MustCompile(`(\d+)\s+(red|blue|green)`)

	for _, line := range lines {

		gameMatch := gameRe.FindStringSubmatch(line)
		if len(gameMatch) < 2 {
			fmt.Println("Could not get gameId from:", line)
			continue
		}
		gameId := toInt(gameMatch[1])
		maxValues := map[string]int{"red": 0, "blue": 0, "green": 0}

		matches := colorRe.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			amount := toInt(match[1])
			color := match[2]

			if amount > maxValues[color] {
				maxValues[color] = amount
			}
		}
		gameGrid = append(gameGrid, []int{gameId, maxValues["red"], maxValues["blue"], maxValues["green"]})
	}
	return gameGrid
}

func calculate(gameGrid [][]int) int {
	var totalSum int = 0

	for _, row := range gameGrid {

		if len(row) < 4 {
			fmt.Println("Error in row in gameGrid:", row)
			continue
		}
		maxRed, maxBlue, maxGreen := row[1], row[2], row[3]
		partSum := maxRed * maxBlue * maxGreen
		totalSum += partSum
	}
	return totalSum

}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting to int", err)
		return 0
	}
	return num
}
