package main

import (
	"reflect"
	"testing"
)

var testGrid = Grid{
	{false, false, false}, {true, true, true}, {false, false, false},
}

func TestCountAliveNeighbors(t *testing.T) {
	var ans int = 0

	// First row
	ans = CountAliveNeighbors(0, 0, testGrid)
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}
	ans = CountAliveNeighbors(1, 0, testGrid)
	if ans != 3 {
		t.Errorf("Expected 3, got %d", ans)
	}
	ans = CountAliveNeighbors(2, 0, testGrid)
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}

	// Second row
	ans = CountAliveNeighbors(0, 1, testGrid)
	if ans != 1 {
		t.Errorf("Expected 1, got %d", ans)
	}
	ans = CountAliveNeighbors(1, 1, testGrid)
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}
	ans = CountAliveNeighbors(2, 1, testGrid)
	if ans != 1 {
		t.Errorf("Expected 1, got %d", ans)
	}

	// Third row
	ans = CountAliveNeighbors(0, 2, testGrid)
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}
	ans = CountAliveNeighbors(1, 2, testGrid)
	if ans != 3 {
		t.Errorf("Expected 3, got %d", ans)
	}
	ans = CountAliveNeighbors(2, 2, testGrid)
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}
}

func TestGameOfLifeStep(t *testing.T) {
	expectedGrid := Grid{
		{false, true, false}, {false, true, false}, {false, true, false},
	}

	// First step should result in the expected grid
	firstResultGrid := GameOfLifeStep(testGrid)
	if reflect.DeepEqual(firstResultGrid, expectedGrid) == false {
		t.Errorf("Expected %v, got %v", expectedGrid, firstResultGrid)
	}

	// Second step should result in the original grid again
	secondResultGrid := GameOfLifeStep(firstResultGrid)
	if reflect.DeepEqual(secondResultGrid, testGrid) == false {
		t.Errorf("Expected %v, got %v", testGrid, secondResultGrid)
	}
}
