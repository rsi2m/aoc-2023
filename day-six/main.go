package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var answer = 1
	races := parseRaces()
	for _, race := range races {
		waysToBeatRecord := checkRace(race)
		fmt.Println("Race:", race, " Ways to beat:", waysToBeatRecord)
		answer = answer * waysToBeatRecord
	}
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

func parseRaces() []Race {
	file, _ := os.Open("day-six/input.txt")
	reader := bufio.NewReader(file)
	var line []byte
	line, _, _ = reader.ReadLine()
	var races []Race
	for _, entry := range tokenize(line) {
		timeNumber, _ := strconv.Atoi(string(entry))
		fmt.Println("ParsedNumber:", timeNumber)
		races = append(races, Race{time: timeNumber})
	}

	line, _, _ = reader.ReadLine()
	for index, entry := range tokenize(line) {
		distanceNumber, _ := strconv.Atoi(string(entry))
		fmt.Println("ParsedNumber:", distanceNumber)
		races[index].distance = distanceNumber
	}
	fmt.Println("Parsed:", races)
	return races
}

func tokenize(inputString []byte) []string {
	stringLine := string(inputString)
	noExtra := strings.ReplaceAll(strings.ReplaceAll(stringLine, "Time:", ""), "Distance:", "")
	space := regexp.MustCompile(`\s+`)
	return strings.Split(strings.TrimSpace(space.ReplaceAllString(noExtra, " ")), " ")
}

type Race struct {
	time     int
	distance int
}
