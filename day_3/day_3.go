package day_3

import (
	"log"
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetPowerConsumption(path string) int64 {
	splitBits := splitBits(path)

	gammaBits := getMostCommonBits(splitBits)
	episolonBits := getEpsilonBits(gammaBits)

	return bitsToDecimal(gammaBits) * bitsToDecimal(episolonBits)
}

func GetLifeSupportRating(path string) int64 {
	splitBits := splitBits(path)

	o2 := getOxygenRating(splitBits)
	co2 := getScrubberRating(splitBits)

	return o2 * co2
}

type comparison func(string, string) bool

func getRating(splitBits [][]string, comparisonFunc comparison) int64 {
	length := len(splitBits[0])
	var final int64

	for i := 0; i < length; i++ {
		var toKeep [][]string
		mostCommonBit := strings.Split(getMostCommonBits(splitBits), "")[i]
		for _, bitLine := range splitBits {
			if comparisonFunc(bitLine[i], mostCommonBit) {
				toKeep = append(toKeep, bitLine)
			}
		}

		if len(toKeep) == 1 {
			final = bitsToDecimal(strings.Join(toKeep[0], ""))
			break
		}

		splitBits = toKeep
	}

	return final
}

func getScrubberRating(splitBits [][]string) int64 {
	compare := func(bitLine string, mostCommonBit string) bool {
		return bitLine != mostCommonBit
	}

	return getRating(splitBits, compare)
}

func getOxygenRating(splitBits [][]string) int64 {
	compare := func(bitLine string, mostCommonBit string) bool {
		return bitLine == mostCommonBit
	}

	return getRating(splitBits, compare)
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

func getMostCommonBits(splitBits [][]string) string {
	length := len(splitBits[0])
	var final []string

	for i := 0; i < length; i++ {
		var oneCount int
		var zeroCount int
		for _, bitLine := range splitBits {
			if bitLine[i] == "0" {
				zeroCount += 1
			} else {
				oneCount += 1
			}
		}

		if oneCount >= zeroCount {
			final = append(final, "1")
		} else {
			final = append(final, "0")
		}
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
