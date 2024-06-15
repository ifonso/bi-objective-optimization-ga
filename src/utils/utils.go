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

// Removes an element from a slice.
func RemoveElement[T types.Equatable[T]](sliceList []T, element T) []T {
	for i, a := range sliceList {
		if a.Equals(element) {
			return append(sliceList[:i], sliceList[i+1:]...)
		}
	}

	return sliceList
}

type Normalizeble interface {
	~int | ~float64
}

// Normalizes a slice of values.
func Normalize[T Normalizeble](values []T) []T {
	if len(values) == 0 {
		panic("Cannot normalize an empty list")
	}

	minValue := values[0]
	maxValue := values[0]

	// Find the min and max values
	for _, value := range values {
		if value < minValue {
			minValue = value
		}

		if value > maxValue {
			maxValue = value
		}
	}

	if minValue == maxValue {
		return make([]T, len(values))
	}

	// Calculate the range once
	rangeValue := maxValue - minValue

	normalized := make([]T, len(values))

	// Normalize the values
	for i, value := range values {
		normalized[i] = (value - minValue) / rangeValue
	}

	return normalized
}

func OrderBy[T types.Equatable[T]](sliceList []T, less func(a, b T) bool) []T {
	sorted := make([]T, len(sliceList))
	copy(sorted, sliceList)

	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if less(sorted[j], sorted[i]) {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	return sorted
}
