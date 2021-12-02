package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	items, err := readInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1Return := part1(&items)
	fmt.Println(part1Return)

	part2Return := part2(&items)
	fmt.Println(part2Return)
}

func convertToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func arraySum(input []int) int {
	total := 0
	for _, num := range input {
		total += num
	}

	return total
}

func readInput(path string) ([]int, error) {
	items := []int{}

	file, err := os.Open(path)
	if err != nil {
		return items, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		converted, err := convertToInt(scanner.Text())
		if err != nil {
			return items, err
		}
		items = append(items, converted)
	}

	if err := scanner.Err(); err != nil {
		return items, err
	}

	return items, nil
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

		if arraySum(right) > arraySum(left) {
			incrementCount++
		}
		left = right
	}

	return incrementCount
}
