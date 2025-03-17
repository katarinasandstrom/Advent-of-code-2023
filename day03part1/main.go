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

func calculateTotal(grid [][]rune) int {
	totalNum := 0
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			if unicode.IsDigit(grid[i][j]) {
				numStr := ""
				k := j
				for k < cols && unicode.IsDigit(grid[i][k]) {
					numStr += string(grid[i][k])
					k++
				}
				num, err := strconv.Atoi(numStr)
				if err != nil {
					j = k - 1
					continue
				}
				validNum := false

				for x := j; x < k; x++ {
					for di := -1; di <= 1; di++ {
						for dj := -1; dj <= 1; dj++ {
							if di == 0 && dj == 0 {
								continue
							}
							ni := i + di
							nj := x + dj

							if ni < 0 || ni >= rows || nj < 0 || nj >= cols {
								continue
							}
							//om intilliggande cell inte är '.' och inte är siffra, sätt validNum = true
							if grid[ni][nj] != '.' && !unicode.IsDigit(grid[ni][nj]) {
								validNum = true
							}
						}
					}
				}
				if validNum {
					totalNum += num
				}
				j = k - 1
			}
		}
	}
	return totalNum
}
