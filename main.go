package main

import (
	"fmt"

	"github.com/ifonso/bi-objective-optimization-ga/src/genetic"
	"github.com/ifonso/bi-objective-optimization-ga/src/utils"
)

func RunEvolution(populationSize int, generations int) {
	harbor := utils.GenerateHarbor()
	population := genetic.GeneratePopulation(harbor, populationSize)

	for i := 0; i < generations; i++ {
		fmt.Printf("\nGeneration: %d\n", i)
		genetic.PrintGenome(population[0], false)
		population = genetic.DoGeneration(population)
	}

	genetic.PrintGenome(population[0], true)
}

func main() {
	RunEvolution(1000, 100)
}
