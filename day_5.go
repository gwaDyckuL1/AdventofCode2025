package main

import (
	"strconv"
	"strings"
)

func Day5Part1(data []string) int {
	answer := 0

	splitPoint := 0
	for char := 0; char < len(data); char++ {
		character := data[char]
		if len(character) == 0 {
			splitPoint = char
		}
	}

	idRangesStr := data[:splitPoint]
	ingredientsStr := data[splitPoint+1:]

	idRange := [][]string{}
	for _, char := range idRangesStr {
		each := strings.Split(char, "-")
		idRange = append(idRange, each)
	}

	for _, ing := range ingredientsStr {
		ingredient := strToint(ing)

		for _, ran := range idRange {
			begin := strToint(ran[0])
			end := strToint(ran[1])
			if ingredient >= begin && ingredient <= end {
				answer++
				break
			}
		}
	}

	return answer
}

func Day5Part2(data []string) int {
	answer := 0

	return answer
}

func strToint(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic("Error in changing string to number")
	}
	return num
}
