package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var numbers = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func getFirstAndLastDigits(line string) (firstDigit, lastDigit int) {

	firstFound := false
	lastfound := false
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			if !firstFound {
				firstDigit = int(ch - '0')
				firstFound = true
			}
			lastDigit = int(ch - '0')
			lastfound = true
		}
	}
	if !lastfound {
		lastDigit = firstDigit
	}

	return
}

func replaceLetterDigitsWithDigits(line string) string {

	result := line
	for old, new := range numbers {
		result = s.Replace(result, old, new, -1)
	}
	return result
}

func main() {

	fmt.Println()

	fmt.Println("input is")
	file, err := os.Open("./day1-input.txt")
	check(err)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		x, y := getFirstAndLastDigits(replaceLetterDigitsWithDigits(fileScanner.Text()))
		sum += x*10 + y
	}

	fmt.Println(sum)

}
