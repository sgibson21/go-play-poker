package main

import (
	"fmt"
	"go-play-poker/poker"
)

func main() {

	deck := poker.GetNewDeck()
	players := poker.GetPlayers()

	poker.DealCards(&deck, players, 2)

	fmt.Println()
	poker.PrintPlayers(players)

	communityCards := poker.GetCommunityCards(&deck, 5)
	fmt.Printf("\nCommunity Cards: %s\n", communityCards)

	winners := poker.GetWinners(players, communityCards)

	fmt.Println("Winners:")
	for _, winner := range winners {
		fmt.Printf("%s %s %s\n", winner.Name, winner.Hand.Cards, winner.Hand.Strength)
	}

}
