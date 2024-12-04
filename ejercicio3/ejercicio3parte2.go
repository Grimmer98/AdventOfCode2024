package ejercicio3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Exercise3Part2Solver() int {
	file, err := os.Open("./ejercicio3/input3.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return -1
	}
	defer file.Close()
	var total int
	scanner := bufio.NewScanner(file)
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()

		reCombined := regexp.MustCompile(`(do\(\)|don't\(\))|mul\((\d{1,3}),(\d{1,3})\)`)
		matches := reCombined.FindAllStringSubmatchIndex(line, -1)

		for _, match := range matches {
			// Extract match and type
			fullMatch := line[match[0]:match[1]]
			fmt.Println(fullMatch)
			if fullMatch == "do()" {
				enabled = true
			} else if fullMatch == "don't()" {
				enabled = false
			} else if match[2] == -1 { //mul(x,y)
				if enabled {
					num1, err1 := strconv.Atoi(line[match[4]:match[5]])
					num2, err2 := strconv.Atoi(line[match[6]:match[7]])

					if err1 != nil || err2 != nil {
						fmt.Println("Error parsing numbers")
						continue
					}

					total += num1 * num2
				}
			}
			fmt.Println(total)
			fmt.Printf("\n")
		}
	}
	return total
}
