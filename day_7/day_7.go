package day_7

import (
	"math"
	"sort"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetFuel(path string) int {
	return getFuel(path, getFuelTotal)
}

func GetNonConstantFuel(path string) int {
	return getFuel(path, getNonConstantFuelTotal)
}

type fuelTotalFunc func([]int, int) int

func getFuel(path string, fuelTotalCalc fuelTotalFunc) int {
	positions := getCrabPositions(path)
	var lowestFuelTotal int
	for i := 0; i < positions[len(positions)-1]; i++ {
		fuelTotal := fuelTotalCalc(positions, i)
		if i == 0 || fuelTotal < lowestFuelTotal {
			lowestFuelTotal = fuelTotal
		}
	}

	return lowestFuelTotal
}

func getNonConstantFuelTotal(positions []int, candidatePosition int) int {
	var total int
	for _, position := range positions {
		distance := int(math.Abs(float64(position - candidatePosition)))
		for i := 1; i <= distance; i++ {
			total += i
		}
	}

	return total
}

func getFuelTotal(positions []int, candidatePosition int) int {
	var total int
	for _, position := range positions {
		total += int(math.Abs(float64(position - candidatePosition)))
	}

	return total
}

func getCrabPositions(path string) []int {
	lines := helpers.ReadLines(path)
	var positions []int
	for _, position := range strings.Split(lines[0], ",") {
		positions = append(positions, helpers.ToI(position))
	}
	sort.Ints(positions)

	return positions
}
