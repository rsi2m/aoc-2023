package day_seven

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	hands := parseHands()
}

type Hand struct {
	cards map[int]Card
}

type Card struct {
	symbol string
	power  int
}

func getCard(symbol string) Card {
	switch symbol {
	case "A":
		return Card{power: 14, symbol: symbol}
	case "K":
		return Card{power: 13, symbol: symbol}
	case "Q":
		return Card{power: 12, symbol: symbol}
	case "J":
		return Card{power: 11, symbol: symbol}
	case "T":
		return Card{power: 10, symbol: symbol}
	default:
		power, _ := strconv.Atoi(symbol)
		return Card{power: power, symbol: symbol}
	}
}

func parseHands() map[int]Hand {
	file, _ := os.Open("day-six/input.txt")
	reader := bufio.NewReader(file)
	var line []byte
	line, _, _ = reader.ReadLine()

}
