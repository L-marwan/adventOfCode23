package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day2-input.txt
var input string

var split = strings.Split
var trim = strings.TrimSpace

func toInt(s string) (res int) {
	res, _ = strconv.Atoi(s)
	return res
}

func parseGame(gameline string) (possible bool) {
	sets := trim(gameline[strings.Index(gameline, ":")+1:])
	possible = true

	for _, set := range split(sets, ";") {

		splitset := split(set, ",")
		for _, pull := range splitset {
			splitpull := split(trim(pull), " ")
			number, color := toInt(splitpull[0]), splitpull[1]
			if (color == "red" && number > 12) || (color == "green" && number > 13) || (color == "blue" && number > 14) {
				possible = false
				break
			}
		}
	}

	return
}

func parseGamePartTwo(gameline string) (blue, red, green int) {
	sets := trim(gameline[strings.Index(gameline, ":")+1:])

	for _, set := range split(sets, ";") {

		splitset := split(set, ",")
		for _, pull := range splitset {
			splitpull := split(trim(pull), " ")
			number, color := toInt(splitpull[0]), splitpull[1]

			switch color {
			case "red":
				if number > red {
					red = number
				}
			case "blue":
				if number > blue {
					blue = number
				}
			case "green":
				if number > green {
					green = number
				}
			}

		}
	}

	return
}

func main() {

	sum := 0
	for _, line := range split(input, "\n") {
		blue, red, green := parseGamePartTwo(line)
		sum += blue * red * green
	}

	fmt.Println(sum)

}
