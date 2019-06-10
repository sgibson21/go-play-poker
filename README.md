# go-play-poker

Command line poker simulator.

## Execution

```
go run main.go
```

Enter how many people are playing:

```
How many people are playing? 4
```

Enter their names:

```
Player 1) name: alice
Player 2) name: bob
Player 3) name: mickey
Player 4) name: minnie
```
Output:

```
{alice, [{6 ♣} {5 ♦}], 1000}
{bob, [{Queen ♥} {9 ♣}], 1000}
{mickey, [{4 ♠} {2 ♥}], 1000}
{minnie, [{2 ♦} {Jack ♠}], 1000}

Community Cards: [{4 ♦} {10 ♥} {2 ♣} {Queen ♣} {4 ♣}]

Best Hands:
alice: [{4 ♣} {4 ♦} {6 ♣} {10 ♥} {Queen ♣}], TWO_OF_A_KIND
bob: [{4 ♣} {4 ♦} {Queen ♣} {Queen ♥} {10 ♥}], TWO_PAIR
mickey: [{4 ♠} {4 ♦} {4 ♣} {2 ♥} {2 ♣}], FULL_HOUSE
minnie: [{2 ♦} {2 ♣} {4 ♦} {4 ♣} {Jack ♠}], TWO_PAIR
Winners:
mickey [{4 ♠} {4 ♦} {4 ♣} {2 ♥} {2 ♣}] FULL_HOUSE
```