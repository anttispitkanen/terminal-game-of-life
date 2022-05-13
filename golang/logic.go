package main

// Return the larger of two integers
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Check neighbors of a cell, and return the number of alive neighbors
func CheckNeighbors(x, y int, grid [][]int) int {
	var aliveNeighbors int = 0
	var coordX = x
	var coordY = y
	var gridSizeY = len(grid)
	var gridSizeX = len(grid[0])

	var yStart = Max(coordY-1, 0)
	var yEnd = Min(coordY+1, gridSizeY-1)
	var xStart = Max(coordX-1, 0)
	var xEnd = Min(coordX+1, gridSizeX-1)

	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			// Don't include self in the count
			if (grid[y][x] == 1) && (x != coordX || y != coordY) {
				aliveNeighbors++
			}
		}
	}

	return aliveNeighbors
}

func GameOfLifeStep(grid [][]int) [][]int {
	/*
		Source: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life

		1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
		2. Any live cell with two or three live neighbours lives on to the next generation.
		3. Any live cell with more than three live neighbours dies, as if by overpopulation.
		4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

		These rules, which compare the behavior of the automaton to real life, can be condensed into the following:

		1. Any live cell with two or three live neighbours survives.
		2. Any dead cell with three live neighbours becomes a live cell.
		3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
	*/
	newGrid := make([][]int, len(grid))

	for y, row := range grid {
		newRow := make([]int, len(row))
		newGrid[y] = newRow

		for x, val := range row {
			originalAlive := val == 1
			aliveNeighborsCount := CheckNeighbors(x, y, grid)

			if originalAlive && (aliveNeighborsCount == 2 || aliveNeighborsCount == 3) {
				// Remain alive
				newGrid[y][x] = 1
			} else if !originalAlive && aliveNeighborsCount == 3 {
				// Be born
				newGrid[y][x] = 1
			} else {
				// Die or stay dead
				newGrid[y][x] = 0
			}
		}
	}

	return newGrid
}
