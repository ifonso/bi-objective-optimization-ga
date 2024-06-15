package genetic

import (
	"fmt"

	"github.com/ifonso/bi-objective-optimization-ga/src/types"
)

func PrintGenome(genome types.Genome, printAllData bool) {
	// Print moviment and stability
	fmt.Printf("Moviment: %.2f\n", FitnessMoviments(genome))
	fmt.Printf("Stability: %.2f\n", FitnessStability(genome))

	// Print genome
	if printAllData {
		for _, displacement := range genome {
			fmt.Printf("%s -> %s\n", displacement.InitialPositionString(), displacement.FinalPositionString())
		}
	}
}
