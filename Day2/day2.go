package Day2

import (
	"AdventOfCode/Utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day2:")
	data := Utils.ReadFileAsString("./Day2/data.txt")
	fmt.Println("Position Part1:", CalculatePosition(data, false))
	fmt.Println("Position Part2:", CalculatePosition(data, true))
	fmt.Println("----------")
}

func CalculatePosition(moves []string, aimMode bool) int {
	horizontal := 0
	vertical := 0
	aim := 0
	for _, move := range moves {
		currentMove := strings.Fields(move)
		distance, _ := strconv.Atoi(currentMove[1])
		switch {
		case currentMove[0] == "forward":
			horizontal += distance
			if aimMode {
				vertical += aim * distance
			}
		case currentMove[0] == "up":
			if aimMode {
				aim -= distance
			} else {
				vertical -= distance
			}
		case currentMove[0] == "down":
			if aimMode {
				aim += distance
			} else {
				vertical += distance
			}
		}

	}
	return horizontal * vertical
}
