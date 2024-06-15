package types

import (
	"fmt"
	"math"
	"reflect"
)

type Equatable[T any] interface {
	Equals(other T) bool
}

// Domain Types

type Position struct {
	X int
	Y int
	Z int
}

func (p Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}

type Container struct {
	Position Position
}

func (c Container) Equals(other Container) bool {
	return c.Position.Equals(other.Position)
}

type Displacement struct {
	Order           int      // loading order
	InitialPosition Position // position in harbor
	FinalPosition   Position // position in ship
}

func (d Displacement) Equals(other Displacement) bool {
	return d.Order == other.Order &&
		d.InitialPosition.Equals(other.InitialPosition) &&
		d.FinalPosition.Equals(other.FinalPosition)
}

func (d Displacement) InitialPositionString() string {
	return fmt.Sprintf("(%d, %d, %d)", d.InitialPosition.X, d.InitialPosition.Y, d.InitialPosition.Z)
}

func (d Displacement) FinalPositionString() string {
	return fmt.Sprintf("(%d, %d, %d)", d.FinalPosition.X, d.FinalPosition.Y, d.FinalPosition.Z)
}

func (d Displacement) MovimentDistance() float64 {
	return (math.Abs(float64(5+d.FinalPosition.X-d.InitialPosition.X)) +
		math.Abs(float64(d.FinalPosition.Y-d.InitialPosition.Y)) +
		math.Abs(float64(d.FinalPosition.Z-d.InitialPosition.Z)))
}

// Genetic Algorithm Types

type Genome []Displacement

func (genome Genome) Equals(other Genome) bool {
	return reflect.DeepEqual(genome, other)
}

type Population = []Genome

// Used for ranking individuals in a population
type GenomeFitness struct {
	Genome Genome

	MovimentFitness  float64
	StabilityFitness float64

	Rank             int
	CrowdingDistance float64

	NormalizedMovimentFitness  float64
	NormalizedStabilityFitness float64
}

func (gf GenomeFitness) Equals(other GenomeFitness) bool {
	return gf.Genome.Equals(other.Genome) &&
		gf.MovimentFitness == other.MovimentFitness &&
		gf.StabilityFitness == other.StabilityFitness &&
		gf.Rank == other.Rank &&
		gf.CrowdingDistance == other.CrowdingDistance &&
		gf.NormalizedMovimentFitness == other.NormalizedMovimentFitness &&
		gf.NormalizedStabilityFitness == other.NormalizedStabilityFitness
}
