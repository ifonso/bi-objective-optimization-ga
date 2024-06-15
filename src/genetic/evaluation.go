package genetic

import (
	"github.com/ifonso/bi-objective-optimization-ga/src/types"
	"github.com/ifonso/bi-objective-optimization-ga/src/utils"
)

// FitnessStability calculates the fitness of a genome based on the stability of the ship,
// the higher the value, the more stable the boat
func FitnessStability(genome types.Genome) float64 {
	totalStability := 0.0
	shipOrdering := make([]types.Container, len(genome))

	for i := 0; i < len(genome); i++ {
		shipOrdering[i] = types.Container{
			Position: genome[i].FinalPosition,
		}
	}

	for i := 0; i < len(shipOrdering); i++ {
		totalStability += utils.GetMetacentricHeight(shipOrdering[:i+1])
	}

	return totalStability
}

// Returns the fitness of a genome, which is the sum of the distances of all displacements.
func FitnessMoviments(genome types.Genome) float64 {
	totalMoviments := 0.0
	for _, displacement := range genome {
		totalMoviments += displacement.MovimentDistance()
	}
	return totalMoviments
}

// Ranking -----------------------------------------
