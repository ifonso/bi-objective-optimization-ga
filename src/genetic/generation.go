package genetic

import (
	"github.com/ifonso/bi-objective-optimization-ga/src/types"
	"github.com/ifonso/bi-objective-optimization-ga/src/utils"
)

// Generates a population of genomes with random displacements.
func GeneratePopulation(harbor []types.Container, populationSize int) types.Population {
	population := types.Population{}

	for len(population) < populationSize {
		genome := utils.GenerateRandomGenome(harbor)
		if !utils.ContainsGenome(population, genome) {
			population = append(population, genome)
		}
	}

	return population
}

// Generates a child genome from two parent genomes using crossover.
func GenerateChildGenome(father types.Genome, mother types.Genome, crossPoint int) types.Genome {
	// Collect initial and final positions from the father genome up to the cross point
	fatherInitialPositions := make(map[types.Position]bool)
	fatherFinalPositions := make(map[types.Position]bool)

	for i := 0; i < crossPoint; i++ {
		fatherInitialPositions[father[i].InitialPosition] = true
		fatherFinalPositions[father[i].FinalPosition] = true
	}

	// Collect initial and final positions from the mother genome that are not in the father's positions

	motherInitialPositions := []types.Position{}
	motherFinalPositions := []types.Position{}

	for _, displacement := range mother {
		if !fatherInitialPositions[displacement.InitialPosition] {
			motherInitialPositions = append(motherInitialPositions, displacement.InitialPosition)
		}
		if !fatherFinalPositions[displacement.FinalPosition] {
			motherFinalPositions = append(motherFinalPositions, displacement.FinalPosition)
		}
	}

	// Construct the child genome

	child := make(types.Genome, len(father))
	copy(child[:crossPoint], father[:crossPoint])

	for i := crossPoint; i < len(father); i++ {
		child[i] = types.Displacement{
			Order:           i + 1,
			InitialPosition: motherInitialPositions[i-crossPoint],
			FinalPosition:   motherFinalPositions[i-crossPoint],
		}
	}

	return child
}

// Crossover between two genomes. Returns two children.
func Crossover(firstParent types.Genome, secondParent types.Genome) (types.Genome, types.Genome) {
	if len(firstParent) != len(secondParent) {
		panic("Genomes must have the same length")
	}

	crossPoint := utils.RandRange(1, len(firstParent)-1)

	return GenerateChildGenome(firstParent, secondParent, crossPoint), GenerateChildGenome(secondParent, firstParent, crossPoint)
}
