package Day4

import (
	"AdventOfCode/Utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day4:")
	data := Utils.ReadFileAsString("./Day4/data.txt")
	fmt.Println("First Bingo score: ", getScore(data, false))
	fmt.Println("Last Bingo score: ", getScore(data, true))
	fmt.Println("----------")
}

func getScore(input []string, last bool) uint {
	moves := ^uint(0)
	movesLast := uint(0)
	var score uint

	chosenInput := input[0]
	chosen := formatChosen(chosenInput)

	var game Game
	var rowIndex uint
	for i := 2; i <= len(input); i++ {
		if i == 2 {
			game = newGame(chosen)
			rowIndex = 0
		}
		if rowIndex == 5 {
			results := game.play()
			if results[0] != 0 {
				if last {
					if results[1] > movesLast {
						movesLast = results[1]
						score = results[0]
					}
				} else {
					if results[1] < moves {
						moves = results[1]
						score = results[0]
					}
				}
			}
			game = newGame(chosen)
			rowIndex = 0
		} else {
			fields := strings.Fields(input[i])
			var row = []uint{}
			for _, numberValue := range fields {
				value, _ := strconv.Atoi(numberValue)
				row = append(row, uint(value))
			}
			game.addRow(row, rowIndex)
			rowIndex++
		}
	}
	return score
}

func formatChosen(chosenInput string) []uint {
	var chosen = []uint{}
	numbersInput := strings.Split(chosenInput, ",")

	for _, numberValue := range numbersInput {
		value, _ := strconv.Atoi(numberValue)
		chosen = append(chosen, uint(value))
	}

	return chosen
}

type number struct {
	value  uint
	marked bool
}

type Game struct {
	chosen []uint
	board  [5][5]number
	bingo  bool
	score  uint
}

func newGame(chosen []uint) Game {
	g := Game{chosen: chosen}
	g.bingo = false
	g.score = 0
	return g
}

func (g *Game) addRow(row []uint, rowIndex uint) {
	var numberRow [5]number
	for i, value := range row {
		numberRow[i].value = value
		numberRow[i].marked = false
	}
	g.board[rowIndex] = numberRow
}

func (g *Game) play() []uint {
	score := uint(0)
	scoreIndex := uint(0)
	for index, chosenNumber := range g.chosen {
		g.markNumber(chosenNumber)
		g.checkBingo(true)
		g.checkBingo(false)
		if g.bingo {
			g.calculateUnmarkedSum()
			score = g.score * chosenNumber
			scoreIndex = uint(index)
			break
		}
	}
	return []uint{score, scoreIndex}
}

func (g *Game) markNumber(number uint) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if g.board[i][j].value == number {
				g.board[i][j].marked = true
			}
		}
	}
}

func (g *Game) checkBingo(vertical bool) {
	for i := 0; i < 5; i++ {
		markedCounter := 0
		for j := 0; j < 5; j++ {
			if vertical {
				if g.board[j][i].marked == true {
					markedCounter++
				}
			} else {
				if g.board[i][j].marked == true {
					markedCounter++
				}
			}
		}
		if markedCounter == 5 {
			g.bingo = true
			break
		}
	}
}

func (g *Game) calculateUnmarkedSum() {
	unMarkedSum := uint(0)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if g.board[i][j].marked == false {
				unMarkedSum += g.board[i][j].value
			}
		}
	}
	g.score = unMarkedSum
}
