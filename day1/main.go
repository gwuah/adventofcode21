package main

import (
	"fmt"
	"log"

	"github.com/gwuah/adventofcode21/common"
)

func main() {
	stringItems, err := common.ReadInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	numberItems := []int{}
	for _, item := range stringItems {
		converted, err := common.ConvertToInt(item)
		if err != nil {
			log.Fatal(err)
		}

		numberItems = append(numberItems, converted)
	}

	part1Return := part1(&numberItems)
	fmt.Println(part1Return)

	part2Return := part2(&numberItems)
	fmt.Println(part2Return)
}

func part1(input *[]int) int {
	var (
		total          int
		decreasedCount int
		previousNumber int
	)

	for _, num := range *input {
		if num < previousNumber {
			decreasedCount++
		}

		total++
		previousNumber = num
	}

	return total - 1 - decreasedCount
}

func part2(input *[]int) int {
	total, incrementCount := len(*input), 0

	for i := -1; i <= total; i++ {
		// an optimization can be made here
		// for every iteration aside the first one, currrent left = right of previous iteration
		// for heavy payloads, that can spare us some extra computation.
		left := (*input)[i+1 : i+4]
		right := (*input)[i+2 : i+5]

		if common.ArraySum(right) > common.ArraySum(left) {
			incrementCount++
		}
		left = right
	}

	return incrementCount
}
