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

	legalMap := make(map[string]int)

	legalMap["red"] = 12
	legalMap["green"] = 13
	legalMap["blue"] = 14

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()

	var sum int
	for err == nil {
		game := readGame(string(line))
		fmt.Println("Game:", game.id, "Legal:", game.isLegal(legalMap))
		sum += game.getMinSet()
		line, _, err = reader.ReadLine()
	}
	fmt.Println("Answer:", sum)

}

// 12 red cubes, 13 green cubes, and 14 blue cubes

func readGame(line string) game {
	gameId, _ := strconv.Atoi(strings.Replace(strings.Fields(line)[1], ":", "", 1))
	gameRecord := parseGameRecord(strings.Split(line, ":")[1])
	return game{
		id:         gameId,
		gameRecord: gameRecord,
	}
}

func parseGameRecord(gameRecordString string) []gameSet {
	var parsedGameRecord []gameSet
	gameSetArray := strings.Split(gameRecordString, ";")
	for i := 0; i < len(gameSetArray); i++ {
		gameSetUnparsed := strings.Split(strings.TrimSpace(gameSetArray[i]), ",")
		parsedGameRecord = append(parsedGameRecord, gameSet{parseGameSet(gameSetUnparsed)})
	}
	return parsedGameRecord
}

func parseGameSet(gameSetUnparsed []string) []observation {
	var parsedObservations []observation
	for j := 0; j < len(gameSetUnparsed); j++ {
		observationUnparsed := strings.Split(strings.TrimSpace(gameSetUnparsed[j]), " ")
		amount, _ := strconv.Atoi(observationUnparsed[0])
		observation := observation{
			amount:    amount,
			diceColor: observationUnparsed[1],
		}
		parsedObservations = append(parsedObservations, observation)
	}
	return parsedObservations
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type game struct {
	id         int
	gameRecord []gameSet
}

func (g game) isLegal(legalMap map[string]int) bool {
	isLegal := true
	for i := range g.gameRecord {
		if !isLegal {
			break
		}
		for j := range g.gameRecord[i].observations {
			observation := g.gameRecord[i].observations[j]
			if observation.amount > legalMap[observation.diceColor] {
				isLegal = false
				break
			}
		}
	}
	return isLegal
}

func (g game) getMinSet() int {
	minimumMap := make(map[string]int)

	minimumMap["red"] = 0
	minimumMap["green"] = 0
	minimumMap["blue"] = 0

	for i := range g.gameRecord {
		for j := range g.gameRecord[i].observations {
			observation := g.gameRecord[i].observations[j]
			if minimumMap[observation.diceColor] < observation.amount {
				minimumMap[observation.diceColor] = observation.amount
			}
		}
	}

	power := 1
	for _, val := range minimumMap {
		power = power * val
	}

	return power
}

type gameSet struct {
	observations []observation
}

type observation struct {
	amount    int
	diceColor string
}
