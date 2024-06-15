package genetic

import (
	"math/rand/v2"

	"github.com/ifonso/bi-objective-optimization-ga/src/types"
)

// Generates a new population based on the current one.
func DoGeneration(population types.Population) types.Population {
	offspring := types.Population{}

	for len(offspring) < len(population) {
		// Select parents
		parent1, parent2 := GetParents(population, 10)

		// Crossover
		child1, child2 := Crossover(parent1, parent2)

		// Mutate
		if rand.Float64() < 0.05 {
			child1 = Mutate(child1)
			child2 = Mutate(child2)
		}

		offspring = append(offspring, child1, child2)
	}

	// Join population and offspring, then rank them
	population = append(population, offspring...)
	population = RankIndividuals(population)

	// Select the best individuals
	return population[:len(population)/2]
}
