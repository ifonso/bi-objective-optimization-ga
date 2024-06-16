package utils

import (
	"github.com/ifonso/bi-objective-optimization-ga/src/types"
)

// CanRemoveContainerFromHarbor checks if a container can be removed from the given position in the harbor.
func CanRemoveContainerFromHarbor(harbor []types.Container, container types.Container) bool {
	// Check if the container is on the second floor
	if container.Position.Z == 2 {
		return true
	}

	// Check if there is any container on the second floor directly above the given container
	for _, c := range harbor {
		if c.Position.Z == 2 && c.Position.X == container.Position.X && c.Position.Y == container.Position.Y {
			return false
		}
	}

	// No container directly above
	return true
}

// CanAddContainerIntoShipPosition checks if a container can be added to the given position on the ship.
func CanAddContainerIntoShipPosition(ship []types.Container, position types.Position) bool {
	// Check if there is a container in the chosen position
	for _, c := range ship {
		if c.Position.X == position.X && c.Position.Y == position.Y && c.Position.Z == position.Z {
			return false
		}
	}

	// Check if the container is on the first floor
	if position.Z == 1 {
		return true
	}

	// Check if there is a container underneath the chosen position
	for _, c := range ship {
		if c.Position.X == position.X && c.Position.Y == position.Y && c.Position.Z == position.Z-1 {
			return true
		}
	}

	// If no container is found underneath the chosen position
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
	}

	for i := 0; i < len(genome)-1; i++ {
		if !CanAddContainerIntoShipPosition(shipOrdered[:i], shipOrdered[i].Position) {
			invalidAddition++
		}
	}

	return invalidRemoval + invalidAddition
}
