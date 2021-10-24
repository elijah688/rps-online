package util

import (
	"math/rand"
	"time"
)

func GenerateCpuMove() *int {
	rand.Seed(time.Now().UnixNano())
	rng := rand.Intn(3)
	return &rng
}

func ProduceWinner(playerMove, cpuMove *int) *string {
	var outcome string
	if *playerMove == *cpuMove {
		outcome = "Draw!"
	} else if *playerMove == (*cpuMove+1)%3 {
		outcome = "Player Wins!"
	} else {
		outcome = "Cpu Wins!"
	}
	return &outcome
}
