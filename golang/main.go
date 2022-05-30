package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func CreateRandomGrid(sideLength int) Grid {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	grid := make(Grid, sideLength)
	for y := 0; y < sideLength; y++ {
		grid[y] = make([]bool, sideLength)
		for x := 0; x < sideLength; x++ {
			grid[y][x] = randGen.Intn(2) == 1
		}
	}
	return grid
}

func ParsePrintableRow(row []bool) string {
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

func PrintGrid(grid Grid) {
	// Terminal clear sequence from https://stackoverflow.com/a/33509850
	fmt.Printf("\033[0;0H")
	for _, row := range grid {
		fmt.Printf(ParsePrintableRow(row) + "\n")
	}
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type args struct {
	sideLength int
	waitTime   float64
}

func ParseArgs() args {
	var sideLength int
	flag.IntVar(&sideLength, "s", 20, "The side length of the grid")
	flag.IntVar(&sideLength, "sideLength", 20, "The side length of the grid")

	var waitTime float64
	flag.Float64Var(&waitTime, "w", 0.4, "The time to wait between each step")
	flag.Float64Var(&waitTime, "waitTime", 0.4, "The time to wait between each step")

	flag.Parse()

	return args{
		sideLength,
		waitTime,
	}
}

func HandleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			ClearScreen()
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}()
}

func main() {
	HandleInterrupt()

	ClearScreen()

	args := ParseArgs()
	grid := CreateRandomGrid(args.sideLength)

	for ok := true; ok; ok = true {
		PrintGrid(grid)
		grid = GameOfLifeStep(grid)
		time.Sleep(time.Duration(args.waitTime*1000) * time.Millisecond)
	}
}
