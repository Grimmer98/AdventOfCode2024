package main

import (
	"adventofcode/ejercicio4"
	"fmt"
)

/*
	func countOccurrences(x int, s []int) int {
		var count int
		for _, value := range s {
			if x == value {
				count++
			}
		}
		return count
	}
*/
func main() {
	/*
		file, err := os.Open("input.txt")
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}
		defer file.Close()
		var firstList []int
		var secondList []int

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, "   ")
			if len(parts) != 2 {
				fmt.Println(len(parts))
				fmt.Println("Incorrect line format: ", line)
				continue
			}

			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])

			if err1 != nil || err2 != nil {
				fmt.Println("Error convertin numbers in line: ", line)
				continue
			}
			firstList = append(firstList, num1)
			secondList = append(secondList, num2)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file: ", err)
		}

		var similarities []int

		for _, value := range firstList {
			ocurrences := countOccurrences(value, secondList)
			fmt.Println(ocurrences)
			similarities = append(similarities, ocurrences)
		}

		var similarityScore int
		for index, value := range similarities {
			similarityScore += value * firstList[index]
		}
		fmt.Println(similarityScore)

		First part - day 1
		sort.Slice(firstList, func(i, j int) bool {
			return firstList[i] < firstList[j]
		})
		sort.Slice(secondList, func(i, j int) bool {
			return secondList[i] < secondList[j]
		})

			var differences []int

			for index := range firstList {
				difference := float64(firstList[index]) - float64(secondList[index])
				differences = append(differences, int(math.Abs(difference)))
			}
			var result int

			for _, value := range differences {
				result += value
			}
			fmt.Println(result)
	*/

	//fmt.Println(ejercicio2.Exercise2Part1Solver())
	//fmt.Println(ejercicio2.Exercise2Part2Solver())
	//fmt.Println(ejercicio3.Exercise3Part1Solver())
	//fmt.Println(ejercicio3.Exercise3Part2Solver())
	//fmt.Println(ejercicio4.Exercise4Part1Solver())
	fmt.Println(ejercicio4.Exercise4Part2Solver())
}
