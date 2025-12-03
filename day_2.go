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

func Day2Part2(data []string) int {
	answer := 0
	idList := strings.Split(data[0], ",")
	for _, idRanges := range idList {
		idRange := strings.Split(idRanges, "-")
		idStart, _ := strconv.Atoi(idRange[0])
		idEnd, _ := strconv.Atoi(idRange[1])

		for id := idStart; id <= idEnd; id++ {
			strID := strconv.Itoa(id)
			strIDLength := len(strID)
			smallSequenceFound := false
			for i := 1; i <= strIDLength/2; i++ {
				if smallSequenceFound {
					break
				}

				firstSet := strID[:i]
				potentialMatch := strID[i : i+i]
				if firstSet == potentialMatch && strIDLength%i == 0 {
					// fmt.Printf("Full ID: %s, The first set is: %s, the potential match is: %s\n", strID, firstSet, potentialMatch)
					if i+i == strIDLength {
						answer += id
						break
					}
					for j := i + i; j+i <= strIDLength; j += i {
						//fmt.Printf("Full ID: %s, The first set is: %s, the potential match is: %s\n", strID, firstSet, strID[j:j+i])

						if firstSet != strID[j:j+i] {
							break
						}
						if j+i == strIDLength {
							//fmt.Printf("Adding %d to answer.\n", id)
							smallSequenceFound = true
							answer += id
						}
					}
				}
			}
		}
	}

	return answer
}
