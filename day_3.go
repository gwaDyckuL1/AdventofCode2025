package main

import (
	"fmt"
	"strconv"
)

func Day3Part1(data []string) int {
	answer := 0

	for bank := 0; bank < len(data); bank++ {
		firstNumber, secondNumber := 0, 0
		for battery := 0; battery < len(data[bank]); battery++ {
			number, err := strconv.Atoi(string(data[bank][battery]))
			if err != nil {
				fmt.Println("Character not a num ", err)
				break
			}

			if number > firstNumber && battery != len(data[bank])-1 {
				firstNumber = number
				secondNumber = 0
			} else if number > secondNumber {
				secondNumber = number
			}
		}
		numbersCombined := firstNumber*10 + secondNumber
		fmt.Printf("First: %d, Second: %d, Combined: %d\n", firstNumber, secondNumber, numbersCombined)
		answer += numbersCombined
	}
	return answer
}

func Day3Part2(data []string) int {
	answer := 0

	return answer
}
