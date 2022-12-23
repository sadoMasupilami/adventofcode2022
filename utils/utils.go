package utils

import (
	"log"
	"os"
	"strconv"
)

func ReadFileIntoString(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ConvertStringToInt(string string) int {
	integer, err := strconv.Atoi(string)
	if err != nil {
		log.Fatal(err)
	}
	return integer
}

func SumUpArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
