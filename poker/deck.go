package poker

import (
	"fmt"
)

type Deck struct {
	cards []Card
}

func GetNewDeck() Deck {
	var cards []Card
	cards = append(cards, GetCardsOfSuit(Heart)...)
	cards = append(cards, GetCardsOfSuit(Spade)...)
	cards = append(cards, GetCardsOfSuit(Diamond)...)
	cards = append(cards, GetCardsOfSuit(Club)...)
	cards = ShuffleCards(cards)
	return Deck{cards}
}

func (d Deck) GetCards() []Card {
	return d.cards
}

func (d Deck) String() string {
	return fmt.Sprintf("%v", d.cards)
}

func (d Deck) Shuffle() Deck {
	d.cards = ShuffleCards(d.cards)
	return d
}

func (d *Deck) TakeCard() Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}
