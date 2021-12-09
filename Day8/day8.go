package Day8

import (
	"AdventOfCode/Utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day8:")
	signalPattern, digitOutput := GetSegmentData("./Day8/data.txt")
	fmt.Println("Simple Digits:", countSimpleDigits(digitOutput))
	fmt.Println("Count:", calculateOutput(signalPattern, digitOutput))
	fmt.Println("----------")
}

func calculateOutput(signalPatterns []string, digitOutputs []string) uint {
	count := uint(0)
	for index, pattern := range signalPatterns {
		digits := strings.Split(digitOutputs[index], " ")
		sortedChars := identifyChars(pattern)
		digitResult := ""
		for i := 0; i < len(digits); i++ {
			digit := sortSignal(digits[i])
			value := getSignalValue(sortedChars, digit)
			dst := []byte(digitResult)
			digitResult = string(strconv.AppendInt(dst, int64(value), 10))

		}
		tmp, _ := strconv.Atoi(digitResult)
		count += uint(tmp)
	}

	return count
}

func getSignalValue(identifiedChars []string, signal string) uint {
	result := ^uint(0)
	for index, char := range identifiedChars {
		if signal == char {
			result = uint(index)
			break
		}
	}

	return result
}

func GetSegmentData(path string) ([]string, []string) {
	var signalPattern []string
	var digitOutput []string

	data := Utils.ReadFileAsString(path)

	for _, stringValue := range data {
		substrings := strings.Split(stringValue, " | ")

		signalPattern = append(signalPattern, substrings[0])
		digitOutput = append(digitOutput, substrings[1])
	}

	return signalPattern, digitOutput
}

func countSimpleDigits(digitOutput []string) uint {
	countOf1 := uint(0)
	countOf4 := uint(0)
	countOf7 := uint(0)
	countOf8 := uint(0)

	for _, digitOutputValue := range digitOutput {
		digits := strings.Split(digitOutputValue, " ")
		for _, digit := range digits {
			digitLength := len(digit)
			if digitLength == 2 {
				countOf1++
			} else if digitLength == 4 {
				countOf4++
			} else if digitLength == 3 {
				countOf7++
			} else if digitLength == 7 {
				countOf8++
			}
		}
	}

	return countOf1 + countOf4 + countOf7 + countOf8
}

func identifyChars(signalPattern string) []string {

	s1, s4, s7, s8 := identify1478(signalPattern)
	s9, bottom := getS9(signalPattern, s4, s7)
	bottomLeft, _ := diff(s8, s9)
	s0, mid := getS0(signalPattern, s8, s9, s1)
	chars, _ := diff(s7, s1)
	top := string(chars[0])
	s3 := getS3(s1, top, mid, bottom)
	s2, _ := getS2(signalPattern, top, mid, bottom, bottomLeft)
	topLeft, _ := diff(s3, s9)
	s5, _ := getS5(signalPattern, top, mid, bottom, topLeft)
	s6 := merge(s5, bottomLeft)

	return []string{sortSignal(s0), sortSignal(s1), sortSignal(s2), sortSignal(s3), sortSignal(s4), sortSignal(s5),
		sortSignal(s6), sortSignal(s7), sortSignal(s8), sortSignal(s9)}
}

func getS9(signalPattern string, s4 string, s7 string) (string, string) {
	signals := strings.Split(signalPattern, " ")

	tmpS9 := merge(s4, s7)
	char := ""
	for _, signal := range signals {
		diffChar, length := diff(tmpS9, signal)
		if length == 1 && len(signal) == 6 {
			char = diffChar
			break
		}
	}

	return tmpS9 + char, char
}

func getS5(signalPattern string, topBar string, midBar string, bottom string, topLeft string) (string, string) {
	signals := strings.Split(signalPattern, " ")

	tmp := merge(topBar, midBar)
	tmp = merge(tmp, bottom)
	tmp = merge(tmp, topLeft)

	s5 := ""
	bottomRight := ""

	for _, signal := range signals {
		diffChar, length := diff(tmp, signal)
		if length == 1 {
			bottomRight = diffChar
			s5 = signal
			break
		}
	}

	return s5, bottomRight
}

func getS3(s1 string, topBar string, midBar string, bottomBar string) string {
	s3 := merge(s1, topBar)
	s3 = merge(s3, midBar)
	s3 = merge(s3, bottomBar)

	return s3
}

func getS2(signalPattern string, topBar string, midBar string, bottomBar string, bottomLeft string) (string, string) {
	signals := strings.Split(signalPattern, " ")

	tmp := merge(topBar, midBar)
	tmp = merge(tmp, bottomBar)
	tmp = merge(tmp, bottomLeft)

	s2 := ""
	topRight := ""

	for _, signal := range signals {
		diffChar, length := diff(tmp, signal)
		if length == 1 {
			topRight = diffChar
			s2 = signal
			break
		}
	}

	return s2, topRight
}

func getS0(signalPattern string, s8 string, s9 string, s1 string) (string, string) {
	signals := strings.Split(signalPattern, " ")

	midBar := ""
	s0 := ""
	for _, signal := range signals {
		diffChar, length := diff(s8, signal)
		if length == 1 && !signalEqual(sortSignal(signal), sortSignal(s9)) && signalIncludes1(s1, signal) {
			midBar = diffChar
			s0 = signal
			break
		}
	}

	return s0, midBar
}

func signalEqual(signal string, otherSignal string) bool {
	_, length := diff(signal, otherSignal)
	return length == 0
}

func identify1478(signalPattern string) (string, string, string, string) {
	string1 := ""
	string4 := ""
	string7 := ""
	string8 := ""

	digits := strings.Split(signalPattern, " ")
	for _, digit := range digits {
		digitLength := len(digit)
		if digitLength == 2 {
			string1 = digit
		} else if digitLength == 4 {
			string4 = digit
		} else if digitLength == 3 {
			string7 = digit
		} else if digitLength == 7 {
			string8 = digit
		}
	}

	return string1, string4, string7, string8
}

func merge(signal string, otherSignal string) string {
	sigArray := strings.Split(signal, "")
	otherSigArray := strings.Split(otherSignal, "")
	var mergeChars []string

	for _, s := range sigArray {
		if !Utils.StringArrayContains(mergeChars, s) {
			mergeChars = append(mergeChars, s)
		}
	}

	for _, s := range otherSigArray {
		if !Utils.StringArrayContains(mergeChars, s) {
			mergeChars = append(mergeChars, s)
		}
	}

	return strings.Join(mergeChars, "")
}

func diff(signal string, otherSignal string) (string, uint) {
	sigArray := strings.Split(signal, "")
	otherSigArray := strings.Split(otherSignal, "")
	var diffChars []string

	for i := 0; i < len(sigArray); i++ {
		sigChar := sigArray[i]
		for j := 0; j < len(otherSigArray); j++ {
			otherSigChar := otherSigArray[j]
			if sigChar == otherSigChar {
				sigArray[i] = ""
				otherSigArray[j] = ""
			}
		}
	}

	for _, s := range sigArray {
		if s != "" {
			diffChars = append(diffChars, s)
		}
	}
	for _, s := range otherSigArray {
		if s != "" {
			diffChars = append(diffChars, s)
		}
	}

	return strings.Join(diffChars, ""), uint(len(diffChars))
}

func sortSignal(signal string) string {
	s := strings.Split(signal, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func signalIncludes1(s1 string, signal string) bool {
	s1Split := strings.Split(s1, "")
	signalSplit := strings.Split(signal, "")
	result := true
	for _, s := range s1Split {
		includes := Utils.StringArrayContains(signalSplit, s)
		if !includes {
			result = false
		}
	}
	return result
}
