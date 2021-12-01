package main

import (
 "AdventOfCode/Day1"
 "fmt"
 "io"
 "log"
 "os"
)

func main() {
 day1()

}


func getValues(path string) []int {
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

func day1() {
 data := getValues("./Day1/data.txt")
 fmt.Println("Counter:", Day1.GetDepths(data))
 fmt.Println("Counter Window:", Day1.GetDepthsWindow(data))
}