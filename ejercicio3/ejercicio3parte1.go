package ejercicio3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Exercise3Part1Solver() int {
	file, err := os.Open("./ejercicio3/test.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return -1
	}
	defer file.Close()
	var total int
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing numbers")
				continue
			}
			product := num1 * num2
			total += product
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
	return total
}
