package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var maxCubes = map[string]int{"red": 12, "green": 13, "blue": 14}

func main() {

	lines := readFile("data.txt")

	for i, line := range lines {
		if i < 5 { // Visa bara de första 5 raderna
			fmt.Println(line)
		}
	}
	var fixedInput [][]int = cleanInput(lines)
	var totalSum = checkPossibleGames(fixedInput)
	fmt.Println("Summa", totalSum)
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
		fmt.Println("Error reading file: ", err)
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

func checkPossibleGames(gameGrid [][]int) int {

	totalSum := 0

	for _, row := range gameGrid {
		gameId, maxRed, maxBlue, maxGreen := row[0], row[1], row[2], row[3]

		if maxRed <= maxCubes["red"] && maxBlue <= maxCubes["blue"] && maxGreen <= maxCubes["green"] {
			totalSum += gameId
		}

	}
	return totalSum
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting to int:", err)
		return 0
	}
	return num
}
