package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var answer = 1
	race := parseRaces()
	waysToBeatRecord := checkRace(race)
	fmt.Println("Race:", race, " Ways to beat:", waysToBeatRecord)
	answer = answer * waysToBeatRecord
	fmt.Println("Answer:", answer)
}

func checkRace(race Race) int {
	var waysToBeatRecord int
	for i := 1; i < race.time; i++ {
		totalDistance := (race.time - i) * i
		if totalDistance > race.distance {
			waysToBeatRecord++
		}
	}
	return waysToBeatRecord
}

func parseRaces() Race {
	file, _ := os.Open("day-six/input.txt")
	reader := bufio.NewReader(file)
	var line []byte
	line, _, _ = reader.ReadLine()
	timeNumber, _ := strconv.Atoi(string(tokenize(line)))
	fmt.Println("ParsedNumber:", timeNumber)

	line, _, _ = reader.ReadLine()
	distanceNumber, _ := strconv.Atoi(tokenize(line))

	race := Race{
		time: timeNumber, distance: distanceNumber,
	}
	fmt.Println("ParsedNumber:", distanceNumber)
	fmt.Println("Parsed:", race)
	return race
}

func tokenize(inputString []byte) string {
	stringLine := string(inputString)
	noExtra := strings.ReplaceAll(strings.ReplaceAll(stringLine, "Time:", ""), "Distance:", "")
	return strings.ReplaceAll(noExtra, " ", "")
}

type Race struct {
	time     int
	distance int
}
