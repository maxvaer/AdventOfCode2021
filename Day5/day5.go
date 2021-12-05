package Day5

import (
	"AdventOfCode/Utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day5:")
	horizontal := true
	lines := getLines("./Day5/data.txt", horizontal)
	maxX, maxY := getMaxXY(lines)
	board := initBoard(maxX, maxY)
	draw(board, lines)
	fmt.Println("Count:", countOverlaps(board))
	horizontal = false
	lines = getLines("./Day5/data.txt", horizontal)
	board = initBoard(maxX, maxY)
	draw(board, lines)
	fmt.Println("Count with Diagonal:", countOverlaps(board))
	fmt.Println("----------")
}

type Point struct {
	x uint
	y uint
}

type Line struct {
	start Point
	end   Point
}

func getLines(path string, onlyHorizontalAndVertical bool) []Line {
	data := Utils.ReadFileAsString(path)
	var lines []Line
	for _, dataValue := range data {
		var line Line
		pointsPair := strings.Split(dataValue, " -> ")
		for pointsIndex, pointsPairValue := range pointsPair {
			pointsValue := strings.Split(pointsPairValue, ",")
			var point Point
			for pointIndex, pointString := range pointsValue {
				pointValue, _ := strconv.Atoi(pointString)
				if pointIndex == 0 {
					point.x = uint(pointValue)
				} else {
					point.y = uint(pointValue)
				}
			}
			if pointsIndex == 0 {
				line.start = point
			} else {
				line.end = point
			}
		}

		if onlyHorizontalAndVertical {
			if line.start.x == line.end.x || line.start.y == line.end.y {
				lines = append(lines, line)
			}
		} else {
			lines = append(lines, line)
		}
	}
	return lines
}

func getMaxXY(lines []Line) (maxX uint, maxY uint) {
	maxX = uint(0)
	maxY = uint(0)
	for _, line := range lines {
		if line.start.x > maxX {
			maxX = line.start.x
		}
		if line.end.x > maxX {
			maxX = line.end.x
		}
		if line.start.y > maxY {
			maxY = line.start.y
		}
		if line.end.y > maxY {
			maxY = line.end.y
		}
	}

	return maxX, maxY
}

func initBoard(maxX uint, maxY uint) [][]uint {
	a := make([][]uint, maxX+1)
	for i, _ := range a {
		a[i] = make([]uint, maxY+1)
	}
	return a
}

func draw(board [][]uint, lines []Line) {
	for _, line := range lines {
		left := line.start.x > line.end.x
		up := line.start.y > line.end.y

		runX := line.start.x
		runY := line.start.y

		for {
			board[runX][runY]++
			if runX == line.end.x && runY == line.end.y {
				break
			}
			if runX != line.end.x {
				if left {
					runX--
				} else {
					runX++
				}
			}

			if runY != line.end.y {
				if up {
					runY--
				} else {
					runY++
				}
			}
		}
	}
}

func countOverlaps(board [][]uint) uint {
	counter := uint(0)

	for _, row := range board {
		for _, value := range row {
			if value > 1 {
				counter++
			}
		}
	}

	return counter
}
