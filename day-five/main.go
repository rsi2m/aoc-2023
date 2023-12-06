package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	soilConverter := readSoilConverter()

	fmt.Println("Answer:", soilConverter)
}

func readSoilConverter() []SoilConverter {
	file, _ := os.Open("./input.txt")
	reader := bufio.NewReader(file)
	var err error
	var line []byte
	start := false

	var output []SoilConverter

	var source, destination, rangeNumber int
	for err == nil {
		line, _, err = reader.ReadLine()
		stringLine := string(line)
		fmt.Println("Parsing:", stringLine)

		if stringLine == "" && start {
			break
		}

		if stringLine == "" {
			continue
		}

		if stringLine == "seed-to-soil map:" {
			start = true
			continue
		}

		if start {
			inputLines := strings.Split(stringLine, " ")
			source, _ = strconv.Atoi(inputLines[0])
			destination, _ = strconv.Atoi(inputLines[1])
			rangeNumber, _ = strconv.Atoi(inputLines[2])

			output = append(output, SoilConverter{
				source:      source,
				destination: destination,
				rangeNumber: rangeNumber,
			})
		}

	}
	return output

}

type SoilConverter struct {
	source      int
	destination int
	rangeNumber int
}
