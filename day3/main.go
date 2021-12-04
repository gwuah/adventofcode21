package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/gwuah/adventofcode21/common"
)

func main() {
	items, err := common.ReadInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1Return, err := part1(&items)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(int(part1Return), "part1")

	part2Return, err := part2(&items)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(int(part2Return))

}

func part1(report *[]string) (int, error) {
	binaryLength := len((*report)[0])

	countPerColum := map[int][]int{}
	for i := 0; i < binaryLength; i++ {
		countPerColum[i] = []int{0, 0}
	}

	for _, binary := range *report {
		for i := 0; i < binaryLength; i++ {
			if string(binary[i]) == "0" {
				countPerColum[i][0]++
			} else {
				countPerColum[i][1]++
			}
		}
	}

	binaryGamma, binaryEpsilon := bytes.Buffer{}, bytes.Buffer{}

	for i := 0; i < binaryLength; i++ {
		if countPerColum[i][0] > countPerColum[i][1] {
			binaryGamma.WriteString("0")
			binaryEpsilon.WriteString("1")
		} else {
			binaryGamma.WriteString("1")
			binaryEpsilon.WriteString("0")
		}
	}

	decimalGamma, err := strconv.ParseInt(binaryGamma.String(), 2, 64)
	if err != nil {
		return 0, err
	}

	decimalEpsilon, err := strconv.ParseInt(binaryEpsilon.String(), 2, 64)
	if err != nil {
		return 0, err
	}

	return int(decimalEpsilon) * int(decimalGamma), nil
}

func part2(report *[]string) (int, error) {

	resChan := make(chan int64, 2)

	go func() {
		o, err := oxygen(*report, 0)
		if err != nil {
			log.Fatal(err)
		}
		resChan <- o
	}()

	go func() {
		o, err := carbondioxide(*report, 0)
		if err != nil {
			log.Fatal(err)
		}
		resChan <- o
	}()

	answer := (1)
	answer *= int(<-resChan)
	answer *= int(<-resChan)

	close(resChan)

	return answer, nil

}

func oxygen(report []string, position int) (int64, error) {
	if len(report) == 1 {
		return strconv.ParseInt(report[0], 2, 64)
	}

	firstBitGroup := groupByBitInPosition(report, position)
	if len(firstBitGroup["0"]) > len(firstBitGroup["1"]) {
		return oxygen(firstBitGroup["0"], position+1)
	} else if len(firstBitGroup["1"]) > len(firstBitGroup["0"]) {
		return oxygen(firstBitGroup["1"], position+1)
	} else {
		return oxygen(firstBitGroup["1"], position+1)
	}
}

func carbondioxide(report []string, position int) (int64, error) {
	if len(report) == 1 {
		return strconv.ParseInt(report[0], 2, 64)
	}

	firstBitGroup := groupByBitInPosition(report, position)
	if len(firstBitGroup["0"]) < len(firstBitGroup["1"]) {
		return carbondioxide(firstBitGroup["0"], position+1)
	} else if len(firstBitGroup["1"]) < len(firstBitGroup["0"]) {
		return carbondioxide(firstBitGroup["1"], position+1)
	} else {
		return carbondioxide(firstBitGroup["0"], position+1)
	}
}

func groupByBitInPosition(report []string, position int) map[string][]string {
	bitInPosition := map[string][]string{}
	for _, binary := range report {
		char := string(binary[position])
		bitInPosition[char] = append(bitInPosition[char], binary)
	}
	return bitInPosition
}
