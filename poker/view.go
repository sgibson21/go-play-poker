package poker

import (
	"fmt"
)

func GetPlayers() []Player {
	numPlayers := getNumPlayers()
	var players []Player
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Player %d) ", i+1)
		player := getPlayer()
		players = append(players, player)
	}
	return players
}

func getPlayer() Player {
	fmt.Print("name: ")
	name := readString()

	return NewPlayer(name)
}

func getNumPlayers() int {
	fmt.Print("How many people are playing? ")
	return readInt()
}

func readString() string {
	var s string
	_, err := fmt.Scanf("%s", &s)
	fmt.Scanln()
	if err != nil {
		panic(err)
	}
	return s
}

func readInt() int {
	var i int
	_, err := fmt.Scanf("%d", &i)
	fmt.Scanln()

	if err != nil {
		panic(err)
	}
	return i
}
