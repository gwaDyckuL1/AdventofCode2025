package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day6Part1(data []string) int {
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

func Day6Part2(data []string) int {
	answer := 0
	symbols := []byte{}
	numList := [][]int{}

	tempNumList := []int{}
	for col := 0; col < len(data[0]); col++ {
		var tempStrNum strings.Builder
		for row := 0; row < len(data); row++ {
			if row == len(data)-1 && data[row][col] != ' ' {
				symbols = append(symbols, data[row][col])
				break
			}
			tempStrNum.WriteString(string(data[row][col]))
		}
		cleanNum := strings.Fields(tempStrNum.String())
		if len(cleanNum) == 0 {
			numList = append(numList, tempNumList)
			tempNumList = []int{}
		} else if col == len(data[0])-1 {
			num, err := strconv.Atoi(cleanNum[0])
			if err != nil {
				log.Fatal("Error converting number:", err)
			}
			tempNumList = append(tempNumList, num)
			numList = append(numList, tempNumList)
		} else {
			num, err := strconv.Atoi(cleanNum[0])
			if err != nil {
				log.Fatal("Error converting number:", err)
			}
			tempNumList = append(tempNumList, num)
		}
	}
	symbolIndex := 0
	for _, row := range numList {
		tempMultNumber := 1
		for idx, num := range row {
			symbol := symbols[symbolIndex]
			switch symbol {
			case '+':
				answer += num
			case '*':
				tempMultNumber *= num
				if idx == len(row)-1 {
					answer += tempMultNumber
				}
			}
		}
		symbolIndex++
	}
	return answer
}
