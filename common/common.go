package common

import (
	"bufio"
	"os"
	"strconv"
)

func ConvertToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ReadInput(path string) ([]string, error) {
	items := []string{}

	file, err := os.Open(path)
	if err != nil {
		return items, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func ArraySum(input []int) int {
	total := 0
	for _, num := range input {
		total += num
	}

	return total
}
