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

	for err == nil {
		fmt.Println(readGame(string(line)))
		line, _, err = reader.ReadLine()
	}

}

func readGame(line string) game {
	gameId, _ := strconv.Atoi(strings.Replace(strings.Fields(line)[1], ":", "", 1))
	gameRecord := strings.Split(line, ":")[1]
	fmt.Println(gameRecord)
	return game{id: gameId}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type game struct {
	id int
	// gameRecord []gameSet
}

type gameSet struct {
	observations []observation
}

type observation struct {
	amount    int
	diceColor string
}
