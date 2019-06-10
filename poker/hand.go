package poker

type Hand struct {
	Cards    []Card
	Strength strength
}

func NewHand(cards []Card, strength strength) Hand {
	return Hand{cards, strength}
}
