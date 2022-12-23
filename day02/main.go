package main

import (
	"adventofcode2022/utils"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputs := utils.ReadFileIntoString("day02/input.txt")
	points := 0
	for _, input := range strings.Split(inputs, "\n") {
		points += evaluate(input)
		if evaluate(input) == -1 {
			log.Fatal("ERROR, invalid input")
		}
	}
	fmt.Println("POINTS: ", points)

	points = 0
	for _, input := range strings.Split(inputs, "\n") {
		points += evaluateComplex(input)
		if evaluateComplex(input) == -1 {
			log.Fatal("ERROR, invalid input")
		}
	}
	fmt.Println("POINTS: ", points)
}

func evaluate(input string) int {
	switch input {
	default:
		return -1
	case "A X":
		return 1 + 3
	case "A Y":
		return 2 + 6
	case "A Z":
		return 3 + 0
	case "B X":
		return 1 + 0
	case "B Y":
		return 2 + 3
	case "B Z":
		return 3 + 6
	case "C X":
		return 1 + 6
	case "C Y":
		return 2 + 0
	case "C Z":
		return 3 + 3
	}
}

func evaluateComplex(input string) int {
	switch input {
	default:
		return -1
	case "A X":
		return 3 + 0
	case "B X":
		return 1 + 0
	case "C X":
		return 2 + 0
	case "A Y":
		return 1 + 3
	case "B Y":
		return 2 + 3
	case "C Y":
		return 3 + 3
	case "A Z":
		return 2 + 6
	case "B Z":
		return 3 + 6
	case "C Z":
		return 1 + 6
	}
}
