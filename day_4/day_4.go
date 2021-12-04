package day_4

import (
	"log"
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2021/helpers"
)

func GetBingoScore(path string) int {
	return getBingoScore(path, getBoardAndLastNum)
}

func GetLosingBingoScore(path string) int {
	return getBingoScore(path, getLosingBoardAndLastNum)
}

type boardWinnerFunc func([]board, []int) (board, int)

func getBingoScore(path string, getBoardWinner boardWinnerFunc) int {
	lines := helpers.ReadLines(path)
	bingoNums := getBingoNums(lines)
	boards := buildBoards(lines)
	board, bingoNum := getBoardWinner(boards, bingoNums)
	unmarkedSum := board.getUnmarkedSum()

	return bingoNum * unmarkedSum
}

func getBingoNums(lines []string) []int {
	var bingoNums []int
	for _, num := range strings.Split(lines[0], ",") {
		bingoNum, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		bingoNums = append(bingoNums, bingoNum)
	}

	return bingoNums
}

func (b *board) getUnmarkedSum() int {
	var sum int
	for i, square := range b.squares {
		addVal := true
		for _, markedIndex := range b.markedIndexes {
			if i == markedIndex {
				addVal = false
				break
			}
		}

		if addVal {
			sum += square
		}
	}

	return sum
}

func getBoardAndLastNum(boards []board, bingoNums []int) (board, int) {
	var b board
	var lastNum int

checkBingo:
	for _, bingoNum := range bingoNums {
		for i, board := range boards {
			board.markBingoNum(bingoNum, i)
			boards[i] = board
		}

		for _, board := range boards {
			if board.hasNewBingo() {
				b = board
				lastNum = bingoNum
				break checkBingo
			}
		}
	}

	return b, lastNum
}

func getLosingBoardAndLastNum(boards []board, bingoNums []int) (board, int) {
	var b board
	var lastNum int
	var boardsWithBingo int

checkBingo:
	for _, bingoNum := range bingoNums {
		for i, board := range boards {
			board.markBingoNum(bingoNum, i)
			boards[i] = board
		}

		for i, board := range boards {
			if board.hasNewBingo() {
				board.bingo = true
				boards[i] = board
				boardsWithBingo += 1

				if boardsWithBingo == len(boards) {
					b = board
					lastNum = bingoNum
					break checkBingo
				}
			}
		}
	}

	return b, lastNum
}

type board struct {
	squares       []int
	markedIndexes []int
	bingo         bool
}

var bingoWinners = [][]int{
	{0, 1, 2, 3, 4},
	{5, 6, 7, 8, 9},
	{10, 11, 12, 13, 14},
	{15, 16, 17, 18, 19},
	{20, 21, 22, 23, 24},
	{0, 5, 10, 15, 20},
	{1, 6, 11, 16, 21},
	{2, 7, 12, 17, 22},
	{3, 8, 13, 18, 23},
	{4, 9, 14, 19, 24},
}

func (b *board) hasNewBingo() bool {
	if b.bingo {
		return false
	}

	for _, winnerSet := range bingoWinners {
		var matches int
		for _, winner := range winnerSet {
			for _, markedIndex := range b.markedIndexes {
				if markedIndex == winner {
					matches += 1
					if matches == 5 {
						b.bingo = true
						return true
					}
				}
			}
		}
	}

	return false
}

func (b *board) markBingoNum(num int, j int) {
	for i, square := range b.squares {
		if square == num {
			b.markedIndexes = append(b.markedIndexes, i)
		}
	}
}

func buildBoards(lines []string) []board {
	var boards []board
	b := board{}

	for _, line := range lines[2:] {
		if line != "" {
			for _, num := range strings.Split(line, " ") {
				if num != "" {
					square, err := strconv.Atoi(num)
					if err != nil {
						log.Fatal(err)
					}
					b.squares = append(b.squares, square)
				}
			}

			if len(b.squares) == 25 {
				boards = append(boards, b)
				b = board{}
			}
		}
	}

	return boards
}
