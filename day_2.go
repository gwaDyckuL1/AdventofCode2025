package main

import (
	"strconv"
	"strings"
)

func Day2Part1(data []string) int {
	idList := strings.Split(data[0], ",")
	answer := 0
	for _, id := range idList {
		idRange := strings.Split(id, "-")
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])

		for i := start; i <= end; i++ {
			strNum := strconv.Itoa(i)
			numLength := len(strNum)
			if strNum[:numLength/2] == strNum[numLength/2:] {
				answer += i
			}
		}
	}

	return answer
}
