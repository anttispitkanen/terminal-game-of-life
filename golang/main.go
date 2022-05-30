package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
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
	oldGrid := CreateRandomGrid(args.sideLength)
	newGrid := oldGrid
	RenderInitialGrid(oldGrid)

	for ok := true; ok; ok = true {
		oldGrid = newGrid
		newGrid = GameOfLifeStep(newGrid)
		diff := GetDiffForRendering(oldGrid, newGrid)
		RenderDiff(diff)
		time.Sleep(time.Duration(args.waitTime*1000) * time.Millisecond)
	}
}
