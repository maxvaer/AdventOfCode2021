package Day6

import (
	"AdventOfCode/Utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day6:")
	fish := getInitialFishFromFile("./Day6/data.txt")
	sortedFish := sortFish(fish)
	fmt.Println("Slow evolve count:", evolve(fish, 80))
	fmt.Println("Fast evolve count:", fastEvolve(sortedFish, 256))
	fmt.Println("----------")
}

func getInitialFishFromFile(path string) []int {
	var fishs []int
	data := Utils.ReadFileAsString(path)
	data = strings.Split(data[0], ",")
	for _, dataValue := range data {
		fishValue, _ := strconv.Atoi(dataValue)
		fish := fishValue
		fishs = append(fishs, fish)
	}

	return fishs
}

func evolve(fishPopulation []int, days int) int {
	for i := 1; i <= days; i++ {
		var newFish []int
		for i := 0; i < len(fishPopulation); i++ {
			fishPopulation[i]--
			if fishPopulation[i] == -1 {
				fishPopulation[i] = 6
				newFish = append(newFish, 8)
			}
		}
		fishPopulation = append(fishPopulation, newFish...)

	}

	return len(fishPopulation)
}

func fastEvolve(sortedFish []int, days int) int {
	for day := 0; day < days; day++ {
		tmpFish := sortedFish[0]
		for i := 1; i <= 8; i++ {
			sortedFish[i-1] = sortedFish[i]
		}
		sortedFish[8] = tmpFish
		sortedFish[6] += tmpFish
	}

	count := 0

	for _, fishValue := range sortedFish {
		count += fishValue
	}

	return count
}

func sortFish(fishPopulation []int) []int {
	var sortedFishPopulation []int
	for i := 0; i <= 8; i++ {
		count := 0
		for _, fish := range fishPopulation {
			if fish == i {
				count++
			}
		}

		sortedFishPopulation = append(sortedFishPopulation, count)
	}

	return sortedFishPopulation
}
