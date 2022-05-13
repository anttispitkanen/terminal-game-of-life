package main

import (
	"reflect"
	"testing"
)

var testGrid = [][]int{
	{0, 0, 0}, {1, 1, 1}, {0, 0, 0},
}

// TestCheckNeighbors tests the CheckNeighbors function
func TestCheckNeighbors(t *testing.T) {
	var ans int = 0

	// First row
	ans = CheckNeighbors(0, 0, testGrid)
	if ans != 2 {
		t.Errorf("Expected 3, got %d", ans)
	}
	ans = CheckNeighbors(1, 0, testGrid)
	if ans != 3 {
		t.Errorf("Expected 3, got %d", ans)
	}
	ans = CheckNeighbors(2, 0, testGrid)
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}

	// Second row
	ans = CheckNeighbors(0, 1, testGrid)
	if ans != 1 {
		t.Errorf("Expected 3, got %d", ans)
	}
	ans = CheckNeighbors(1, 1, testGrid)
	if ans != 2 {
		t.Errorf("Expected 3, got %d", ans)
	}
	ans = CheckNeighbors(2, 1, testGrid)
	if ans != 1 {
		t.Errorf("Expected 2, got %d", ans)
	}

	// Third row
	ans = CheckNeighbors(0, 2, testGrid)
	if ans != 2 {
		t.Errorf("Expected 3, got %d", ans)
	}
	ans = CheckNeighbors(1, 2, testGrid)
	if ans != 3 {
		t.Errorf("Expected 3, got %d", ans)
	}
	ans = CheckNeighbors(2, 2, testGrid)
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}
}

func TestGameOfLifeStep(t *testing.T) {
	expectedGrid := [][]int{
		{0, 1, 0}, {0, 1, 0}, {0, 1, 0},
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
