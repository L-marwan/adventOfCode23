package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day4-input.txt
var input string

func solvePart1(lines []string) int {

	result := 0
	for idx, line := range lines {
		cardScore := 0
		numbers := strings.Split(line[strings.Index(line, ":")+1:], "|")
		winning := strings.Split(strings.TrimSpace(numbers[0]), " ")
		guesses := strings.Split(strings.TrimSpace(numbers[1]), " ")

		common := findCommonElements(winning, guesses)
		fmt.Printf("line: %v: %v \n", idx, common)

		for i := 0; i < len(common); i++ {
			if cardScore == 0 {
				cardScore = 1
			} else {
				cardScore *= 2
			}
		}

		fmt.Println(len(common))
		result += cardScore
	}

	return result
}

func findCommonElements(slice1, slice2 []string) []string {
	set := make(map[string]bool)
	var common []string

	// Add elements from the first slice to the set
	for _, item := range slice1 {
		set[item] = true
	}

	// Check for common elements in the second slice
	for _, item := range slice2 {
		if set[item] {
			common = append(common, item)
			// Remove the element to avoid duplicates
			delete(set, item)
		}
	}

	return removeBlankStrings(common)
}

func removeBlankStrings(slice []string) []string {
	result := make([]string, 0)

	for _, str := range slice {
		if strings.TrimSpace(str) != "" {
			result = append(result, str)
		}
	}

	return result
}

func solvePart2(lines []string) int64 {

	var cardsNumber int64
	cardsMap := make(map[int]int64)

	//
	for idx := 0; idx < len(lines); idx++ {
		cardsMap[idx] = cardsMap[idx] + 1
		ammount := cardsMap[idx]
		fmt.Printf("card %v  copies: %v\n", idx, ammount)

		solvePart2WithStart(lines, idx, cardsMap, ammount)

	}

	fmt.Println(cardsMap)
	for _, ammount := range cardsMap {
		cardsNumber += ammount
	}
	return cardsNumber

}

func solvePart2WithStart(lines []string, start int, cardsMap map[int]int64, ammount int64) {

	cards := getScore(lines[start])
	fmt.Printf("card %v  score: %v\n", start, cards)
	for j := 1; j <= cards; j++ {
		if start+j < len(lines) {
			cardsMap[start+j] = cardsMap[start+j] + ammount
		}
	}
}

func getScore(line string) int {
	numbers := strings.Split(line[strings.Index(line, ":")+1:], "|")
	winning := strings.Split(strings.TrimSpace(numbers[0]), " ")
	guesses := strings.Split(strings.TrimSpace(numbers[1]), " ")
	return len(findCommonElements(winning, guesses))
}

func main() {
	fmt.Println(solvePart2(strings.Split(input, "\n")))

}
