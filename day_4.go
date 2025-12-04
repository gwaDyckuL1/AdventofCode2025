package main

func Day4Part1(data []string) int {
	answer := 0
	directions := [][2]int{
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
	}
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[row]); col++ {
			if data[row][col] == '@' {
				count := 0
				for idx, direction := range directions {
					nextRow := row + direction[0]
					nextCol := col + direction[1]

					if nextRow >= 0 && nextCol >= 0 && nextRow < len(data) && nextCol < len(data[row]) {
						if data[nextRow][nextCol] == '@' {
							count++
						}
					}
					if count > 3 {
						break
					}
					if idx == 7 {
						answer++
					}
				}
			}
		}
	}

	return answer
}

func Day4Part2(data []string) int {
	answer := 0

	grid := make([][]byte, len(data))
	for i := range data {
		grid[i] = []byte(data[i])
	}
	moreRolls := true
	for moreRolls {
		count := Part1Modified(grid)
		answer += count
		if count == 0 {
			moreRolls = false
		}
	}
	return answer
}

func Part1Modified(data [][]byte) int {
	answer := 0
	directions := [][2]int{
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
	}
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[row]); col++ {
			if data[row][col] == '@' {
				count := 0
				for idx, direction := range directions {
					nextRow := row + direction[0]
					nextCol := col + direction[1]

					if nextRow >= 0 && nextCol >= 0 && nextRow < len(data) && nextCol < len(data[row]) {
						if data[nextRow][nextCol] == '@' {
							count++
						}
					}
					if count > 3 {
						break
					}
					if idx == 7 {
						data[row][col] = '.'
						answer++
					}
				}
			}
		}
	}

	return answer
}
