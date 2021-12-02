package main

import (
	"aoc2021/days/utils"
	"fmt"
	"log"
)

func main() {
	data, err := utils.ReadInput(1)

	if err != nil {
		log.Fatal(data)
	}

	numbers, err := utils.InputToIntLines(data)

	if err != nil {
		log.Fatal(err)
	}

	part1Increase := Part1(numbers)
	part2Increases := Part2(numbers)

	fmt.Printf("Part 1: %d\n", part1Increase)
	fmt.Printf("Part 2: %d\n", part2Increases)
}

func Part1(numbers []int) int {
	timesRisen := 0
	lastNum := numbers[0]

	for _, number := range numbers[1:] {
		if number > lastNum {
			timesRisen++
		}
		lastNum = number
	}

	return timesRisen
}

func windowSum(window []int) int {
	return window[0] + window[1] + window[2]
}

func Part2(numbers []int) int {
	numIncreases := 0
	prevSum := windowSum(numbers[0:3])
	for i := 1; i <= len(numbers)-3; i++ {
		window := windowSum(numbers[i : i+3])
		if window > prevSum {
			numIncreases++
		}
		prevSum = window
	}

	return numIncreases
}
