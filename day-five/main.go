package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {

	seeds := parseSeeds()
	fmt.Println("Seeds:", seeds)

	converterNames := []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	var converters [][]Converter
	for _, converterName := range converterNames {
		converters = append(converters, parseConverter(converterName))
	}

	chanRes := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(len(seeds))
	// for i := 0; i < len(seeds); i++ {
	// 	go execute(chanRes, i, &wg)
	// }

	var x []int

	fmt.Println("Return:", x)

	for _, seed := range seeds {
		go executeSeed(chanRes, seed, converters, &wg)
	}

	for {
		oneResult := <-chanRes
		x = append(x, oneResult)
		if len(x) == len(seeds) {
			break
		}
	}

	fmt.Println("Answer:", slices.Min(x))

}

func executeSeed(out chan<- int, seed Seed, converters [][]Converter, wg *sync.WaitGroup) {
	fmt.Println("Doin:", seed)
	defer wg.Done()
	var results []int
	for seedNumber := seed.start; seedNumber < seed.start+seed.rangeNumber; seedNumber++ {
		results = append(results, process(seedNumber, converters))
	}
	fmt.Println(seed.start, "->", slices.Min(results))
	out <- slices.Min(results)
}

func execute(out chan<- int, seed int, wg *sync.WaitGroup) {
	fmt.Println("Doin:", seed)
	defer wg.Done()
	out <- seed + 10
}

func process(seedNumber int, converters [][]Converter) int {
	result := seedNumber
	for _, converterList := range converters {
		for _, converter := range converterList {
			newResult := convert(converter, result)
			if newResult != result {
				result = newResult
				break
			}

		}
	}
	return result
}

func convert(converter Converter, seed int) int {
	if seed >= converter.source && seed <= converter.source+converter.rangeNumber {
		return seed - converter.source + converter.destination
	} else {
		return seed
	}
}

func parseSeeds() []Seed {
	file, _ := os.Open("./input.txt")
	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()

	var seedNumbers []Seed
	var start int

	for index, seedString := range strings.Split(strings.Replace(string(line), "seeds:", "", -1), " ") {
		seedNumber, _ := strconv.Atoi(strings.TrimSpace(seedString))

		if seedNumber == 0 {
			continue
		}

		if index%2 != 0 {
			start = seedNumber
		} else {
			seedNumbers = append(seedNumbers, Seed{start: start, rangeNumber: seedNumber})
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

type Seed struct {
	start       int
	rangeNumber int
}

type Converter struct {
	source      int
	destination int
	rangeNumber int
}
