package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	legalMap := make(map[string]int)

	legalMap["red"] = 12
	legalMap["green"] = 13
	legalMap["blue"] = 14

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()
	matrix := matrix{[]string{}, [][]PartNumber{}}
	for err == nil {
		stringLine := string(line)
		matrix.data = append(matrix.data, stringLine)
		matrix.partNumbers = append(matrix.partNumbers, readPartNumbers(stringLine))
		line, _, err = reader.ReadLine()
	}

	var answer int
	for x := range matrix.data {
		for y := range matrix.data[x] {
			if matrix.isSymbol(x, y) {
				if matrix.get(x, y) == "*" {
					adjacentPartNumbers := matrix.getAdjacentPartNumbers(x, y)
					if len(adjacentPartNumbers) == 2 {
						first, _ := strconv.Atoi(adjacentPartNumbers[0].Number)
						second, _ := strconv.Atoi(adjacentPartNumbers[1].Number)
						fmt.Println("Engine at x:", x, " y:", y, "Ratio:", first*second)
						answer += first * second
					}
				}

			}
		}
	}

	fmt.Println("Anser:", answer)
}

type matrix struct {
	data        []string
	partNumbers [][]PartNumber
}

func (m matrix) get(x int, y int) string {
	if x < 0 || y < 0 || x >= len(m.data) || y >= len(m.data[0]) {
		return ""
	}
	return string(m.data[x][y])
}

func (m matrix) isSymbol(x int, y int) bool {
	value := m.get(x, y)
	return value != "." && !unicode.IsDigit([]rune(value)[0])
}

func (m matrix) getAdjacentPartNumbers(x int, y int) []PartNumber {
	partMap := make(map[PartNumber]bool)

	left := m.getPartNumber(x, y-1)
	putIntoMapIfOK(left, partMap)

	right := m.getPartNumber(x, y+1)
	putIntoMapIfOK(right, partMap)

	top := m.getPartNumber(x-1, y)
	putIntoMapIfOK(top, partMap)

	bottom := m.getPartNumber(x+1, y)
	putIntoMapIfOK(bottom, partMap)

	leftTop := m.getPartNumber(x-1, y-1)
	putIntoMapIfOK(leftTop, partMap)

	leftBot := m.getPartNumber(x+1, y-1)
	putIntoMapIfOK(leftBot, partMap)

	rightTop := m.getPartNumber(x-1, y+1)
	putIntoMapIfOK(rightTop, partMap)

	rightBot := m.getPartNumber(x+1, y+1)
	putIntoMapIfOK(rightBot, partMap)

	fmt.Println("Part Numbers:", Keys(partMap))
	return Keys(partMap)
}

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func putIntoMapIfOK(p PartNumber, partMap map[PartNumber]bool) {
	if p.Number != "" {
		partMap[p] = true
	}
}

func (m matrix) getPartNumber(x int, y int) PartNumber {
	if x < 0 || y < 0 || x >= len(m.data) || y >= len(m.data[0]) {
		return PartNumber{}
	}

	for partIndex := range m.partNumbers[x] {
		if y >= m.partNumbers[x][partIndex].start && y <= m.partNumbers[x][partIndex].end {
			return m.partNumbers[x][partIndex]
		}
	}
	return PartNumber{}
}

type PartNumber struct {
	Number string
	start  int
	end    int
}

func readPartNumbers(stringLine string) []PartNumber {
	var result []PartNumber
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringIndex(stringLine, -1)
	for _, match := range matches {
		number := stringLine[match[0]:match[1]]
		result = append(result, PartNumber{
			Number: number,
			start:  match[0],
			end:    match[1] - 1,
		})
	}

	return result
}
