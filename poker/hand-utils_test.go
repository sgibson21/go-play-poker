package poker

import "testing"

func TestGroupByNumbers(t *testing.T) {
	input := []Card{
		NewCard(2, Heart), // assume cards are sorted in order first
		NewCard(2, Heart),
		NewCard(6, Heart),
		NewCard(6, Heart),
		NewCard(10, Heart),
		NewCard(11, Heart),
	}

	output := groupByNumbers(input)

	t.Logf("output: %v", output)
}
