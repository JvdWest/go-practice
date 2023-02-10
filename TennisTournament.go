package main

import "math"

func Solution10(P int, C int) int {
	gamesForPlayers := P / 2
	return int(math.Min(float64(gamesForPlayers), float64(C)))
}
