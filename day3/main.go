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
	fmt.Println(int(part1Return))
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
