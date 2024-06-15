package utils_test

import (
	"testing"

	"github.com/ifonso/bi-objective-optimization-ga/src/types"
	"github.com/ifonso/bi-objective-optimization-ga/src/utils"
)

func TestRemoveDuplicate(t *testing.T) {
	ints := []int{1, 2, 2, 3, 4, 4, 5}
	expectedInts := []int{1, 2, 3, 4, 5}
	resultInts := utils.RemoveDuplicate(ints)

	if len(resultInts) != len(expectedInts) {
		t.Errorf("Expected length %d, got %d", len(expectedInts), len(resultInts))
	}

	for i, v := range expectedInts {
		if resultInts[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, resultInts[i])
		}
	}

	positions := []types.Position{
		{X: 1, Y: 2, Z: 3},
		{X: 1, Y: 2, Z: 3},
		{X: 4, Y: 5, Z: 6},
	}

	expectedPositions := []types.Position{
		{X: 1, Y: 2, Z: 3},
		{X: 4, Y: 5, Z: 6},
	}

	resultPositions := utils.RemoveDuplicate(positions)

	if len(resultPositions) != len(expectedPositions) {
		t.Errorf("Expected length %d, got %d", len(expectedPositions), len(resultPositions))
	}

	for i, v := range expectedPositions {
		if !resultPositions[i].Equals(v) {
			t.Errorf("Expected %v at index %d, got %v", v, i, resultPositions[i])
		}
	}
}

func TestContains(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	if !utils.Contains(ints, 3) {
		t.Errorf("Expected to find 3 in %v", ints)
	}

	if utils.Contains(ints, 6) {
		t.Errorf("Did not expect to find 6 in %v", ints)
	}

	positions := []types.Position{
		{X: 1, Y: 2, Z: 3},
		{X: 4, Y: 5, Z: 6},
	}

	if !utils.Contains(positions, types.Position{X: 1, Y: 2, Z: 3}) {
		t.Errorf("Expected to find %v in %v", types.Position{X: 1, Y: 2, Z: 3}, positions)
	}

	if utils.Contains(positions, types.Position{X: 7, Y: 8, Z: 9}) {
		t.Errorf("Did not expect to find %v in %v", types.Position{X: 7, Y: 8, Z: 9}, positions)
	}
}

func TestContainsGenome(t *testing.T) {
	genome1 := types.Genome{
		{Order: 1, InitialPosition: types.Position{X: 1, Y: 2, Z: 3}, FinalPosition: types.Position{X: 4, Y: 5, Z: 6}},
	}

	genome2 := types.Genome{
		{Order: 1, InitialPosition: types.Position{X: 7, Y: 8, Z: 9}, FinalPosition: types.Position{X: 10, Y: 11, Z: 12}},
	}

	population := types.Population{genome1, genome2}

	if !utils.ContainsGenome(population, genome1) {
		t.Errorf("Expected to find genome %v in population %v", genome1, population)
	}

	if utils.ContainsGenome(population, types.Genome{
		{Order: 1, InitialPosition: types.Position{X: 1, Y: 1, Z: 1}, FinalPosition: types.Position{X: 2, Y: 2, Z: 2}},
	}) {
		t.Errorf("Did not expect to find genome %v in population %v", types.Genome{
			{Order: 1, InitialPosition: types.Position{X: 1, Y: 1, Z: 1}, FinalPosition: types.Position{X: 2, Y: 2, Z: 2}},
		}, population)
	}
}
