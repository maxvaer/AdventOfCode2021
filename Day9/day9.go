package Day9

import (
	"AdventOfCode/Utils"
	"fmt"
	"sort"
	"strconv"
)

func Run() {
	fmt.Println("Day9:")
	heightMap := getHeightMapFromFile("./Day9/data.txt")
	fmt.Println("Sum of LowPoints:", sumLowPoints(heightMap))
	fmt.Println("Basin score:", calculateBasinScore(heightMap))
	fmt.Println("----------")
}

func sumLowPoints(heightMap [][]int) int {
	sum := 0

	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[0]); x++ {
			if checkLowPoint(x, y, heightMap) {
				value := heightMap[y][x]
				value++
				sum += value
			}
		}
	}

	return sum
}

func calculateBasinScore(heightMap [][]int) int {

	heightMapCopy := make([][]int, len(heightMap))
	for i := range heightMap {
		heightMapCopy[i] = make([]int, len(heightMap[i]))
		copy(heightMapCopy[i], heightMap[i])
	}

	var sizes []int

	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[0]); x++ {
			if checkLowPoint(x, y, heightMap) {
				size := getSizeOfBasin(x, y, heightMapCopy)
				sizes = append(sizes, size)
			}
		}
	}

	sort.Ints(sizes)

	score := 1
	for i := len(sizes); i > len(sizes)-3; i-- {
		score *= sizes[i-1]
	}

	return score
}

func getHeightMapFromFile(path string) [][]int {
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

func getSizeOfBasin(x int, y int, board [][]int) int {
	maxY := len(board) - 1
	maxX := len(board[0]) - 1

	minRow := y == 0
	maxRow := y == maxY

	minCol := x == 0
	maxCol := x == maxX

	sum := 0

	if board[y][x] == 9 {
		return 0
	}

	sum += 1
	board[y][x] = 9
	if !minRow {
		sum += getSizeOfBasin(x, y-1, board)
	}
	if !minCol {
		sum += getSizeOfBasin(x-1, y, board)
	}
	if !maxRow {
		sum += getSizeOfBasin(x, y+1, board)
	}
	if !maxCol {
		sum += getSizeOfBasin(x+1, y, board)
	}

	return sum
}

func checkLowPoint(x int, y int, board [][]int) bool {
	maxY := len(board) - 1
	maxX := len(board[0]) - 1
	value := board[y][x]

	minRow := y == 0
	maxRow := y == maxY

	minCol := x == 0
	maxCol := x == maxX

	isLowPoint := true

	if !minRow {
		if board[y-1][x] <= value {
			isLowPoint = false
		}
	}
	if !minCol {
		if board[y][x-1] <= value {
			isLowPoint = false
		}
	}
	if !maxRow {
		if board[y+1][x] <= value {
			isLowPoint = false
		}
	}
	if !maxCol {
		if board[y][x+1] <= value {
			isLowPoint = false
		}
	}

	return isLowPoint
}

func initBoard(maxX int, maxY int) [][]int {
	a := make([][]int, maxX)
	for i, _ := range a {
		a[i] = make([]int, maxY)
	}
	return a
}
