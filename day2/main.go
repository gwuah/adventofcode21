package main

import (
	"fmt"
	"log"
	"strings"

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

func part1(movements *[]string) (float64, error) {
	movementMap := map[string]float64{
		"forward": 0,
		"up":      0,
		"down":    0,
	}

	for _, movement := range *movements {
		splitted := strings.Split(movement, " ")

		direction := splitted[0]

		units, err := common.ConvertToInt(splitted[1])
		if err != nil {
			return 0, err
		}
		movementMap[direction] += float64(units)
	}

	depth := movementMap["down"] - movementMap["up"]

	return (movementMap["forward"] * depth), nil
}
