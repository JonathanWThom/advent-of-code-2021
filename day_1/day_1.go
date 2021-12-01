package day_1

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []int
	for scanner.Scan() {
		integer, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, integer)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
