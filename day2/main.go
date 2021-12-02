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

	part2Return, err := part2(&items)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part2Return)

}

func part1(movements *[]string) (int, error) {
	movementMap := map[string]int{
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
		movementMap[direction] += int(units)
	}

	return (movementMap["forward"] * (movementMap["down"] - movementMap["up"])), nil
}

func part2(movements *[]string) (int, error) {
	movementMap := map[string]int{
		"forward": 0,
		"aim":     0,
		"depth":   0,
	}

	for _, movement := range *movements {
		splitted := strings.Split(movement, " ")

		direction := splitted[0]

		units, err := common.ConvertToInt(splitted[1])
		if err != nil {
			return 0, err
		}

		if direction == "down" {
			movementMap["aim"] += units
		}

		if direction == "up" {
			movementMap["aim"] -= units
		}

		if direction == "forward" {
			movementMap["forward"] += units
			movementMap["depth"] += movementMap["aim"] * units
		}
	}

	return (movementMap["forward"] * movementMap["depth"]), nil
}
