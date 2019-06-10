package poker

import (
	"fmt"
)

type Player struct {
	Name    string
	Cards   []Card
	Balance int
	Hand    Hand
}

func NewPlayer(name string) Player {
	return Player{name, []Card{}, 1000, Hand{}}
}

func (p *Player) TakeCard(card Card) {
	p.Cards = append(p.Cards, card)
}

func (p *Player) String() string {
	return fmt.Sprintf("{%s, %s, %d}", p.Name, p.Cards, p.Balance)
}

func PrintPlayers(players []Player) {
	for i := 0; i < len(players); i++ {
		fmt.Println(players[i].String())
	}
}
