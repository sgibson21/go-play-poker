package poker

import (
	"errors"
	"fmt"
	"strconv"
)

var handLength = 5

// Sort once in GetBestHand method then when appending to new array, append (prepend) to the start of the array to maintain order

/**
 * Gets the best hand within a group of cards
 * @returns []Card - the best hand, strength - the strength of the hand eg:
 * 		Straight Flush  = 9
 * 		Four of a Kind  = 8
 * 		Full House      = 7
 * 		Flush           = 6
 * 		Straight        = 5
 * 		Three of a Kind = 4
 * 		Two Pair        = 3
 * 		Two of a Kind   = 2
 * 		High Card       = 1
 */
func GetBestHand(hand []Card) ([]Card, strength) {

	if len(hand) < 5 {
		panic(fmt.Sprintf("Not enough cards. Need %d got: %s ", handLength, hand))
	}

	SortCards(hand)
	groups := groupByNumbers(hand)

	bestHand, err := getStraightFlush(hand)

	if err == nil {
		return bestHand, STRAIGHT_FLUSH
	} else {
		bestHand, err = getFourOfAKind(groups)
	}

	if err == nil {
		return bestHand, FOUR_OF_A_KIND
	} else {
		bestHand, err = getFullHouse(groups)
	}

	if err == nil {
		return bestHand, FULL_HOUSE
	} else {
		bestHand, err = getFlush(hand)
	}

	if err == nil {
		return bestHand, FLUSH
	} else {
		bestHand, err = getStraight(hand)
	}

	if err == nil {
		return bestHand, STRAIGHT
	} else {
		bestHand, err = getThreeOfAKind(groups)
	}

	if err == nil {
		return bestHand, THREE_OF_A_KIND
	} else {
		bestHand, err = getTwoPair(groups)
	}

	if err == nil {
		return bestHand, TWO_PAIR
	} else {
		bestHand, err = getTwoOfAKind(groups)
	}

	if err == nil {
		return bestHand, TWO_OF_A_KIND
	} else {
		bestHand = getHighCards(hand)
	}

	return bestHand, HIGH_CARD

}

func getHighCards(hand []Card) []Card {
	length := len(hand)
	outputLength := handLength
	return hand[length-outputLength:]
}

func getTwoOfAKind(groups map[int][]Card) ([]Card, error) {
	return getXOfAKind(groups, 2)
}

func getTwoPair(groups map[int][]Card) ([]Card, error) {
	var remainder, cardsToReturn []Card
	for _, cards := range groups {
		if len(cards) == 2 {
			cardsToReturn = append(cardsToReturn, cards...)
		} else {
			remainder = append(remainder, cards...)
		}
	}
	if len(cardsToReturn) == 4 {
		SortCards(cardsToReturn)
		return append(cardsToReturn, remainder[len(remainder)-1]), nil
	} else {
		return nil, errors.New("Two pair hand not found")
	}
}

func getThreeOfAKind(groups map[int][]Card) ([]Card, error) {
	return getXOfAKind(groups, 3)
}

/*
 * Start with the highest card and check if any 5 cards in a row are sequential.
 * This ensures you get the highest straight available from a set of cards whose length is more than 5.
 */
func getStraight(hand []Card) ([]Card, error) {
	length := len(hand)
	lastCardInHand := hand[length-1]
	lastIndex := lastCardInHand.Index
	straight := []Card{lastCardInHand} // start with last element

	// start at the penultimate element and work your way back to the start
	for i := length - 2; i >= 0; i-- {
		if lastIndex-1 == hand[i].Index {
			straight = append([]Card{hand[i]}, straight...) // prepend the cards to the hand to keep in order
			if len(straight) == handLength {
				return straight, nil
			}
		} else if lastIndex != hand[i].Index {
			// reset the straight slice to start at this card if the index is not equal to, or one below the last index
			straight = []Card{hand[i]}
		}
		lastIndex = hand[i].Index
	}

	// check for low ace
	if len(straight) == handLength-1 && straight[0].Index == 2 && hand[length-1].Index == 14 {
		return append([]Card{hand[length-1]}, straight...), nil
	}

	return nil, errors.New("Straight hand not found")
}

func getFlush(hand []Card) ([]Card, error) {
	var hearts []Card
	var spades []Card
	var diamonds []Card
	var clubs []Card
	for i := len(hand) - 1; i >= 0; i-- {
		if hand[i].Suit == Heart {
			hearts = append([]Card{hand[i]}, hearts...)
			if len(hearts) == handLength {
				return hearts, nil
			}
		} else if hand[i].Suit == Spade {
			spades = append([]Card{hand[i]}, spades...)
			if len(spades) == handLength {
				return spades, nil
			}
		} else if hand[i].Suit == Diamond {
			diamonds = append([]Card{hand[i]}, diamonds...)
			if len(diamonds) == handLength {
				return diamonds, nil
			}
		} else if hand[i].Suit == Club {
			clubs = append([]Card{hand[i]}, clubs...)
			if len(clubs) == handLength {
				return clubs, nil
			}
		}
	}
	return nil, errors.New("Flush hand not found")
}

func getFullHouse(groups map[int][]Card) ([]Card, error) {
	threeOfAKind, err := getXOfAKind(groups, 3)

	if err == nil {
		threeOfAKind = threeOfAKind[:3] // the three of a kind will be the first 3 cards
		twoOfAKind, err := getXOfAKind(groups, 2)

		if err == nil {
			twoOfAKind = twoOfAKind[:2] // the two of a kind will be the first 2 cards
			return append(threeOfAKind, twoOfAKind...), nil
		}
	}

	return nil, errors.New("Full house hand not found")
}

func getFourOfAKind(groups map[int][]Card) ([]Card, error) {
	return getXOfAKind(groups, 4)
}

func getStraightFlush(hand []Card) ([]Card, error) {
	groups := groupBySuit(hand)
	var straightFlushes [][]Card

	for _, group := range groups {
		if len(group) >= handLength {
			straight, err := getStraight(group)

			if err == nil {
				straightFlushes = append(straightFlushes, straight)
			}
		}
	}

	if len(straightFlushes) == 0 {
		return nil, errors.New("Straight flush hand not found")
	}

	highestStraightFlush := straightFlushes[0]

	if len(straightFlushes) > 1 {
		for i := 1; i < len(straightFlushes); i++ {
			straightFlush := straightFlushes[i]
			highestCardOfStraightFlush := straightFlush[len(straightFlush)-1]
			highestCardOfHighestStraightFlush := highestStraightFlush[len(highestStraightFlush)-1]
			if highestCardOfStraightFlush.Index > highestCardOfHighestStraightFlush.Index {
				highestStraightFlush = straightFlush
			}
		}
	}

	return highestStraightFlush, nil
}

func getXOfAKind(groups map[int][]Card, x int) ([]Card, error) {
	var remainder, xOfAKind []Card
	for _, cards := range groups {
		if len(cards) == x {
			xOfAKind = cards
		} else {
			remainder = append(remainder, cards...)
		}
	}

	SortCards(remainder)

	highCardsNeeded := handLength - x
	nextHighestCards := remainder[len(remainder)-highCardsNeeded:]

	if xOfAKind != nil {
		output := append(xOfAKind, nextHighestCards...)
		return output, nil
	} else {
		return nil, errors.New(strconv.Itoa(x) + " of a kind not found")
	}
}
