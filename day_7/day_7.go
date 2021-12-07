package day_7

import (
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetFuel(path string) int {
	positions := getCrabPositions(path)
	var lowestFuelTotal int
	for i := 0; i < positions[len(positions)-1]; i++ {
		fuelTotal := getFuelTotal(positions, i)
		if i == 0 || fuelTotal < lowestFuelTotal {
			lowestFuelTotal = fuelTotal
		}
	}

	return lowestFuelTotal
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
