package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func convertToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		total          int
		decreasedCount int
		previousNumber int
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentNumber := convertToInt(scanner.Text())

		if currentNumber < previousNumber {
			decreasedCount++
		}

		total++
		previousNumber = currentNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println((total - 1) - decreasedCount)
}
