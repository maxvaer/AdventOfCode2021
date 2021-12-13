package Day13

import (
	"AdventOfCode/Utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day13:")
	points, folds, maxX, maxY := getInstructionsFromFile("./Day13/data.txt")
	board := initBoard(maxY, maxX)
	board = drawPoints(board, points)
	counts, board := fold(board, folds, maxX, maxY)
	fmt.Println("Number of Dots:", counts)
	printBoard(board)
	fmt.Println("----------")
}

func getInstructionsFromFile(path string) ([][]int, []string, int, int) {
	var points [][]int
	var folds []string

	maxX := 0
	maxY := 0

	data := Utils.ReadFileAsString(path)

	isPoints := true

	for _, d := range data {
		if isPoints {
			if d == "" {
				isPoints = false
				continue
			}
			pointValues := strings.Split(d, ",")
			x, _ := strconv.Atoi(pointValues[0])
			if x > maxX {
				maxX = x
			}
			y, _ := strconv.Atoi(pointValues[1])
			if y > maxY {
				maxY = y
			}
			points = append(points, []int{x, y})
		} else {
			folds = append(folds, d)
		}

	}

	return points, folds, maxX + 1, maxY + 1
}

func drawPoints(board [][]string, points [][]int) [][]string {
	for _, point := range points {
		board[point[1]][point[0]] = "#"
	}

	return board
}

func fold(board [][]string, folds []string, maxX int, maxY int) ([]int, [][]string) {
	var newBoard [][]string
	var counts []int

	for _, s := range folds {
		verticalFold := true
		tmp := strings.Split(s, " ")
		tmp = strings.Split(tmp[2], "=")
		foldingPoint, _ := strconv.Atoi(tmp[1])

		if tmp[0] == "y" {
			verticalFold = false
		}

		if !verticalFold {
			//Horizontal fold
			maxY = len(board) - foldingPoint - 1
			newBoard = initBoard(maxY, maxX)
			//Copy Old
			for y := 0; y < len(newBoard); y++ {
				for x := 0; x < len(newBoard[0]); x++ {
					newBoard[y][x] = board[y][x]
				}
			}

			//Copy New
			for y := len(board) - 1; y > foldingPoint; y-- {
				for x := 0; x < len(newBoard[0]); x++ {
					newY := len(board) - 1 - y
					if newBoard[newY][x] == "." {
						newBoard[newY][x] = board[y][x]
					}
				}
			}

			board = newBoard
		} else {
			//vertical Fold
			maxX = len(board[0]) - foldingPoint - 1
			newBoard = initBoard(maxY, maxX)

			//Copy Old
			for y := 0; y < len(newBoard); y++ {
				for x := 0; x < len(newBoard[0]); x++ {
					newBoard[y][x] = board[y][x]
				}
			}

			//Copy New
			for y := 0; y < len(newBoard); y++ {
				for x := len(board[0]) - 1; x > foldingPoint; x-- {
					newX := len(board[0]) - 1 - x
					if newBoard[y][newX] == "." {
						newBoard[y][newX] = board[y][x]
					}
				}
			}
			board = newBoard
		}
		count := 0
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[0]); x++ {
				if board[y][x] == "#" {
					count++
				}
			}
		}
		counts = append(counts, count)
	}

	return counts, board
}

func printBoard(board [][]string) {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			value := board[y][x]
			if value == "." {
				fmt.Print(" ")
			} else {
				fmt.Print(value)
			}
		}
		fmt.Println("")
	}
}

func initBoard(maxX int, maxY int) [][]string {
	a := make([][]string, maxX)
	for i, _ := range a {
		a[i] = make([]string, maxY)
	}

	for y := 0; y < len(a); y++ {
		for x := 0; x < len(a[0]); x++ {
			a[y][x] = "."
		}
	}

	return a
}
