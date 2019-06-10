package poker

func groupByNumbers(hand []Card) map[int][]Card {
	groups := make(map[int][]Card)

	for i := 0; i < len(hand); i++ {
		group := groups[hand[i].Index]
		groups[hand[i].Index] = append(group, hand[i])
	}

	return groups
}

func groupBySuit(hand []Card) map[Suit][]Card {
	groups := make(map[Suit][]Card)

	for i := 0; i < len(hand); i++ {
		group := groups[hand[i].Suit]
		groups[hand[i].Suit] = append(group, hand[i])
	}

	return groups
}
