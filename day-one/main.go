package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()
	var sum int

	digitSpelling := make(map[int]string)

	digitSpelling[0] = "zero"
	digitSpelling[1] = "one"
	digitSpelling[2] = "two"
	digitSpelling[3] = "three"
	digitSpelling[4] = "four"
	digitSpelling[5] = "five"
	digitSpelling[6] = "six"
	digitSpelling[7] = "seven"
	digitSpelling[8] = "eight"
	digitSpelling[9] = "nine"

	for err == nil {
		result := getCalibrationValue(string(line), digitSpelling)
		sum += result
		fmt.Println("Total:", sum, ", added:", result)
		line, _, err = reader.ReadLine()
	}
	fmt.Println("Total:", sum)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCalibrationValue(inputString string, digitSpelling map[int]string) int {
	var firstDigit, lastDigit int
	var firstDigitIndex = 9999999
	var lastDigitIndex = 0

	for key, val := range digitSpelling {
		if foundIndexStart := strings.Index(inputString, val); foundIndexStart != -1 && foundIndexStart < firstDigitIndex {
			firstDigit, firstDigitIndex = key, foundIndexStart
		}
		if foundIndexEnd := strings.LastIndex(inputString, val); foundIndexEnd != -1 && foundIndexEnd > lastDigitIndex {
			lastDigit, lastDigitIndex = key, foundIndexEnd
		}
	}

	var firstDigitFound, lastDigitFound bool
	for i, j := 0, len(inputString)-1; i < len(inputString); i, j = i+1, j-1 {
		if _, err := strconv.Atoi(string(inputString[i])); err == nil && i <= firstDigitIndex && !firstDigitFound {
			firstDigit, _ = strconv.Atoi(string(inputString[i]))
			firstDigitFound = true
		}
		if _, err := strconv.Atoi(string(inputString[j])); err == nil && j >= lastDigitIndex && !lastDigitFound {
			lastDigit, _ = strconv.Atoi(string(inputString[j]))
			lastDigitFound = true
		}
	}

	calibrationString := string(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))
	fmt.Println("Input:", inputString, "Result:", calibrationString)
	output, _ := strconv.Atoi(calibrationString)
	return output
}
