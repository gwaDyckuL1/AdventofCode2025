package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day9Part1(data []string) int {
	answer := 0

	intData := [][2]int{}
	for _, set := range data {
		strNum := strings.Split(set, ",")
		num1, err := strconv.Atoi(strNum[0])
		if err != nil {
			fmt.Println("Error converting num1. ", err)
		}
		num2, err := strconv.Atoi(strNum[1])
		if err != nil {
			fmt.Println("Error converting num2. ", err)
		}
		intData = append(intData, [2]int{num1, num2})
	}

	largestArea := 0
	for i := 0; i < len(intData); i++ {
		for j := i + 1; j < len(intData); j++ {
			x1, y1 := intData[i][0], intData[i][1]
			x2, y2 := intData[j][0], intData[j][1]
			area := (abs(x1, x2) + 1) * (abs(y1, y2) + 1)
			if area > largestArea {
				largestArea = area
			}
		}
	}
	answer = largestArea
	return answer
}

func abs(x, y int) int {
	answer := x - y
	if answer < 0 {
		return y - x
	}
	return answer
}
