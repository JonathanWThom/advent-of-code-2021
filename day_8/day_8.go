package day_8

import (
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

var uniques = []int{2, 3, 4, 7}

func GetUniqueDigits(path string) int {
	var counter int
	lines := helpers.ReadLines(path)
	for _, line := range lines {
		outputs := strings.Split(line, " | ")[1]
		for _, output := range strings.Split(outputs, " ") {
			if contains(uniques, len(output)) {
				counter += 1
			}

		}
	}

	return counter
}

func contains(i []int, integer int) bool {
	for _, v := range i {
		if v == integer {
			return true
		}
	}

	return false
}
