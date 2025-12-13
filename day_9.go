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

func Day9Part2(data []string) int {
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
	length := len(intData)
	maxArea := 0
	for idx := range intData {
		for jdx := idx + 1; jdx < len(intData); jdx++ {
			cSet := intData[idx]
			nSet := intData[jdx]
			xMax := max(cSet[0], nSet[0])
			xMin := min(cSet[0], nSet[0])
			yMax := max(cSet[1], nSet[1])
			yMin := min(cSet[1], nSet[1])
			xMid := float64(xMax+xMin) / 2.0
			yMid := float64(yMax+yMin) / 2.0
			if xMax == xMin || yMax == yMin {
				continue
			}

			area := (xMax - xMin + 1) * (yMax - yMin + 1)
			if area > maxArea {
				count := 0
				good := true
				for j, set := range intData {
					x := set[0]
					y := set[1]
					nextSet := intData[(j+1)%length]
					xNext := nextSet[0]
					yNext := nextSet[1]
					yLow := min(y, yNext)
					yHigh := max(y, yNext)
					xLow := min(x, xNext)
					xHigh := max(x, xNext)
					if x == xNext && x > xMin && x < xMax {
						if yLow < yMax && yHigh > yMin {
							good = false
							break
						}
					}
					if y == yNext && y > yMin && y < yMax {
						if xLow < xMax && xHigh > xMin {
							good = false
							break
						}
					}
					if x == xNext && yNext > y {
						if float64(x) > xMid && yMid >= float64(yLow) && yMid < float64(yHigh) {
							count++
						}
					}
				}
				if good && count%2 == 1 {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func abs(x, y int) int {
	answer := x - y
	if answer < 0 {
		return y - x
	}
	return answer
}
