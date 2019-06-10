package poker

type strength string

const (
	HIGH_CARD       strength = "HIGH_CARD"
	TWO_OF_A_KIND   strength = "TWO_OF_A_KIND"
	TWO_PAIR        strength = "TWO_PAIR"
	THREE_OF_A_KIND strength = "THREE_OF_A_KIND"
	STRAIGHT        strength = "STRAIGHT"
	FLUSH           strength = "FLUSH"
	FULL_HOUSE      strength = "FULL_HOUSE"
	FOUR_OF_A_KIND  strength = "FOUR_OF_A_KIND"
	STRAIGHT_FLUSH  strength = "STRAIGHT_FLUSH"
)

var strengthMap = map[strength]int{
	HIGH_CARD:       1,
	TWO_OF_A_KIND:   2,
	TWO_PAIR:        3,
	THREE_OF_A_KIND: 4,
	STRAIGHT:        5,
	FLUSH:           6,
	FULL_HOUSE:      7,
	FOUR_OF_A_KIND:  8,
	STRAIGHT_FLUSH:  9,
}

var comparisonMap = map[strength]func(players []Player) []Player{
	HIGH_CARD:       compareHighCard,
	TWO_OF_A_KIND:   getCompareXOfAKindFunction(2),
	TWO_PAIR:        compareTwoPair,
	THREE_OF_A_KIND: getCompareXOfAKindFunction(3),
	STRAIGHT:        compareHighCard,
	FLUSH:           compareHighCard,
	FULL_HOUSE:      compareFullHouse,
	FOUR_OF_A_KIND:  getCompareXOfAKindFunction(4),
	STRAIGHT_FLUSH:  compareHighCard,
}

func GetWinningPlayers(players []Player) []Player {
	var winningPlayers []Player
	strength := HIGH_CARD // default to lowest strength
	for i := 0; i < len(players); i++ {
		_strength := players[i].Hand.Strength
		if strengthMap[_strength] > strengthMap[strength] {
			winningPlayers = []Player{players[i]}
			strength = _strength
		} else if strengthMap[_strength] == strengthMap[strength] {
			compareFunction := comparisonMap[strength]
			winningPlayers = compareFunction(append(winningPlayers, players[i]))
		}
	}
	return winningPlayers
}

/**
 * Compare hands by highest card. Highest card is the last card when the cards are sorted with
 * the kickers at the end.
 */
func compareHighCard(players []Player) []Player {
	return comparePlayersByCardIndex(players, -1)
}

/**
 * Compare hands by the highest kicker in the hand. The highest kicker is the last card in the hand.
 * We work back from the end of the hand by the amount of kickers we are told there is.
 * eg: a two of a kind hand has 3 kickers and a four of a kind hand has 1 kicker
 */
func compareKickers(players []Player, numberOfKickers int) []Player {
	var playersWithHighestCard []Player
	cardIndex := 0

	for numberOfKickers > 0 && len(playersWithHighestCard) != 1 {
		numberOfKickers--
		cardIndex--
		playersWithHighestCard = comparePlayersByCardIndex(players, cardIndex)
	}

	return playersWithHighestCard
}

/**
 * Function that returns a comparison function that compares hand of X of a kind.
 * When comparing X of a kind hands (assuming the cards are ordered by
 * the X of a kind then by kickers) we just need to check the card at index 0.
 *
 * If this comparison can not get a single winner, then we look for the high card in the kickers.
 */
func getCompareXOfAKindFunction(x int) func(players []Player) []Player {
	return func(players []Player) []Player {
		playersWithHighestCard := comparePlayersByCardIndex(players, 0)

		if len(playersWithHighestCard) > 1 {
			playersWithHighestCard = compareKickers(players, handLength-x)
		}

		return playersWithHighestCard
	}
}

/**
 * When comparing two pair hands, we compare the highest of the two pairs.
 * If they are equal, we compare the other pair.
 * If they are equal we look for the high card in the kickers.
 */
func compareTwoPair(players []Player) []Player {
	playersWithHighestCard := comparePlayersByCardIndex(players, 0)

	if len(playersWithHighestCard) > 1 {
		playersWithHighestCard = comparePlayersByCardIndex(players, 2)
	}

	if len(playersWithHighestCard) > 1 {
		playersWithHighestCard = compareKickers(players, handLength-4)
	}

	return playersWithHighestCard
}

/**
 * When comparing full houses, we first compare the three of a kind.
 * If those are equal, we then compare the two of a kind.
 */
func compareFullHouse(players []Player) []Player {
	playersWithHighestCard := comparePlayersByCardIndex(players, 0)

	if len(playersWithHighestCard) > 1 {
		playersWithHighestCard = comparePlayersByCardIndex(players, 3)
	}

	return playersWithHighestCard
}

/**
 * Function to compare players by their cards at a given index.
 */
func comparePlayersByCardIndex(players []Player, cardIndex int) []Player {
	var playersWithHighestCard []Player
	var highCard Card
	actualIndex := cardIndex // default to what is passed in

	for i := 0; i < len(players); i++ {

		playerHand := players[i].Hand

		// if the index passed in is negative, start at the end of the array.
		if cardIndex < 0 {
			actualIndex = len(playerHand.Cards) + cardIndex
		}

		if playerHand.Cards[actualIndex].Index > highCard.Index {
			playersWithHighestCard = []Player{players[i]}
			highCard = playerHand.Cards[actualIndex]
		} else if playerHand.Cards[actualIndex].Index == highCard.Index {
			playersWithHighestCard = append(playersWithHighestCard, players[i])
		}
	}

	return playersWithHighestCard
}
