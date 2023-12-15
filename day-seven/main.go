package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hands []Hand

func main() {
	hands := parseHands()
	sort.Sort(Hands(hands))
	var answer int
	for i, i2 := range hands {
		answer = answer + (i+1)*i2.bet
		fmt.Println("#", i+1, "*", i2.bet)
	}
	fmt.Println("Answer:", answer)
	//for _, hand := range hands {
	//	fmt.Println("Freq:", hand.frequencyTable().getFreqNumber())
	//}
}

type Hand struct {
	cards map[int]Card
	bet   int
}
type freqTable map[string]int

func (h Hand) frequencyTable() freqTable {
	frequencyTable := make(map[string]int)
	for _, card := range h.cards {
		frequencyTable[card.symbol]++
	}
	//fmt.Println("Cards:", frequencyTable)
	return frequencyTable
}

func (f freqTable) getFreqNumber() int {
	var maxFrequency, secondBiggest int
	var frequentestCard string
	for symbol, val := range f {
		//fmt.Println("index:", index, "val:", val)
		if val > maxFrequency {
			maxFrequency = val
			frequentestCard = symbol
		}
	}

	for symbol, val := range f {
		//fmt.Println("index:", index, "val:", val)
		if val > secondBiggest && symbol != frequentestCard {
			secondBiggest = val
		}
	}

	fmt.Println("Table:", f, "Max:", maxFrequency, "second:", secondBiggest)

	if maxFrequency > 3 {
		return maxFrequency
	} else if maxFrequency == 3 && secondBiggest == 2 {
		return maxFrequency
	} else if maxFrequency == 3 && secondBiggest != 2 {
		return 2
	}
	return maxFrequency
}

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Less(i, j int) bool {

	iMaxFreq := h[i].frequencyTable().getFreqNumber()
	jMaxFreq := h[j].frequencyTable().getFreqNumber()

	if iMaxFreq == jMaxFreq {
		if h[i].cards[0].power == h[j].cards[0].power {
			if h[i].cards[1].power == h[j].cards[1].power {
				if h[i].cards[2].power == h[j].cards[2].power {
					if h[i].cards[3].power == h[j].cards[3].power {
						return h[i].cards[4].power < h[j].cards[4].power
					} else {
						return h[i].cards[3].power < h[j].cards[3].power
					}
				} else {
					return h[i].cards[2].power < h[j].cards[2].power
				}
			} else {
				return h[i].cards[1].power < h[j].cards[1].power
			}
		} else {
			return h[i].cards[0].power < h[j].cards[0].power
		}
	} else {
		return iMaxFreq < jMaxFreq

	}
}

func (h Hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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

func parseHands() []Hand {
	var hands []Hand
	file, _ := os.Open("day-seven/input.txt")
	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	for err == nil {
		var result = make(map[int]Card)
		deal := strings.Split(string(line), " ")
		if len(deal) > 1 {
			bet, _ := strconv.Atoi(deal[1])
			for i := 0; i < len(deal[0]); i++ {
				result[i] = getCard(string(deal[0][i]))
			}
			hands = append(hands, Hand{cards: result, bet: bet})

			line, _, _ = reader.ReadLine()
		} else {
			break
		}
	}
	return hands

}
