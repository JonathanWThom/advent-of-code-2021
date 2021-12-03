package day_3

import (
	"log"
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetPowerConsumption(path string) int64 {
	splitBits := splitBits(path)

	gammaBits := getGammaBits(splitBits)
	episolonBits := getEpsilonBits(gammaBits)

	return bitsToDecimal(gammaBits) * bitsToDecimal(episolonBits)
}

func splitBits(path string) [][]string {
	bitLines := helpers.ReadLines(path)
	var splitBits [][]string

	for _, bitLine := range bitLines {
		splitBitLine := strings.Split(bitLine, "")
		splitBits = append(splitBits, splitBitLine)
	}

	return splitBits
}

func getGammaBits(splitBits [][]string) string {
	length := len(splitBits[0])
	var final []string
	var oneCount int
	var zeroCount int

	for i := 0; i < length; i++ {
		for _, bitLine := range splitBits {
			if bitLine[i] == "0" {
				zeroCount += 1
			} else {
				oneCount += 1
			}
		}

		if oneCount > zeroCount {
			final = append(final, "1")
		} else {
			final = append(final, "0")
		}

		oneCount = 0
		zeroCount = 0
	}

	return strings.Join(final, "")
}

func getEpsilonBits(gammaBits string) string {
	var flipped []string
	split := strings.Split(gammaBits, "")

	for _, s := range split {
		if s == "0" {
			flipped = append(flipped, "1")
		} else {
			flipped = append(flipped, "0")
		}
	}

	return strings.Join(flipped, "")
}

func bitsToDecimal(bits string) int64 {
	dec, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return dec
}
