package day_2

import (
	"log"
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

type instruction struct {
	direction string
	magnitude int
}

func GetFinalPositionProduct(path string) int {
	instructions := createInstructionsFromFile(path)
	var depth int
	var horizontal int

	for _, instruction := range instructions {
		switch instruction.direction {
		case "down":
			depth += instruction.magnitude
		case "up":
			depth -= instruction.magnitude
		case "forward":
			horizontal += instruction.magnitude
		}
	}

	return depth * horizontal
}

func createInstructionsFromFile(path string) []instruction {
	rawInstructions := helpers.ReadLines(path)
	var instructions []instruction

	for _, raw := range rawInstructions {
		rawParsed := strings.Split(raw, " ")
		magnitude, err := strconv.Atoi(rawParsed[1])
		if err != nil {
			log.Fatal(err)
		}

		inst := instruction{direction: rawParsed[0], magnitude: magnitude}
		instructions = append(instructions, inst)
	}

	return instructions
}
