package day_6

import (
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetLaternfishCount(path string, days int) int {
	fishes := getStartingFishes(path)
	var fishCount = len(fishes)

	for _, f := range fishes {
		fishCount += f.tickFor(days)
	}

	return fishCount
}

type fish struct {
	timer int
}

func (f *fish) tickFor(days int) int {
	var descendents int
	for i := days; i > 0; i-- {
		var newFish bool
		_, newFish = f.tick()

		if newFish {
			descendents += 1
			descendent := fish{timer: 8}
			descendents += descendent.tickFor(i - 1)
		}
	}

	return descendents
}

func (f *fish) tick() (fish, bool) {
	var newFish bool
	if f.timer == 0 {
		f.timer = 6
		newFish = true
	} else {
		f.timer -= 1
	}

	return *f, newFish
}

func getStartingFishes(path string) []fish {
	var fishes []fish
	lines := helpers.ReadLines(path)
	for _, timer := range strings.Split(lines[0], ",") {
		f := fish{timer: helpers.ToI(timer)}
		fishes = append(fishes, f)
	}

	return fishes
}
