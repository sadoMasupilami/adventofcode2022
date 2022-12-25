package main

import (
	"adventofcode2022/utils"
	"errors"
	"fmt"
	"log"
	"strings"
)

const letters = "-abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	inputs := utils.ReadFileIntoString("day03/input.txt")
	inputLines := strings.Split(inputs, "\n")
	points := 0
	for _, input := range inputLines {
		firstHalf, secondHalf := input[0:len(input)/2], input[len(input)/2:]
		//fmt.Println(firstHalf, "|", secondHalf)
		err, letter := findLetterInBothStrings(firstHalf, secondHalf)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("LETTER IN BOTH: ", letter)
		points += strings.Index(letters, letter)
	}
	fmt.Println("POINTS: ", points)

	points = 0
	for i := 0; i < len(inputLines); i += 3 {
		err, letter := findLetterInMultipleStrings(inputLines[i], inputLines[i+1], inputLines[i+2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----------")
		fmt.Println(inputLines[i], inputLines[i+1], inputLines[i+2])
		fmt.Println(letter)
		points += strings.Index(letters, letter)
	}
	fmt.Println("POINTS: ", points)
}

func findLetterInBothStrings(s, t string) (error, string) {
	for i := 0; i < len(s); i++ {
		if strings.Contains(t, string(s[i])) {
			return nil, string(s[i])
		}
	}
	return errors.New("no letter found in both strings: " + s + " and " + t), ""
}

func findLetterInMultipleStrings(strs ...string) (error, string) {
	if len(strs) < 2 {
		return errors.New("at least 2 strings are needed"), ""
	}

	foundLetters := strs[0]

	for j := 1; j < len(strs); j++ {
		tempFoundLetters := ""
		for i := 0; i < len(strs[j]); i++ {
			if strings.Contains(foundLetters, string(strs[j][i])) {
				if !strings.Contains(tempFoundLetters, string(strs[j][i])) {
					tempFoundLetters = tempFoundLetters + string(strs[j][i])
				}
			}
		}
		foundLetters = tempFoundLetters
	}
	return nil, foundLetters
}
