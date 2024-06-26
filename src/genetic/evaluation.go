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
		metacentricHeight := utils.GetMetacentricHeight(shipOrdering[:i+1])
		if metacentricHeight < 0 {
			totalStability += metacentricHeight * 5
		} else {
			totalStability += metacentricHeight
		}
	}

	return totalStability
}

// Returns the fitness of a genome, which is the sum of the distances of all displacements.
func FitnessMoviments(genome types.Genome) float64 {
	totalMoviments := 0.0

	for _, displacement := range genome {
		totalMoviments += displacement.MovimentDistance()
	}

	penality := float64(float64(utils.InvalidOrderingCount(genome))/float64(len(genome))) * totalMoviments

	return totalMoviments + penality
}

// Returns the best genome from a tournament selection.
func TournamentSelection(population types.Population, tournamentSize int) (types.Genome, types.Genome) {
	if tournamentSize > len(population) {
		panic("Tournament size is greater than the population size")
	}

	selected := make(types.Population, tournamentSize)

	for i := 0; i < tournamentSize; i++ {
		selected[i] = population[utils.RandRange(0, len(population)-1)]
	}

	selected = RankIndividuals(selected)

	return selected[0], selected[1]
}

// Returns different parents based on tournament selection.
func GetParents(population types.Population, tournamentSize int) (types.Genome, types.Genome) {
	parent1, parent2 := TournamentSelection(population, tournamentSize)
	return parent1, parent2
}
