package poker

import (
	"testing"
)

func TestGetWinningPlayers(t *testing.T) {

	steve := NewPlayer("Steve")
	steve.Hand = NewHand([]Card{
		NewCard(2, Club),
		NewCard(2, Heart),
		NewCard(3, Club),
		NewCard(5, Club),
		NewCard(6, Club),
	}, TWO_OF_A_KIND)
	matt := NewPlayer("Matt")
	matt.Hand = NewHand([]Card{
		NewCard(2, Spade),
		NewCard(2, Diamond),
		NewCard(4, Heart),
		NewCard(5, Heart),
		NewCard(6, Heart),
	}, TWO_OF_A_KIND)

	players := []Player{
		steve,
		matt,
	}

	winningPlayers := GetWinningPlayers(players)

	if len(winningPlayers) == 1 && winningPlayers[0].Name == "Matt" {
		t.Logf("Success: Matt has higher 3rd kicker 4 vs 3: %v", winningPlayers)
	} else {
		t.Errorf("Error: Matt was not returned: %v", winningPlayers)
	}
}

func TestGetPlayersWithHighestCard(t *testing.T) {

	steve := NewPlayer("Steve")
	steve.Hand = NewHand([]Card{
		NewCard(2, Club),
		NewCard(3, Club),
		NewCard(4, Club),
		NewCard(5, Club),
		NewCard(6, Club),
	}, STRAIGHT_FLUSH)
	matt := NewPlayer("Matt")
	matt.Hand = NewHand([]Card{
		NewCard(3, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
		NewCard(6, Heart),
		NewCard(7, Heart),
	}, STRAIGHT_FLUSH)

	players := []Player{
		steve,
		matt,
	}

	playersWithHighestCard := compareHighCard(players)

	if len(playersWithHighestCard) == 1 && playersWithHighestCard[0].Name == "Matt" {
		t.Log("Success: Matt has highest card")
	} else {
		t.Errorf("Error: Matt was not returned: %v", playersWithHighestCard)
	}
}

func TestCompareHighCardEqualHands(t *testing.T) {

	steve := NewPlayer("Steve")
	steve.Hand = NewHand([]Card{
		NewCard(2, Club),
		NewCard(3, Club),
		NewCard(4, Club),
		NewCard(5, Club),
		NewCard(6, Club),
	}, STRAIGHT_FLUSH)
	matt := NewPlayer("Matt")
	matt.Hand = NewHand([]Card{
		NewCard(2, Heart),
		NewCard(3, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
		NewCard(6, Heart),
	}, STRAIGHT_FLUSH)

	players := []Player{
		steve,
		matt,
	}

	playersWithHighestCard := compareHighCard(players)

	if len(playersWithHighestCard) == 2 {
		t.Log("Success: Players have same hand")
	} else {
		t.Errorf("Error: Both were not returned: %v", playersWithHighestCard)
	}
}

func TestCompareTwoOfAKind(t *testing.T) {

	steve := NewPlayer("Steve")
	steve.Hand = NewHand([]Card{
		NewCard(2, Club),
		NewCard(2, Club),
		NewCard(3, Club),
		NewCard(5, Club),
		NewCard(6, Club),
	}, TWO_OF_A_KIND)
	matt := NewPlayer("Matt")
	matt.Hand = NewHand([]Card{
		NewCard(2, Heart),
		NewCard(2, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
		NewCard(6, Heart),
	}, TWO_OF_A_KIND)

	players := []Player{
		steve,
		matt,
	}

	playersWithHighestCard := getCompareXOfAKindFunction(2)(players)

	if len(playersWithHighestCard) == 1 && playersWithHighestCard[0].Name == "Matt" {
		t.Logf("Success: Matt has a higher 3rd kicker: 4 vs 3: %v", playersWithHighestCard)
	} else {
		t.Errorf("Error: Matt was not returned: %v", playersWithHighestCard)
	}
}

func TestCompareTwoPair(t *testing.T) {

	steve := NewPlayer("Steve")
	steve.Hand = NewHand([]Card{
		NewCard(2, Club),
		NewCard(2, Club),
		NewCard(3, Club),
		NewCard(3, Club),
		NewCard(5, Club),
	}, STRAIGHT_FLUSH)
	matt := NewPlayer("Matt")
	matt.Hand = NewHand([]Card{
		NewCard(2, Heart),
		NewCard(2, Heart),
		NewCard(4, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
	}, STRAIGHT_FLUSH)

	players := []Player{
		steve,
		matt,
	}

	playersWithHighestCard := compareTwoPair(players)

	if len(playersWithHighestCard) == 1 && playersWithHighestCard[0].Name == "Matt" {
		t.Log("Success: Matt has pair of 4s, higher than 3s")
	} else {
		t.Errorf("Error: Matt was not returned: %v", playersWithHighestCard)
	}
}
