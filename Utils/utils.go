package Utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFileAsIntegers(path string) []int {
	var perline int
	var nums []int

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}
	return nums
}