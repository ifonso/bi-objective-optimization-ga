package utils

import "github.com/ifonso/bi-objective-optimization-ga/src/types"

// Returns a copy of the slice without duplicates.
func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}

	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}

// Checks if a value is in a slice.
func Contains[T comparable](sliceList []T, value T) bool {
	for _, a := range sliceList {
		if a == value {
			return true
		}
	}

	return false
}

// Checks if a genome is already in the population.
func ContainsGenome(population types.Population, individual types.Genome) bool {
	for _, a := range population {
		if a.Equals(individual) {
			return true
		}
	}

	return false
}
