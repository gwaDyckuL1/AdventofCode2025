package main

import (
	"math"
	"strconv"
)

func Part1(data []string) int {
	pointer := 50
	zeroes := 0
	for _, each := range data {
		direction := each[0]
		clicks, _ := strconv.Atoi(each[1:])

		if direction == 'L' {
			pointer -= clicks
		} else {
			pointer += clicks
		}

		pointer = ((pointer % 100) + 100) % 100

		if pointer == 0 {
			zeroes += 1
		}
	}

	return zeroes
}

func Part2(data []string) int {
	zeroes := 0
	pointer := 50

	for _, each := range data {
		direction := each[0]
		clicks, _ := strconv.Atoi(each[1:])
		fullSpins := clicks / 100
		startingPointer := pointer

		zeroes += fullSpins

		remainingClicks := int(math.Abs(float64(clicks - (fullSpins * 100))))

		if direction == 'L' {
			pointer -= remainingClicks
		} else {
			pointer += remainingClicks
		}

		if startingPointer != 0 && (pointer < 0 || pointer > 100) {
			zeroes += 1
		}

		pointer = ((pointer % 100) + 100) % 100

		if pointer == 0 {
			zeroes += 1
		}
	}

	return zeroes
}
