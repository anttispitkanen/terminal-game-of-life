package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Generic utility
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//
// Initial rendering
//
func GetRowForRendering(row []bool) string {
	var parsedRow string = ""
	for _, val := range row {
		if val == true {
			parsedRow += fmt.Sprint("🟪 ")
		} else {
			parsedRow += fmt.Sprint("⬜️ ")
		}
	}
	return parsedRow
}

func RenderInitialGrid(grid Grid) {
	// Terminal clear sequence from https://stackoverflow.com/a/33509850
	fmt.Printf("\033[0;0H")
	for _, row := range grid {
		fmt.Printf(GetRowForRendering(row) + "\n")
	}
}

//
// Iterative rendering, optimized to only render the diff.
//
type Diff struct {
	x, y    int
	isAlive bool
}

func GetDiffForRendering(oldGrid, newGrid Grid) []Diff {
	var diffs []Diff
	for y := 0; y < len(oldGrid); y++ {
		for x := 0; x < len(oldGrid); x++ {
			if oldGrid[y][x] != newGrid[y][x] {
				diffs = append(diffs, Diff{x, y, newGrid[y][x]})
			}
		}
	}
	return diffs
}

func RenderDiff(diff []Diff) {
	// The x axis needs to be offset by a multiple of 3, as each rendered square,
	// including the following space, is 3 characters wide. Also the terminal indices
	// are one-based, so we need to offset by 1 from the grid indices that are 0-based.
	for _, d := range diff {
		if d.isAlive {
			fmt.Printf("\033[%d;%dH🟪 ", d.y+1, (d.x*3)+1)
		} else {
			fmt.Printf("\033[%d;%dH⬜️ ", d.y+1, (d.x*3)+1)
		}
	}
}
