package poker

import (
	"testing"
)

func areHandsEqual(a, b []Card) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestCardsEqual(t *testing.T) {
	card1 := NewCard(1, Heart)
	card2 := NewCard(1, Heart)

	if card1 != card2 {
		t.Error("Cards are not equal", card1, card2)
	} else {
		t.Log("Cards are equal")
	}
}

func TestHandsAreEqual(t *testing.T) {
	slice1 := []Card{NewCard(1, Heart)}
	slice2 := []Card{NewCard(1, Heart)}

	if !areHandsEqual(slice1, slice2) {
		t.Error("Hands are not equal", slice1, slice2)
	} else {
		t.Log("Hands are equal")
	}
}

func TestGetHighCard(t *testing.T) {
	input := []Card{
		NewCard(2, Heart), // assume cards are sorted in order first
		NewCard(5, Heart),
		NewCard(6, Heart),
		NewCard(7, Heart),
		NewCard(10, Heart),
		NewCard(11, Heart),
	}
	expectedOutput := []Card{
		NewCard(5, Heart),
		NewCard(6, Heart),
		NewCard(7, Heart),
		NewCard(10, Heart),
		NewCard(11, Heart),
	}
	highCards := getHighCards(input)
	if areHandsEqual(highCards, expectedOutput) {
		t.Log(highCards)
	} else {
		t.Errorf("High cards: Expected %s to equal %s", highCards, expectedOutput)
	}
}

func TestGetTwoOfAKind(t *testing.T) {
	input := []Card{
		NewCard(2, Heart), // assume cards are sorted in order first
		NewCard(5, Heart),
		NewCard(6, Heart),
		NewCard(6, Heart),
		NewCard(10, Heart),
		NewCard(11, Heart),
	}
	expectedOutput := []Card{
		NewCard(6, Heart),
		NewCard(6, Heart),
		NewCard(5, Heart),
		NewCard(10, Heart),
		NewCard(11, Heart),
	}

	groups := groupByNumbers(input)
	twoOfAKindHand, _ := getTwoOfAKind(groups)

	if areHandsEqual(twoOfAKindHand, expectedOutput) {
		t.Log(twoOfAKindHand)
	} else {
		t.Errorf("Two of a kind hand: Expected %s to equal %s", twoOfAKindHand, expectedOutput)
	}
}

func TestGetTwoPair(t *testing.T) {
	input := []Card{
		NewCard(2, Heart), // assume cards are sorted in order first
		NewCard(2, Heart),
		NewCard(6, Heart),
		NewCard(6, Heart),
		NewCard(10, Heart),
		NewCard(11, Heart),
	}
	expectedOutput := []Card{
		NewCard(2, Heart),
		NewCard(2, Heart),
		NewCard(6, Heart),
		NewCard(6, Heart),
		NewCard(11, Heart),
	}

	groups := groupByNumbers(input)
	twoPair, _ := getTwoPair(groups)

	if areHandsEqual(twoPair, expectedOutput) {
		t.Log(twoPair)
	} else {
		t.Errorf("Two pair hand: Expected %s to equal %s", twoPair, expectedOutput)
	}
}

func TestGetStraight(t *testing.T) {
	input := []Card{
		NewCard(1, Heart),
		NewCard(2, Heart),
		NewCard(3, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
	}
	straight, err := getStraight(input)
	if err == nil {
		t.Log(straight)
	} else {
		t.Error(err)
	}
}

func TestGetStraight7Cards(t *testing.T) {
	input := []Card{
		NewCard(1, Heart),
		NewCard(2, Heart),
		NewCard(3, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
		NewCard(6, Heart),
		NewCard(7, Heart),
	}
	straight, err := getStraight(input)
	if err == nil {
		t.Log(straight)
		lastCard := straight[len(straight)-1]
		if lastCard.Index != 7 {
			t.Errorf("Highest straight was not found: %s", straight)
		}

	} else {
		t.Error(err)
	}
}

func TestGetStraightWithInterruptedStraightHand(t *testing.T) {
	input := []Card{
		NewCard(1, Heart),
		NewCard(2, Heart),
		NewCard(2, Heart), // second "2" card interrupts the straight (cards are sorted in order first)
		NewCard(3, Heart),
		NewCard(3, Heart), // second "3" card interrupts the straight (cards are sorted in order first)
		NewCard(4, Heart),
		NewCard(5, Heart),
	}
	expectedHighestIndexOfStraight := 5
	straight, err := getStraight(input)
	if err == nil {
		t.Log(straight)
		lastCard := straight[len(straight)-1]
		if lastCard.Index != expectedHighestIndexOfStraight {
			t.Errorf("Highest straight was not found: %s", straight)
		}

	} else {
		t.Error(err)
	}
}

func TestGetStraightWithAceAsLow(t *testing.T) {
	input := []Card{
		NewCard(2, Heart),
		NewCard(3, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
		NewCard(14, Heart),
	}
	expectedOutput := []Card{
		NewCard(14, Heart),
		NewCard(2, Heart),
		NewCard(3, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
	}
	straight, err := getStraight(input)
	if err == nil {
		if areHandsEqual(straight, expectedOutput) {
			t.Log(straight)
		} else {
			t.Errorf("Expected %s to equal %s", straight, expectedOutput)
		}
	} else {
		t.Error(err)
	}
}

func TestGetFlush(t *testing.T) {
	input := []Card{
		NewCard(1, Heart),
		NewCard(2, Heart),
		NewCard(3, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
	}
	flush, err := getFlush(input)
	if err == nil {
		t.Log(flush)
	} else {
		t.Error(err)
	}
}

func TestGetFlushWithHighestCards(t *testing.T) {
	input := []Card{
		NewCard(1, Heart),
		NewCard(2, Heart),
		NewCard(3, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
		NewCard(6, Heart),
		NewCard(7, Heart),
	}
	flush, err := getFlush(input)
	if err == nil {
		if flush[len(flush)-1].Index != 7 {
			t.Errorf("Highest flush not found: %s", flush)
		} else {
			t.Log(flush)
		}
	} else {
		t.Error(err)
	}
}

func TestGetFullHouse(t *testing.T) {
	hand := []Card{
		NewCard(1, Club),
		NewCard(1, Heart),
		NewCard(4, Diamond),
		NewCard(4, Heart),
		NewCard(6, Diamond),
		NewCard(6, Heart),
		NewCard(6, Spade),
	}

	expectedOutput := []Card{
		NewCard(6, Diamond),
		NewCard(6, Heart),
		NewCard(6, Spade),
		NewCard(4, Diamond),
		NewCard(4, Heart),
	}

	groups := groupByNumbers(hand)
	fullHouse, err := getFullHouse(groups)

	if err == nil {
		if !areHandsEqual(fullHouse, expectedOutput) {
			t.Errorf("Expected %s to equal %s", fullHouse, expectedOutput)
		} else {
			t.Log(fullHouse)
		}
	} else {
		t.Error(err)
	}
}

func TestGetFourOfAKind(t *testing.T) {
	hand := []Card{
		NewCard(7, Club),
		NewCard(7, Diamond),
		NewCard(5, Diamond),
		NewCard(6, Diamond),
		NewCard(11, Club),
	}
	groups := groupByNumbers(hand)

	fourOfAKind, err := getFourOfAKind(groups)

	if err != nil {
		t.Logf("Expected error recieved: %s", err)
	} else {
		t.Errorf("getFourOfAKind: Expected top recieve an error. Instead got: %s", fourOfAKind)
	}
}

func TestGetStraightFlush(t *testing.T) {
	cards := []Card{
		NewCard(1, Heart),
		NewCard(2, Heart),
		NewCard(3, Heart),
		NewCard(4, Heart),
		NewCard(5, Heart),
		NewCard(6, Heart),
		NewCard(7, Heart),
	}
	straightFlush, err := getStraightFlush(cards)

	if err == nil {
		highestCardInStraightFlush := straightFlush[len(straightFlush)-1]
		if highestCardInStraightFlush.Index != 7 || highestCardInStraightFlush.Suit != Heart {
			t.Errorf("Highest straight flush not found: %s", straightFlush)
		} else {
			t.Log(straightFlush)
		}
	} else {
		t.Error(err)
	}
}

func TestGetBestHand(t *testing.T) {
	input := []Card{
		NewCard(7, Club),
		NewCard(7, Diamond),
		NewCard(5, Diamond),
		NewCard(6, Diamond),
		NewCard(11, Club),
	}
	expectedOutput := []Card{
		NewCard(7, Club),
		NewCard(7, Diamond),
		NewCard(5, Diamond),
		NewCard(6, Diamond),
		NewCard(11, Club),
	}
	output, _ := GetBestHand(input)

	if areHandsEqual(output, expectedOutput) {
		t.Log(output)
	} else {
		t.Errorf("GetBestHand: expected %s to equal %s", output, expectedOutput)
	}
}
