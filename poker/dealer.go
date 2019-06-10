package poker

import (
	"fmt"
)

func DealCards(deck *Deck, players []Player, cardsEach int) {
	for i := 0; i < cardsEach; i++ {
		for i := 0; i < len(players); i++ {
			players[i].TakeCard(deck.TakeCard())
		}
	}
}

func GetCommunityCards(deck *Deck, num int) []Card {
	var communityCards []Card
	for i := 0; i < num; i++ {
		communityCards = append(communityCards, deck.TakeCard())
	}
	return communityCards
}

func GetWinners(players []Player, communityCards []Card) []Player {
	fmt.Println("\nBest Hands:")
	for i := 0; i < len(players); i++ {
		bestHand, strength := GetBestHand(append(players[i].Cards, communityCards...))
		players[i].Hand = NewHand(bestHand, strength)
		fmt.Printf("%s: %s, %s\n", players[i].Name, bestHand, strength)
	}
	return GetWinningPlayers(players)
}
