package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	reader := bufio.NewReader(file)

	var answer int

	line, _, err := reader.ReadLine()
	for err == nil {
		stringLine := string(line)
		gameField := strings.Split(strings.Split(stringLine, ":")[1], "|")
		winningNumberField := strings.Split(strings.TrimSpace(gameField[0]), " ")
		actualNumberField := strings.Split(strings.TrimSpace(gameField[1]), " ")

		var winningNumbers []int
		for i := range winningNumberField {
			parsedNumber, _ := strconv.Atoi(strings.TrimSpace(string(winningNumberField[i])))
			if parsedNumber > 0 {
				winningNumbers = append(winningNumbers, parsedNumber)
			}
		}
		var actualNumbers []int
		for i := range actualNumberField {
			parsedNumber, _ := strconv.Atoi(strings.TrimSpace(string(actualNumberField[i])))
			if parsedNumber > 0 {
				actualNumbers = append(actualNumbers, parsedNumber)
			}
		}
		card := card{
			actualNumbers:  actualNumbers,
			winningNumbers: winningNumbers,
		}
		fmt.Println("Card:", card, " Score:", card.calculateScore())
		answer += card.calculateScore()
		line, _, err = reader.ReadLine()
	}

	fmt.Println("Anser:", answer)
}

type card struct {
	actualNumbers  []int
	winningNumbers []int
}

func (c card) calculateScore() int {
	var score int = 0
	for i := range c.actualNumbers {
		number := c.actualNumbers[i]
		if slices.Contains(c.winningNumbers, number) {
			if score == 0 {
				score = 1
			} else {
				score = score * 2
			}
		}
	}
	return score
}
