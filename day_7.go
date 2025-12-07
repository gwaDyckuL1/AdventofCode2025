package main

type Coordinates [2]int

func Day7Part1(data []string) int {
	answer := 0
	start := Coordinates{}
	for s := 0; s < len(data[0]); s++ {
		if data[0][s] == 'S' {
			start = Coordinates{0, s}
			break
		}
	}

	queue := []Coordinates{start}
	splitterVisited := map[Coordinates]bool{}
	splitterVisited[start] = true
	head := 0
	for head < len(queue) {
		currLocation := queue[head]
		row, col := currLocation[0], currLocation[1]
		head++
		currSymbol := data[row][col]
		for currSymbol != '^' {
			row++
			currSymbol = data[row][col]
			if row == len(data)-1 {
				break
			}
		}
		if currSymbol == '^' {
			currLocation = Coordinates{row, col}
			if !splitterVisited[currLocation] {
				answer++
				splitterVisited[currLocation] = true
				queue = append(queue, Coordinates{row, col - 1})
				queue = append(queue, Coordinates{row, col + 1})
			}
		}
	}
	return answer
}

func Day7Part2(data []string) int {
	answer := 0

	return answer
}
