package day_1

import (
	"log"
	"strconv"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetDepthIncreaseCount(path string) int {
	depths := getDepthsFromFile(path)
	var counter int
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			counter++
		}
	}

	return counter
}

func GetSlidingDepthIncreaseCount(path string) int {
	depths := getDepthsFromFile(path)
	var counter int
	for i := 1; i < len(depths)-2; i++ {
		previous_set_sum := depths[i-1] + depths[i] + depths[i+1]
		next_set_sum := depths[i] + depths[i+1] + depths[i+2]

		if next_set_sum > previous_set_sum {
			counter++
		}
	}

	return counter
}

func getDepthsFromFile(path string) []int {
	lines := helpers.ReadLines(path)
	var depths []int
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, depth)
	}

	return depths
}
