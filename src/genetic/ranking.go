package genetic

import (
	"math"

	"github.com/ifonso/bi-objective-optimization-ga/src/types"
	"github.com/ifonso/bi-objective-optimization-ga/src/utils"
)

// Returns true if the genome is not dominated by any other genome in the population.
func IsNotDominated(genome types.GenomeFitness, population []types.GenomeFitness) bool {
	// If the individual is the only one in the population, it is not dominated.
	if len(population) == 1 {
		return true
	}

	for _, individual := range population {
		if individual.Equals(genome) {
			continue
		}

		if (individual.MovimentFitness <= genome.MovimentFitness && individual.StabilityFitness >= genome.StabilityFitness) &&
			(individual.MovimentFitness < genome.MovimentFitness || individual.StabilityFitness > genome.StabilityFitness) {
			return false
		}
	}

	return true
}

// Crowding Distance: The sum of the distances of the individual to the nearest neighbors.
func CalculateCrowdingDistanceForGroup(group []types.GenomeFitness) []types.GenomeFitness {
	orderedByMoviment := utils.OrderBy(group, func(a, b types.GenomeFitness) bool {
		return a.MovimentFitness < b.MovimentFitness
	})

	frontSize := len(orderedByMoviment)

	for i := 0; i < frontSize; i++ {
		if i == 0 || i == frontSize-1 {
			orderedByMoviment[i].CrowdingDistance = 100000
			continue
		} else {
			current := orderedByMoviment[i]
			previous := orderedByMoviment[i-1]
			next := orderedByMoviment[i+1]

			previusDistance := math.Pow(math.Pow(current.MovimentFitness-previous.MovimentFitness, 2)+math.Pow(current.StabilityFitness-previous.StabilityFitness, 2), 0.5)
			nextDistance := math.Pow(math.Pow(current.MovimentFitness-next.MovimentFitness, 2)+math.Pow(current.StabilityFitness-next.StabilityFitness, 2), 0.5)

			orderedByMoviment[i].CrowdingDistance = previusDistance + nextDistance
		}
	}

	return orderedByMoviment
}

// Calculate the normalized fitness for a group of individuals.
func CalculateNormalizedFitnessForGroup(group []types.GenomeFitness) []types.GenomeFitness {
	moviments := make([]float64, len(group))
	stability := make([]float64, len(group))

	for i, genome := range group {
		moviments[i] = genome.MovimentFitness
		stability[i] = genome.StabilityFitness
	}

	normalizedMoviments := utils.Normalize(moviments)
	normalizedStability := utils.Normalize(stability)

	for i, genome := range group {
		genome.NormalizedMovimentFitness = normalizedMoviments[i]
		genome.NormalizedStabilityFitness = normalizedStability[i]
	}

	return group
}

// Ranks the individuals in the population based on the non-dominated sorting algorithm
// and returns a new population sorted by rank.
func RankIndividuals(population types.Population) types.Population {
	remaningToBeRanked := make([]types.GenomeFitness, len(population))

	// - Calculate fitness values
	for i, individual := range population {
		remaningToBeRanked[i] = types.GenomeFitness{
			Genome:           individual,
			MovimentFitness:  FitnessMoviments(individual),
			StabilityFitness: FitnessStability(individual),

			Rank:             -1,
			CrowdingDistance: -1,

			NormalizedMovimentFitness:  -1,
			NormalizedStabilityFitness: -1,
		}
	}

	// - Rank individuals
	rank := 1
	ranked := []types.GenomeFitness{}

	for len(remaningToBeRanked) > 0 {
		toRemove := []types.GenomeFitness{}

		for _, genome := range remaningToBeRanked {
			if IsNotDominated(genome, remaningToBeRanked) {
				toRemove = append(toRemove, genome)
				genome.Rank = rank
				ranked = append(ranked, genome)
			}
		}

		for _, genome := range toRemove {
			remaningToBeRanked = utils.RemoveElement(remaningToBeRanked, genome)
		}

		rank++
	}

	// - Group by rank
	groupedByRank := make([][]types.GenomeFitness, rank-1)

	for _, genome := range ranked {
		groupedByRank[genome.Rank-1] = append(groupedByRank[genome.Rank-1], genome)
	}

	// Calculate normalized fitness values
	for i, group := range groupedByRank {
		groupedByRank[i] = CalculateNormalizedFitnessForGroup(group)
	}

	// Calculate crowding distance
	for i, group := range groupedByRank {
		groupedByRank[i] = CalculateCrowdingDistanceForGroup(group)
	}

	// Sort groups by crowding distance
	for i, group := range groupedByRank {
		groupedByRank[i] = utils.OrderBy(group, func(a, b types.GenomeFitness) bool {
			return a.CrowdingDistance > b.CrowdingDistance
		})
	}

	// - Order population

	orderedPopulation := []types.Genome{}

	for _, group := range groupedByRank {
		for _, genome := range group {
			orderedPopulation = append(orderedPopulation, genome.Genome)
		}
	}

	return orderedPopulation
}
