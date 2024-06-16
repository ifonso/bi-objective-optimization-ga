package utils_test

import (
	"testing"

	"github.com/ifonso/bi-objective-optimization-ga/src/types"
	"github.com/ifonso/bi-objective-optimization-ga/src/utils"
)

func TestCanRemoveContainerFromHarbor(t *testing.T) {
	harbor := []types.Container{
		{Position: types.Position{X: 1, Y: 1, Z: 1}},
		{Position: types.Position{X: 1, Y: 1, Z: 2}},
	}
	container := types.Container{Position: types.Position{X: 1, Y: 1, Z: 1}}

	if utils.CanRemoveContainerFromHarbor(harbor, container) {
		t.Error("Expected false, got true")
	}

	container = types.Container{Position: types.Position{X: 1, Y: 1, Z: 2}}

	if !utils.CanRemoveContainerFromHarbor(harbor, container) {
		t.Error("Expected true, got false")
	}
}

func TestCanAddContainerIntoShipPosition(t *testing.T) {
	ship := []types.Container{
		{Position: types.Position{X: 1, Y: 1, Z: 1}},
	}
	position := types.Position{X: 1, Y: 1, Z: 2}

	if !utils.CanAddContainerIntoShipPosition(ship, position) {
		t.Error("Expected true, got false")
	}

	position = types.Position{X: 1, Y: 1, Z: 1}

	if utils.CanAddContainerIntoShipPosition(ship, position) {
		t.Error("Expected false, got true")
	}
}

func TestIsValidSolutionGenome(t *testing.T) {
	genome := types.Genome{
		{InitialPosition: types.Position{X: 1, Y: 1, Z: 1}, FinalPosition: types.Position{X: 1, Y: 1, Z: 1}},
		{InitialPosition: types.Position{X: 2, Y: 2, Z: 2}, FinalPosition: types.Position{X: 2, Y: 2, Z: 2}},
	}

	if !utils.IsValidSolutionGenome(genome) {
		t.Error("Expected true, got false")
	}

	genome = types.Genome{
		{InitialPosition: types.Position{X: 1, Y: 1, Z: 1}, FinalPosition: types.Position{X: 1, Y: 1, Z: 1}},
		{InitialPosition: types.Position{X: 1, Y: 1, Z: 1}, FinalPosition: types.Position{X: 2, Y: 2, Z: 2}},
	}

	if utils.IsValidSolutionGenome(genome) {
		t.Error("Expected false, got true")
	}
}

func TestInvalidOrderingCount(t *testing.T) {
	genome := types.Genome{
		{InitialPosition: types.Position{X: 5, Y: 2, Z: 2}, FinalPosition: types.Position{X: 3, Y: 2, Z: 3}},
		{InitialPosition: types.Position{X: 2, Y: 3, Z: 2}, FinalPosition: types.Position{X: 3, Y: 3, Z: 1}},
		{InitialPosition: types.Position{X: 4, Y: 2, Z: 2}, FinalPosition: types.Position{X: 3, Y: 2, Z: 1}},
		{InitialPosition: types.Position{X: 3, Y: 1, Z: 2}, FinalPosition: types.Position{X: 2, Y: 2, Z: 1}},
		{InitialPosition: types.Position{X: 2, Y: 1, Z: 1}, FinalPosition: types.Position{X: 3, Y: 4, Z: 1}},
		{InitialPosition: types.Position{X: 4, Y: 2, Z: 1}, FinalPosition: types.Position{X: 2, Y: 3, Z: 1}},
	}

	if count := utils.InvalidOrderingCount(genome); count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}

	genome = types.Genome{
		{InitialPosition: types.Position{X: 1, Y: 1, Z: 2}, FinalPosition: types.Position{X: 1, Y: 1, Z: 1}},
		{InitialPosition: types.Position{X: 1, Y: 1, Z: 1}, FinalPosition: types.Position{X: 1, Y: 1, Z: 2}},
	}

	if count := utils.InvalidOrderingCount(genome); count != 0 {
		t.Errorf("Expected 0, got %d", count)
	}
}
