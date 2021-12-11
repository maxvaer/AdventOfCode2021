package Day11

import (
	"AdventOfCode/Utils"
	"fmt"
	"strconv"
)

func Run() {
	fmt.Println("Day11:")
	board := getPositionsFromFile("./Day11/data.txt")
	fmt.Println("Count Flashes:", calculateFlashes(board, 100))
	board = getPositionsFromFile("./Day11/data.txt")
	fmt.Println("First Sync Flash :", calculateSyncFlash(board))
	fmt.Println("----------")
}

func calculateSyncFlash(board [][]int) int {
	sum := 0
	boardSize := len(board) * len(board[0])
	step := 1

	for {
		sum = 0
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[0]); x++ {
				flash(board, x, y)
			}
		}

		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[0]); x++ {
				if board[y][x] > 9 {
					board[y][x] = 0
					sum++
				}
			}
		}

		if sum == boardSize {
			break
		}
		step++
	}

	return step
}

func calculateFlashes(board [][]int, steps int) int {
	sum := 0
	for i := 0; i < steps; i++ {
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[0]); x++ {
				sum += flash(board, x, y)
			}
		}

		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[0]); x++ {
				if board[y][x] > 9 {
					board[y][x] = 0
				}
			}
		}
	}

	return sum
}

func flash(board [][]int, x int, y int) int {
	maxY := len(board) - 1
	maxX := len(board[0]) - 1

	minRow := y == 0
	maxRow := y == maxY

	minCol := x == 0
	maxCol := x == maxX

	energy := board[y][x]
	if energy > 9 {
		return 0
	}

	board[y][x]++
	energy = board[y][x]
	if energy < 10 {
		return 0
	}
	count := 1

	if !minRow {
		count += flash(board, x, y-1)
	}
	if !minCol {
		count += flash(board, x-1, y)
	}
	if !maxRow {
		count += flash(board, x, y+1)
	}
	if !maxCol {
		count += flash(board, x+1, y)
	}

	if !minRow && !minCol {
		count += flash(board, x-1, y-1)
	}
	if !minRow && !maxCol {
		count += flash(board, x+1, y-1)
	}
	if !maxRow && !minCol {
		count += flash(board, x-1, y+1)
	}
	if !maxRow && !maxCol {
		count += flash(board, x+1, y+1)
	}

	return count
}

func getPositionsFromFile(path string) [][]int {
	data := Utils.ReadFileAsString(path)
	board := initBoard(len(data), len(data[0]))

	for y := 0; y < len(board); y++ {
		row := data[y]
		for x := 0; x < len(board[0]); x++ {
			tmp, _ := strconv.Atoi(string(row[x]))
			board[y][x] = tmp
		}
	}

	return board
}

func initBoard(maxX int, maxY int) [][]int {
	a := make([][]int, maxX)
	for i, _ := range a {
		a[i] = make([]int, maxY)
	}
	return a
}
