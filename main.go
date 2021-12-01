package main

import (
 "AdventOfCode/Day1"
 "AdventOfCode/Utils"
 "fmt"
)

func main() {
 day1()
}




func day1() {
 data := Utils.ReadFileAsIntegers("./Day1/data.txt")
 fmt.Println("Counter:", Day1.GetDepths(data))
 fmt.Println("Counter Window:", Day1.GetDepthsWindow(data))
}