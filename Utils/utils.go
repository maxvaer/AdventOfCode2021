package Utils

import (
	"bufio"
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

func ReadFileAsString(path string) []string {
	var data []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}
