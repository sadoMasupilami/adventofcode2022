package main

import (
	"adventofcode2022/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := utils.ReadFileIntoString("day01/input.txt")
	separatedInput := strings.Split(input, "\n\n")
	maxCalories := []int{0, 0, 0}
	for _, elf := range separatedInput {
		foods := strings.Split(elf, "\n")
		calories := 0
		for _, food := range foods {
			calory := utils.ConvertStringToInt(food)
			calories += calory
			if calories > maxCalories[0] {
				maxCalories = append(maxCalories, calories)
				sort.Ints(maxCalories)
				maxCalories = maxCalories[1:]
			}
		}
	}
	fmt.Println("MAX CALORIES: ", maxCalories[2])
	summedUpCalories := utils.SumUpArray(maxCalories)
	fmt.Println("TOP 3 CALORIES: ", summedUpCalories)
}
