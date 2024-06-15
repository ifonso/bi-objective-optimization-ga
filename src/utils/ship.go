package utils

import (
	"math"

	"github.com/ifonso/bi-objective-optimization-ga/src/types"
)

// GravityCenter calculates the gravity center of the containers
func gravityCenter(containers []types.Container) float64 {
	length := 6.0
	width := 2.4
	height := 2.6
	containerWeight := 10.0

	var x, y, z float64

	for _, container := range containers {
		x += float64(container.Position.X) * (length / 2) * containerWeight
		y += float64(container.Position.Y) * (width / 2) * containerWeight
		z += float64(container.Position.Z) * (height / 2) * containerWeight
	}

	totalWeight := float64(len(containers)) * containerWeight

	cgX := x / totalWeight
	cgY := y / totalWeight
	cgZ := z / totalWeight

	return math.Sqrt(cgX*cgX + cgY*cgY + cgZ*cgZ)
}

// GetMetacenter calculates the metacenter of the ship
func getMetacenter(length, width, draft float64) float64 {
	momentOfInertia := (width * math.Pow(length, 3)) / 12
	volumeOfDisplacement := width * length * draft
	return momentOfInertia / volumeOfDisplacement
}

// GetMetacentricHeight calculates the metacentric height of the ship with given containers
func GetMetacentricHeight(containers []types.Container) float64 {
	return gravityCenter(containers) - getMetacenter(18, 9.6, 4)
}
