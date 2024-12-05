package ejercicio4

import (
	"bufio"
	"fmt"
	"os"
)

func Exercise4Part2Solver() int {
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
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if string(input[i][j]) == "A" {
				total += countXMasMas(input, i, j)
			}
		}
	}
	return total
}

func countXMasMas(matrix [][]rune, i int, j int) int {
	var result int
	if (string(matrix[i-1][j-1]) == "M" && string(matrix[i-1][j+1]) == "M" &&
		string(matrix[i+1][j-1]) == "S" && string(matrix[i+1][j+1]) == "S") ||
		(string(matrix[i-1][j-1]) == "S" && string(matrix[i-1][j+1]) == "M" &&
			string(matrix[i+1][j-1]) == "S" && string(matrix[i+1][j+1]) == "M") ||
		(string(matrix[i-1][j-1]) == "M" && string(matrix[i-1][j+1]) == "S" &&
			string(matrix[i+1][j-1]) == "M" && string(matrix[i+1][j+1]) == "S") ||
		(string(matrix[i-1][j-1]) == "S" && string(matrix[i-1][j+1]) == "S" &&
			string(matrix[i+1][j-1]) == "M" && string(matrix[i+1][j+1]) == "M") {
		result++
	}
	return result
}
