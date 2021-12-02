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
}

func convertToInt(s string) (int, error) {
	return strconv.Atoi(s)
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
