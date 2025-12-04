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

	for bank := 0; bank < len(data); bank++ {
		var jolts [12]byte
		counter := 0
		for battery := len(data[bank]) - 1; battery >= 0; battery-- {
			if counter < 12 {
				jolts[11-counter] = data[bank][battery]
				counter++
				continue
			}
			currBattery := data[bank][battery]
			for idx, currJolt := range jolts {
				if currBattery < currJolt {
					break
				}
				jolts[idx] = currBattery
				currBattery = currJolt
			}
		}
		number, err := strconv.Atoi(string(jolts[:]))
		if err != nil {
			fmt.Println("Problem turning bytes to number")
		}
		answer += number
	}

	return answer
}
