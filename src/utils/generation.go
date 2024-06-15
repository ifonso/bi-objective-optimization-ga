package utils

import (
	"math/rand/v2"

	"github.com/ifonso/bi-objective-optimization-ga/src/types"
)

func RandRange(min, max int) int {
	return rand.IntN(max+1-min) + min
}

// Generate a harbor with containers for fixed dimension (5x3x2).
func GenerateHarbor() []types.Container {
	harbor := make([]types.Container, 0, 5*3*2)

	for z := 1; z < 3; z++ {
		for x := 1; x < 6; x++ {
			for y := 1; y < 4; y++ {
				harbor = append(harbor, types.Container{
					Position: types.Position{X: x, Y: y, Z: z},
				})
			}
		}
	}

	return harbor
}

// Generate a random available position in the ship.
func GenerateRandomAvailableShipPosition(shipContainers []types.Container) types.Position {
	for {
		x := RandRange(1, 3)
		y := RandRange(1, 4)
		z := RandRange(1, 3)

		position := types.Position{X: x, Y: y, Z: z}

		if CanAddContainerIntoShipPosition(shipContainers, position) {
			return position
		}
	}
}

func removeRandomContainer(harbor *[]types.Container) types.Container {
	for {
		containerIndex := RandRange(0, len(*harbor)-1)
		container := (*harbor)[containerIndex]

		if CanRemoveContainerFromHarbor(*harbor, container) {
			*harbor = append((*harbor)[:containerIndex], (*harbor)[containerIndex+1:]...)
			return container
		}
	}
}

// Generate a random genome for the containers displacement.
func GenerateRandomGenome(harbor []types.Container) types.Genome {
	currentHarbor := make([]types.Container, len(harbor))
	copy(currentHarbor, harbor)

	var genome types.Genome
	var shipContainers []types.Container

	for i := 0; i < len(currentHarbor); i++ {
		// remove a container from harbor
		harborRemovedContainer := removeRandomContainer(&currentHarbor)
		// get an available position in the ship
		shipPosition := GenerateRandomAvailableShipPosition(shipContainers)
		// add the removed container to the ship
		shipContainers = append(shipContainers, types.Container{Position: shipPosition})

		genome = append(genome, types.Displacement{
			Order:           i + 1,
			InitialPosition: harborRemovedContainer.Position,
			FinalPosition:   shipPosition,
		})
	}

	return genome
}

// Mutates a genome getting two random displacements and swapping their final positions.
func Mutate(genome types.Genome) types.Genome {
	firstIndex, secondIndex := RandRange(0, len(genome)-1), RandRange(0, len(genome)-1)

	for firstIndex == secondIndex {
		secondIndex = RandRange(0, len(genome)-1)
	}

	genome[firstIndex], genome[secondIndex] = genome[secondIndex], genome[firstIndex]
	genome[firstIndex].Order, genome[secondIndex].Order = genome[secondIndex].Order, genome[firstIndex].Order

	return genome
}
