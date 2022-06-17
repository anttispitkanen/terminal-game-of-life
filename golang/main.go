package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
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

func ParseArgs(maxSideLength int) args {
	var sideLength int
	flag.IntVar(&sideLength, "s", 20, "The side length of the grid")
	flag.IntVar(&sideLength, "sideLength", 20, "The side length of the grid, max "+strconv.Itoa(maxSideLength))

	var waitTime float64
	flag.Float64Var(&waitTime, "w", 0.4, "The time to wait between each step")
	flag.Float64Var(&waitTime, "waitTime", 0.4, "The time to wait between each step")

	flag.Parse()

	if sideLength > maxSideLength {
		log.Fatalf("Side length can be at most %d with your current terminal window size", maxSideLength)
	}

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

func GetMaxSideLength() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	terminalDimensions := strings.Split(strings.TrimSpace(string(out)), " ")
	heightString := terminalDimensions[0]
	widthString := terminalDimensions[1]

	width, err := strconv.Atoi(widthString)
	if err != nil {
		log.Fatal(err)
	}
	height, err := strconv.Atoi(heightString)
	if err != nil {
		log.Fatal(err)
	}

	// Terminal height as rows. Due to how the grid is printed, we can have max
	// N-1 rows.
	maxHeight := height - 1

	// Terminal width as columns, 1 column = 1 character. Since we render "🟪 ",
	// each game cell takes 3 characters. Thus max width is N/3.
	maxWidth := width / 3

	// Game must fit in terminal view, so the shorter length is the max size.
	if maxHeight < maxWidth {
		return maxHeight
	} else {
		return maxWidth
	}
}

func main() {
	HandleInterrupt()

	ClearScreen()

	args := ParseArgs(GetMaxSideLength())
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
