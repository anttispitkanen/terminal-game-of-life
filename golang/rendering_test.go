package main

import (
	"reflect"
	"testing"
)

func TestGetDiffForRendering(t *testing.T) {
	var oldGrid = Grid{
		{false, false, false}, {true, true, true}, {false, false, false},
	}
	var newGrid = Grid{
		{false, true, false}, {false, true, false}, {false, true, false},
	}
	var expectedDiff = []Diff{
		{1, 0, true}, {0, 1, false}, {2, 1, false}, {1, 2, true},
	}

	var diff = GetDiffForRendering(oldGrid, newGrid)
	if !reflect.DeepEqual(diff, expectedDiff) {
		t.Errorf("Expected %v, got %v", expectedDiff, diff)
	}
}
