package main

type Coordinates [3]int

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
	start := Coordinates{}
	for s := 0; s < len(data[0]); s++ {
		if data[0][s] == 'S' {
			start = Coordinates{0, s}
			break
		}
	}
	visited := map[Coordinates]int{}
	answer = helper(data, start, visited)
	return answer
}
func helper(data []string, coords Coordinates, visited map[Coordinates]int) int {
	row, col := coords[0], coords[1]
	if row == len(data) {
		return 1
	}
	if data[row][col] == '^' {
		value, exists := visited[Coordinates{row, col}]
		if exists {
			return value
		} else {
			paths := helper(data, Coordinates{row, col - 1}, visited) + helper(data, Coordinates{row, col + 1}, visited)
			visited[Coordinates{row, col}] = paths
			return paths
		}
	}

	return helper(data, Coordinates{row + 1, col}, visited)
}

// func Day7Part2(data []string) int {
// 	answer := 0
// 	start := Coordinates{}
// 	for s := 0; s < len(data[0]); s++ {
// 		if data[0][s] == 'S' {
// 			start = Coordinates{0, s}
// 			break
// 		}
// 	}
// 	stack := []Coordinates{start}
// 	for len(stack) > 0 {
// 		startLocation := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		row, col := startLocation[0], startLocation[1]
// 		for row < len(data) {
// 			symbol := data[row][col]
// 			if symbol == '^' {
// 				stack = append(stack, Coordinates{row, col + 1})
// 				col--
// 			}
// 			row++
// 			if row == len(data) {
// 				answer++
// 			}
// 		}
// 	}

// 	return answer
// }
