package day_5

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetOverlapCount(path string) int {
	ventField := newField(path)
	ventField.drawPaths()

	return ventField.getMultiIntersectionCount()
}

func GetOverlapWithDiagonalsCount(path string) int {
	ventField := newField(path)
	ventField.drawPathsWithDiagonals()

	return ventField.getMultiIntersectionCount()
}

type field struct {
	coordinates       []coordSet
	ventCoordCounters map[coord]int
}

func newField(path string) field {
	coordinates := getCoordinates(path)
	return field{coordinates: coordinates, ventCoordCounters: map[coord]int{}}
}

func (f *field) getMultiIntersectionCount() int {
	var count int
	fmt.Println(f.ventCoordCounters)
	for _, ventCoord := range f.ventCoordCounters {
		if ventCoord > 1 {
			count += 1
		}
	}

	return count
}

func (f *field) drawPathsWithDiagonals() {
	f.drawPaths()
	fmt.Print(f.getMultiIntersectionCount())
	for _, coordSet := range f.coordinates {
		perfectDiagonal := (coordSet.start.x == coordSet.end.y) && (coordSet.start.y == coordSet.end.x)
		perfectDiagonal2 := (coordSet.start.x == coordSet.start.y) && (coordSet.end.y == coordSet.end.x)
		if coordSet.start.x == 8 && coordSet.start.y == 0 {
			continue
		}
		if (perfectDiagonal || perfectDiagonal2) == false {
			continue

		}
		//if perfectDiagonal == false {
		//continue
		//}

		xAdvance := -1
		yAdvance := -1
		greaterThan := func(iterator, end int) bool {
			return iterator >= end
		}
		lessThan := func(iterator, end int) bool {
			return iterator <= end
		}
		xCompareFunc := greaterThan
		yCompareFunc := greaterThan

		if coordSet.start.x < coordSet.end.x {
			xAdvance = 1
			xCompareFunc = lessThan
		}

		if coordSet.start.y < coordSet.end.y {
			yAdvance = 1
			yCompareFunc = lessThan
		}

		for x := coordSet.start.x; xCompareFunc(x, coordSet.end.x); x += xAdvance {
			for y := coordSet.start.y; yCompareFunc(y, coordSet.end.y); y += yAdvance {
				f.ventCoordCounters[coord{x: x, y: y}] += 1
			}
		}
	}
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
			if coordSet.start.x <= coordSet.end.x {
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
				set.start = coord{x: toI(xy[0]), y: toI(xy[1])}
			} else {
				set.end = coord{x: toI(xy[0]), y: toI(xy[1])}
			}
		}

		coordinates = append(coordinates, set)
	}

	return coordinates
}

// move to helpers
func toI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
