package ejercicio4

import (
	"bufio"
	"fmt"
	"os"
)

func Exercise4Part1Solver() int {
	file, err := os.Open("./ejercicio4/input4.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return -1
	}
	defer file.Close()
	var input [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		input = append(input, runes)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
	var total int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if string(input[i][j]) == "X" {
				total += countXmas(input, i, j)
			}
		}
	}
	return total
}

func countXmas(matrix [][]rune, i int, j int) int {
	var result int
	//check row,reverse row
	if j+3 < len(matrix[i]) && string(matrix[i][j+1]) == "M" &&
		string(matrix[i][j+2]) == "A" && string(matrix[i][j+3]) == "S" {
		result++
	}
	if j-3 >= 0 && string(matrix[i][j-1]) == "M" &&
		string(matrix[i][j-2]) == "A" && string(matrix[i][j-3]) == "S" {
		result++
	}
	//check vertically
	if i-3 >= 0 && string(matrix[i-1][j]) == "M" &&
		string(matrix[i-2][j]) == "A" && string(matrix[i-3][j]) == "S" {
		result++
	}
	if i+3 < len(matrix) && string(matrix[i+1][j]) == "M" &&
		string(matrix[i+2][j]) == "A" && string(matrix[i+3][j]) == "S" {
		result++
	}

	//diagonals
	if i+3 < len(matrix) && j+3 < len(matrix[i]) && string(matrix[i+1][j+1]) == "M" &&
		string(matrix[i+2][j+2]) == "A" && string(matrix[i+3][j+3]) == "S" {
		result++
	}
	if j-3 >= 0 && i+3 < len(matrix) && string(matrix[i+1][j-1]) == "M" &&
		string(matrix[i+2][j-2]) == "A" && string(matrix[i+3][j-3]) == "S" {
		result++
	}
	if i-3 >= 0 && j+3 < len(matrix[i]) && string(matrix[i-1][j+1]) == "M" &&
		string(matrix[i-2][j+2]) == "A" && string(matrix[i-3][j+3]) == "S" {
		result++
	}
	if i-3 >= 0 && j-3 >= 0 && string(matrix[i-1][j-1]) == "M" &&
		string(matrix[i-2][j-2]) == "A" && string(matrix[i-3][j-3]) == "S" {
		result++
	}
	return result
}
