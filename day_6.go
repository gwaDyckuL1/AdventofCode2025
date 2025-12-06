package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day6Part1(data []string) int {
	answer := 0

	numList := [][]int{}
	symbolList := []string{}
	for idx, line := range data {
		strToarray := strings.Split(line, " ")
		tempList := []int{}
		for _, num := range strToarray {
			if len(num) == 0 {
				continue
			}

			if idx == len(data)-1 {
				symbolList = append(symbolList, num)
				continue
			}

			number, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Problem converting number", err)
			}
			tempList = append(tempList, number)
		}
		if len(tempList) > 0 {
			numList = append(numList, tempList)
		}
	}

	for col := 0; col < len(numList[0]); col++ {
		tempMultiply := 1
		for row := 0; row < len(numList); row++ {
			symbol := symbolList[col]
			switch symbol {
			case "+":
				answer += numList[row][col]
			case "*":
				tempMultiply *= numList[row][col]
				if row == len(numList)-1 {
					answer += tempMultiply
				}
			}
		}
	}
	return answer
}
