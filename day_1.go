package main

import (
	"fmt"
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

		fmt.Println(pointer)

		if pointer == 0 {
			zeroes += 1
		}
	}

	return zeroes
}
