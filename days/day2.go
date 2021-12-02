package main

import (
	"aoc2021/days/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Direction struct {
	Where  string
	Amount int
}

func parseLines(lines []string) []Direction {
	buf := make([]Direction, 0)

	for _, line := range lines {
		split := strings.Split(line, " ")
		num, _ := strconv.Atoi(split[1])

		buf = append(buf, Direction{
			Where:  split[0],
			Amount: num,
		})
	}

	return buf
}

func main() {
	lines, err := utils.ReadInput(2)

	if err != nil {
		log.Fatal(err)
	}

	directions := parseLines(lines)

	amount := part1(directions)
	fmt.Printf("Part1: %d\n", amount[0])
	fmt.Printf("Part2: %d\n", amount[1])
}

func part1(directions []Direction) []int {
	horizontal := 0
	depth := 0
	depthAim := 0
	aim := 0

	for _, direction := range directions {
		switch direction.Where {
		case "forward":
			horizontal += direction.Amount
			depthAim += aim * direction.Amount
		case "down":
			depth += direction.Amount
			aim += direction.Amount
		case "up":
			depth -= direction.Amount
			aim -= direction.Amount
		}
	}

	return []int{horizontal * depth, horizontal * depthAim}
}
