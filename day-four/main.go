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

	cardCount := make(map[int]int)

	var answer int

	line, _, err := reader.ReadLine()
	for err == nil {
		stringLine := string(line)
		cardLine := strings.Split(stringLine, ":")

		cardNumber, _ := strconv.Atoi(strings.TrimSpace(strings.Replace(cardLine[0], "Card", "", -1)))
		cardCount[cardNumber]++

		gameField := strings.Split(cardLine[1], "|")
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

		fmt.Println("cardNumber:", cardNumber)
		for i := 0; i < cardCount[cardNumber]; i++ {
			apply(cardCount, cardNumber, card.winNumbersAmount())
			answer++
		}
		line, _, err = reader.ReadLine()
	}

	fmt.Println("Answer:", answer)
}

func apply(cardCount map[int]int, currentCardNumber int, winNumbersAmount int) {
	adjustedAmount := winNumbersAmount
	for i := 1; i <= adjustedAmount; i++ {
		cardCount[currentCardNumber+i]++
	}
}

type card struct {
	actualNumbers  []int
	winningNumbers []int
}

func (c card) winNumbersAmount() int {
	var score int = 0
	for i := range c.actualNumbers {
		number := c.actualNumbers[i]
		if slices.Contains(c.winningNumbers, number) {
			score++
		}
	}
	return score
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
