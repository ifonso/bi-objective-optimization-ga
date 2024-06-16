package utils

import (
	"github.com/ifonso/bi-objective-optimization-ga/src/types"
)

// GravityCenter calculates the gravity center of the containers
func gravityCenter(containers []types.Container) (float64, float64, float64) {
	length := 6.0
	width := 2.4
	height := 2.6
	containerWeight := 10.0

	var x, y, z float64

	for _, container := range containers {
		x += float64(container.Position.X) * length * containerWeight
		y += float64(container.Position.Y) * width * containerWeight
		z += float64(container.Position.Z) * height * containerWeight
	}

	totalWeight := float64(len(containers)) * containerWeight

	if totalWeight == 0 {
		return 0, 0, 0
	}

	cgX := x / totalWeight
	cgY := y / totalWeight
	cgZ := z / totalWeight

	return cgX, cgY, cgZ
}

func GetMetacentricHeight(containers []types.Container) float64 {
	// Dimensions of the boat
	boatLength := 20.0
	boatWidth := 20.0
	boatHeight := 10.0
	boatWeight := 200.0

	// Center of gravity of containers
	cgX, cgY, cgZ := gravityCenter(containers)

	// Combined center of gravity (G') of the boat + containers
	totalContainerWeight := float64(len(containers)) * 10.0
	totalWeight := boatWeight + totalContainerWeight

	cgBoatX := boatLength / 2.0
	cgBoatY := boatWidth / 2.0
	cgBoatZ := boatHeight / 2.0

	combinedCgX := (cgBoatX*boatWeight + cgX*totalContainerWeight) / totalWeight
	combinedCgY := (cgBoatY*boatWeight + cgY*totalContainerWeight) / totalWeight
	combinedCgZ := (cgBoatZ*boatWeight + cgZ*totalContainerWeight) / totalWeight

	// Metacenter (M) calculation
	KB := boatHeight / 2.0
	BM := (boatWidth * boatWidth) / (12 * (boatLength / boatHeight))
	KM := KB + BM

	// Metacentric height (GM)
	GM := KM*1.5 - (combinedCgZ + combinedCgX + combinedCgY)
	return GM
}
