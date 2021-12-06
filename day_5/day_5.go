package day_5

import (
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetOverlapCount(path string) int {
	coordinates := getCoordinates(path)
	ventField := newField(coordinates)
	ventField.drawPaths()

	return ventField.getMultiIntersectionCount()
}

type field struct {
	coordinates       []coordSet
	ventCoordCounters map[coord]int
}

func newField(coordinates []coordSet) field {
	return field{coordinates: coordinates, ventCoordCounters: map[coord]int{}}
}

func (f *field) getMultiIntersectionCount() int {
	var count int
	for _, ventCoord := range f.ventCoordCounters {
		if ventCoord > 1 {
			count += 1
		}
	}

	return count
}

// clean this up...
func (f *field) drawPaths() {
	for _, coordSet := range f.coordinates {
		if coordSet.start.x == coordSet.end.x {
			if coordSet.start.y < coordSet.end.y {
				for i := coordSet.start.y; i <= coordSet.end.y; i++ {
					f.ventCoordCounters[coord{x: coordSet.start.x, y: i}] += 1
				}
			} else {
				for i := coordSet.start.y; i >= coordSet.end.y; i-- {
					f.ventCoordCounters[coord{x: coordSet.start.x, y: i}] += 1
				}
			}

		} else if coordSet.start.y == coordSet.end.y {
			if coordSet.start.x < coordSet.end.x {
				for i := coordSet.start.x; i <= coordSet.end.x; i++ {
					f.ventCoordCounters[coord{y: coordSet.start.y, x: i}] += 1
				}
			} else {
				for i := coordSet.start.x; i >= coordSet.end.x; i-- {
					f.ventCoordCounters[coord{y: coordSet.start.y, x: i}] += 1
				}
			}
		}
	}
}

type coord struct {
	x int
	y int
}

type coordSet struct {
	start coord
	end   coord
}

func getCoordinates(path string) []coordSet {
	lines := helpers.ReadLines(path)

	var coordinates []coordSet
	for _, line := range lines {
		var set coordSet

		rawCoords := strings.Split(line, " -> ")
		for i, rawCoord := range rawCoords {
			xy := strings.Split(rawCoord, ",")
			if i == 0 {
				set.start = coord{x: helpers.ToI(xy[0]), y: helpers.ToI(xy[1])}
			} else {
				set.end = coord{x: helpers.ToI(xy[0]), y: helpers.ToI(xy[1])}
			}
		}

		coordinates = append(coordinates, set)
	}

	return coordinates
}
