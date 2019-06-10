package poker

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit string

const (
	Heart   Suit = "♥"
	Spade   Suit = "♠"
	Diamond Suit = "♦"
	Club    Suit = "♣"
)

var suits = []Suit{Heart, Spade, Diamond, Club}

type Card struct {
	Index  int
	Number string
	Suit   Suit
}

func (c Card) String() string {
	num := fmt.Sprintf("%d", c.Index)
	if c.Index > 10 {
		num = c.Number
	}
	return fmt.Sprintf("{%s %s}", num, c.Suit)
}

var cardNumberMap = map[int]string{
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Jack",
	12: "Queen",
	13: "King",
	14: "Ace",
}

func NewCard(number int, suit Suit) Card {
	return Card{number, getCardNumber(number), suit}
}

func GetCardsOfSuit(suit Suit) []Card {
	const cardsInSuitCount = 13
	var cards []Card
	for i := 2; i <= cardsInSuitCount; i++ {
		cards = append(cards, NewCard(i, suit))
	}
	return cards
}

func ShuffleCards(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

func SortCards(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Index < cards[j].Index
	})
}

func getCardNumber(num int) string {
	return cardNumberMap[num]
}
