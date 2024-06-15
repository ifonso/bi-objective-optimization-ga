package utils

import (
	"github.com/ifonso/bi-objective-optimization-ga/src/types"
)

// Check if a container can be removed from the harbor by checking if is there a container on top of it.
func CanRemoveContainerFromHarbor(harbor []types.Container, contaienr types.Container) bool {
	if contaienr.Position.Z == 2 {
		return true
	}

	for _, c := range harbor {
		if c.Position.Z == 2 && c.Position.X == contaienr.Position.X && c.Position.Y == contaienr.Position.Y {
			return false
		}
	}

	return true
}

// Check if a container can be added to the ship by checking if is there a container on its bottle to be stacked.
func CanAddContainerIntoShipPosition(ship []types.Container, position types.Position) bool {
	// Check if there is a container in choosen position
	for _, c := range ship {
		if c.Position.X == position.X && c.Position.Y == position.Y && c.Position.Z == position.Z {
			return false
		}
	}
	// Check if the container is in first floor
	if position.Z == 1 {
		return true
	}
	// Check if there is a container underneath the choosen position
	for _, c := range ship {
		if c.Position.X == position.X && c.Position.Y == position.Y && c.Position.Z == position.Z-1 {
			return true
		}
	}

	return false
}

// Check if the genome is valid by checking if there are no repeated positions.
func IsValidSolutionGenome(genome types.Genome) bool {
	var initialPositions []types.Position
	var finalPositionos []types.Position

	for _, gene := range genome {
		initialPositions = append(initialPositions, gene.InitialPosition)
		finalPositionos = append(finalPositionos, gene.FinalPosition)
	}

	initialPositions = RemoveDuplicate(initialPositions)
	finalPositionos = RemoveDuplicate(finalPositionos)

	return len(initialPositions) == len(genome) && len(finalPositionos) == len(genome)
}

// Count the number of invalid orderings in the genome.
func InvalidOrderingCount(genome types.Genome) int {
	harborOrdered := make([]types.Container, len(genome))
	shipOrdered := make([]types.Container, len(genome))

	for i, d := range genome {
		harborOrdered[i] = types.Container{Position: d.InitialPosition}
		shipOrdered[i] = types.Container{Position: d.FinalPosition}
	}

	invalidRemoval, invalidAddition := 0, 0

	for i := 0; i < len(genome)-1; i++ {
		if !CanRemoveContainerFromHarbor(harborOrdered[i+1:], harborOrdered[i]) {
			invalidRemoval++
		}
		if !CanAddContainerIntoShipPosition(shipOrdered[:i+1], shipOrdered[i+1].Position) {
			invalidAddition++
		}
	}

	return invalidRemoval + invalidAddition
}
