package ejercicio2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Exercise2Part1Solver() int {
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
		safety += isSafe(report)
	}
	return safety
}

func isSafe(report []int) int {

	if len(report) < 2 {
		return 1 // Single or empty report is trivially safe
	}

	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if int(math.Abs(float64(diff))) < 1 || int(math.Abs(float64(diff))) > 3 || (diff > 0) != isIncreasing {
			return 0 // Violates step constraint or monotonicity
		}
	}

	return 1
}
