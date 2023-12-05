package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed day3-input.txt
var input string

var numbersRegex = regexp.MustCompile(`\d+`)
var symbolsRegex = regexp.MustCompile(`[^a-zA-Z0-9\s.]`)
var asterixRegex = regexp.MustCompile(`\*`)

// PART1
func findNumbersAdjacentToSymbols(lines []string) int {
	sum := 0
	for lineNumber, line := range lines {
		numbers := extractRegexMatchAndIndexes(line, *numbersRegex)
		for idx, number := range numbers {
			isPartNumber := false
			if lineNumber != 0 { // line before
				symbolsBefore := extractRegexMatchAndIndexes(lines[lineNumber-1], *symbolsRegex)
				isPartNumber = hasAdjacentSymbol(symbolsBefore, idx, number)
			}

			if !isPartNumber {
				//current line
				symbolsCurrent := extractRegexMatchAndIndexes(lines[lineNumber], *symbolsRegex)
				isPartNumber = hasAdjacentSymbol(symbolsCurrent, idx, number)

				if lineNumber+1 < len(lines) && !isPartNumber { // line after
					symbolsAfter := extractRegexMatchAndIndexes(lines[lineNumber+1], *symbolsRegex)
					isPartNumber = hasAdjacentSymbol(symbolsAfter, idx, number)
				}
			}

			if isPartNumber {
				converted, _ := strconv.Atoi(number)
				sum += converted
			}
		}

	}
	return sum
}

func hasAdjacentSymbol(symbolsBefore map[int]string, idx int, number string) bool {
	for symIdx := range symbolsBefore {
		if idx-1 <= symIdx && symIdx < idx+len(number)+1 {
			return true
		}
	}
	return false
}

func extractRegexMatchAndIndexes(input string, exp regexp.Regexp) (result map[int]string) {
	matches := exp.FindAllStringIndex(input, -1)

	result = make(map[int]string)
	for _, match := range matches {
		number := input[match[0]:match[1]]
		result[match[0]] = number
	}

	return result
}

// PART2

func findGearsAdjacentToTwoNumbers(lines []string) int {
	sum := 0

	for lineNumber, line := range lines {
		gears := extractRegexMatchAndIndexes(line, *asterixRegex)
		for gearIdx, _ := range gears {
			foundNumbers := make([]int, 0)
			if lineNumber != 0 { // line before
				numbersBefore := extractRegexMatchAndIndexes(lines[lineNumber-1], *numbersRegex)
				x, y := getAdjacentNumber(numbersBefore, gearIdx)
				if x > 0 {
					foundNumbers = append(foundNumbers, x)
				}
				if y > 0 {
					foundNumbers = append(foundNumbers, y)
				}
			}

			numbersCurrent := extractRegexMatchAndIndexes(lines[lineNumber], *numbersRegex)
			x, y := getAdjacentNumber(numbersCurrent, gearIdx)
			if x > 0 {
				foundNumbers = append(foundNumbers, x)
			}
			if y > 0 {
				foundNumbers = append(foundNumbers, y)
			}

			if len(foundNumbers) > 2 {
				continue
			}

			if lineNumber+1 < len(lines) { // line after
				numbersAfter := extractRegexMatchAndIndexes(lines[lineNumber+1], *numbersRegex)
				x, y := getAdjacentNumber(numbersAfter, gearIdx)
				if x > 0 {
					foundNumbers = append(foundNumbers, x)
				}
				if y > 0 {
					foundNumbers = append(foundNumbers, y)
				}

			}

			if len(foundNumbers) == 2 {
				fmt.Printf("found gear: %v \n", lineNumber)
				sum += foundNumbers[0] * foundNumbers[1]
			}
		}
	}

	return sum
}

func getAdjacentNumber(numbersMap map[int]string, idx int) (first, second int) {
	for numberIdx, number := range numbersMap {
		isAdjacent := numberIdx-1 <= idx && numberIdx+len(number)+1 > idx
		if isAdjacent {
			res, _ := strconv.Atoi(number)
			if first == 0 {
				first = res
			} else if second == 0 {
				second = res
			} else {
				return 0, 0
			}
		}
	}
	return
}

func main() {
	fmt.Println(findGearsAdjacentToTwoNumbers(strings.Split(input, "\n")))

}
