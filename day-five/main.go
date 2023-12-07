package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	seeds := parseSeeds()
	fmt.Println("Seeds:", seeds)

	converters := []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

	// for _, converterName := range converters {
	// 	converterList := parseConverter(converterName)
	// 	for _, seed := range seeds {
	// 		result := seed
	// 		for _, converter := range converterList {
	// 			result = process(converter, seed)
	// 			if result != seed {
	// 				break
	// 			}

	// 		}
	// 		fmt.Println(converterName, ":", seed, "->", result)
	// 	}

	// }

	var answer int = 999999999999999999
	for _, seed := range seeds {
		result := seed
		for _, converterName := range converters {
			converterList := parseConverter(converterName)
			for _, converter := range converterList {
				newResult := process(converter, result)
				if newResult != result {
					result = newResult
					break
				}

			}
		}
		fmt.Println(seed, "->", result)
		if result < answer {
			answer = result
		}
	}
	fmt.Println("Answer:", answer)

}

func process(converter Converter, seed int) int {
	if seed >= converter.source && seed <= converter.source+converter.rangeNumber {
		return seed - converter.source + converter.destination
	} else {
		return seed
	}
}

func parseSeeds() []int {
	file, _ := os.Open("./input.txt")
	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()

	var seedNumbers []int

	for _, seedString := range strings.Split(strings.Replace(string(line), "seeds:", "", -1), " ") {
		seedNumber, _ := strconv.Atoi(strings.TrimSpace(seedString))
		if seedNumber > 0 {
			seedNumbers = append(seedNumbers, seedNumber)
		}
	}

	return seedNumbers
}

func parseConverter(converterName string) []Converter {
	file, _ := os.Open("./input.txt")
	reader := bufio.NewReader(file)
	var err error
	var line []byte
	start := false

	converterStartLine := fmt.Sprintf("-to-%s map:", converterName)

	var output []Converter

	var source, destination, rangeNumber int
	for err == nil {
		line, _, err = reader.ReadLine()
		stringLine := string(line)

		if stringLine == "" && start {
			break
		}

		if stringLine == "" {
			continue
		}

		if strings.Contains(stringLine, converterStartLine) {
			start = true
			continue
		}

		if start {
			inputLines := strings.Split(stringLine, " ")
			source, _ = strconv.Atoi(inputLines[1])
			destination, _ = strconv.Atoi(inputLines[0])
			rangeNumber, _ = strconv.Atoi(inputLines[2])

			output = append(output, Converter{
				source:      source,
				destination: destination,
				rangeNumber: rangeNumber,
			})
		}

	}
	return output

}

type Converter struct {
	source      int
	destination int
	rangeNumber int
}
