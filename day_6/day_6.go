package day_6

import (
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetLaternfishCount(path string) int {
	fishes := getStartingFishes(path)

	for i := 0; i < 80; i++ {
		for i, f := range fishes {
			f, addNewFish := f.tick()
			fishes[i] = f
			if addNewFish {
				fishes = append(fishes, fish{timer: 8})
			}
		}
	}

	return len(fishes)
}

type fish struct {
	timer int
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
