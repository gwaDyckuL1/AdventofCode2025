package main

import (
	"sort"
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
	splitPoint := 0
	for char := 0; char < len(data); char++ {
		character := data[char]
		if len(character) == 0 {
			splitPoint = char
		}
	}
	idRangesStr := data[:splitPoint]
	idRange := [][]int{}
	for _, char := range idRangesStr {
		each := strings.Split(char, "-")
		begin := strToint(each[0])
		end := strToint(each[1])
		idRange = append(idRange, []int{begin, end})
	}
	sort.Slice(idRange, func(i, j int) bool {
		return idRange[i][0] < idRange[j][0]
	})
	combinedRanges := [][]int{}
	start := idRange[0][0]
	for i := 0; i < len(idRange); i++ {
		end := idRange[i][1]
		for j := i + 1; j < len(idRange); j++ {
			if j == len(idRange)-1 {
				end = max(end, idRange[j][1])
				combinedRanges = append(combinedRanges, []int{start, end})
				i = j + 1
				break
			}
			if end > idRange[j][1] {
				continue
			}
			if end < idRange[j][0] {
				combinedRanges = append(combinedRanges, []int{start, end})
				start = idRange[j][0]
				i = j - 1
				break
			}
			if end >= idRange[j][0] && end <= idRange[j][1] {
				end = idRange[j][1]
			}
		}
	}

	for _, set := range combinedRanges {
		answer += set[1] - set[0] + 1
	}
	return answer
}

func strToint(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic("Error in changing string to number")
	}
	return num
}
