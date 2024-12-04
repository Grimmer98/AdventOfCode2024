package ejercicio2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Exercise2Part2Solver() int {
	file, err := os.Open("./ejercicio2/input2.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return -1
	}
	defer file.Close()
	var matrix [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var report []int
		for _, value := range parts {
			level, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error converting numbers: ", err)
			}
			report = append(report, level)
		}
		matrix = append(matrix, report)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
	var safety int
	for _, report := range matrix {
		safety += isSafeRelaxed(report)
	}
	return safety
}

func isSafeRelaxed(report []int) int {
	if isSafe(report) == 1 {
		return 1
	}
	fmt.Println(report)
	// Try removing each level
	for i := 0; i < len(report); i++ {
		// Make a copy of the slice without modifying the original
		withoutLevel := make([]int, 0, len(report)-1)
		withoutLevel = append(withoutLevel, report[:i]...)
		withoutLevel = append(withoutLevel, report[i+1:]...)
		fmt.Println(withoutLevel)
		// Check if the modified slice is safe
		if isSafe(withoutLevel) == 1 {
			return 1 // Safe after removing one level
		}
	}

	return 0 // Unsafe even with the Problem Damp
}
