package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	input := readFile()
	totalSum := calculateTotal(input)

	fmt.Println("Sum is ", totalSum)
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

func extractHorizontalNumber(grid [][]rune, row, col int) string {
	cols := len(grid[0])

	start := col
	for start > 0 && unicode.IsDigit((grid[row][start-1])) {
		start--
	}
	end := col
	for end < cols-1 && unicode.IsDigit(grid[row][end+1]) {
		end++
	}
	numStr := ""
	for k := start; k <= end; k++ {
		numStr += string(grid[row][k])
	}
	return numStr
}

func calculateTotal(grid [][]rune) int {
	totalSum := 0
	rows := len(grid)
	if rows == 0 {
		return 0
	}

	cols := len(grid[0])

	directions := []struct{ di, dj int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != '*' {
				continue
			}
			//hittat *, letar angränsande siffror
			var foundNums []int

			extracted := make(map[string]bool)

			for _, d := range directions {
				r := i + d.di
				c := j + d.dj

				//om dellen hamnar utanför griden -> fortsätt
				if r < 0 || r >= rows || c < 0 || c >= cols {
					continue
				}
				//Om cellen inte innehåller en siffra -> fortsätt
				if !unicode.IsDigit(grid[r][c]) {
					continue
				}
				//Hämta hela siffersekvensen, sparat som sträng
				numStr := extractHorizontalNumber(grid, r, c)

				//Undvik dubletter
				key := fmt.Sprintf("%d:%s", r, numStr)
				if extracted[key] {
					continue
				}
				extracted[key] = true

				num, err := strconv.Atoi(numStr)
				if err != nil {
					continue
				}
				foundNums = append(foundNums, num)

				if len(foundNums) > 2 {
					break
				}
				if len(foundNums) == 2 {
					totalSum += foundNums[0] * foundNums[1]
				}
			}
		}

	}
	return totalSum
}
