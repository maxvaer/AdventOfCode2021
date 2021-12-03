package Day3

import (
	"AdventOfCode/Utils"
	"fmt"
	"math"
	"strconv"
)

func Run() {
	data := Utils.ReadFileAsString("./Day3/test.txt")
	fmt.Println(CalculatePower(data))
	fmt.Println(CalculateLifeSupport(data))
}

func CalculatePower(data []string) int64 {
	gamma := make([]int, len(data[0]))
	epsilon := make([]int, len(data[0]))

	for _, byteData := range data {
		bit, _ := strconv.ParseInt(byteData, 2, 32)
		for i := 0; i < len(data[0]); i++ {
			bitValue := (bit >> i) & 1
			if bitValue == 1 {
				gamma[i]++
			}
		}
	}

	for i := 0; i < len(data[0]); i++ {
		zeroes := len(data) - gamma[i]
		if zeroes > gamma[i] {
			epsilon[i] = 1
			gamma[i] = 0
		} else {
			gamma[i] = 1

		}
	}

	gammaValue := float64(0)
	epsilonValue := float64(0)
	for i := 0; i < len(gamma); i++ {
		if gamma[i] == 1 {
			gammaValue += math.Pow(float64(2), float64(i))
		} else {
			epsilonValue += math.Pow(float64(2), float64(i))
		}
	}

	return int64(gammaValue * epsilonValue)
}

func CalculateLifeSupport(data []string) int64 {
	oxygen := calculateOxygenRating(data)
	co2 := calculateCO2Rating(data)
	return oxygen * co2
}

func filterData(data []string, position int, mostSignificantBit bool) []string {

	ones := 0
	zeroes := 0

	//Find most and least significant bit at position
	for _, byteData := range data {
		bit, _ := strconv.ParseInt(byteData, 2, 32)
		if (bit>>position)&1 == 1 {
			ones++
		} else {
			zeroes++
		}
	}

	filter := 0
	if mostSignificantBit {
		if ones >= zeroes {
			filter = 1
		} else {
			filter = 0
		}
	} else {
		if ones < zeroes {
			filter = 1
		} else {
			filter = 0
		}
	}

	//filter Data
	var dataFiltered []string
	for _, byteData := range data {
		bit, _ := strconv.ParseInt(byteData, 2, 32)
		bitValue := (bit >> position) & 1
		if filter == int(bitValue) {
			dataFiltered = append(dataFiltered, byteData)
		}
	}

	return dataFiltered
}

func calculateOxygenRating(data []string) int64 {

	for position := len(data[0]) - 1; position >= 0; position-- {
		filteredData := filterData(data, position, true)
		data = filteredData
		if len(data) == 1 {
			break
		}
	}

	oxygenRating, _ := strconv.ParseInt(data[0], 2, 32)
	return oxygenRating
}

func calculateCO2Rating(data []string) int64 {

	for position := len(data[0]) - 1; position >= 0; position-- {
		filteredData := filterData(data, position, false)
		data = filteredData
		if len(data) == 1 {
			break
		}
	}

	co2Rating, _ := strconv.ParseInt(data[0], 2, 32)
	return co2Rating
}
