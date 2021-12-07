package Day7

import (
	"AdventOfCode/Utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day7:")
	positions, maxPos := getPositionsFromFile("./Day7/data.txt")
	cost, position := calculateFuelCostForPositions(positions, maxPos, false)
	costSum, positionSum := calculateFuelCostForPositions(positions, maxPos, true)
	fmt.Println("Cheapest Position:", position, " Cost:", cost)
	fmt.Println("Cheapest Position with Sum:", positionSum, " Cost:", costSum)
	fmt.Println("----------")
}

func getPositionsFromFile(path string) ([]uint, uint) {
	var positions []uint
	maxPox := uint(0)

	data := Utils.ReadFileAsString(path)
	positionData := strings.Split(data[0], ",")
	for _, positionValue := range positionData {
		value, _ := strconv.Atoi(positionValue)
		if value >= int(maxPox) {
			maxPox = uint(value)
		}
		positions = append(positions, uint(value))
	}

	return positions, maxPox
}

func calculateFuelCostForPositions(positions []uint, max uint, sum bool) (uint, uint) {
	cheapestCost := ^uint(0)
	cheapestPosition := ^uint(0)
	for i := 0; i < int(max); i++ {
		totalCost := uint(0)
		for _, position := range positions {
			if sum {
				totalCost += calculateFuelCostSum(int(position), i)
			} else {
				totalCost += calculateFuelCost(int(position), i)
			}
		}
		if totalCost < cheapestCost {
			cheapestCost = totalCost
			cheapestPosition = uint(i)
		}
	}

	return cheapestCost, cheapestPosition
}

func calculateFuelCost(x int, y int) uint {
	return uint(math.Abs(float64(x - y)))
}

func calculateFuelCostSum(x int, y int) uint {
	cost := uint(0)

	difference := uint(math.Abs(float64(x - y)))
	for i := 0; i <= int(difference); i++ {
		cost += uint(i)
	}

	return cost
}
