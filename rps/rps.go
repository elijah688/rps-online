package rps

import (
	"rps/rps/util"
	"strings"
)

var rpsMap map[int]string = map[int]string{0: "rock", 1: "paper", 2: "scissors"}

type GameResponse struct {
	PlayerMover string `json:"player_move"`
	CpuMove     string `json:"cpu_move"`
	Winner      string `json:"winner"`
}

func RPS(
	playerInChan *chan *int,
	gameChan *chan *GameResponse,
) {
	for {
		var cpuMove *int = util.GenerateCpuMove()
		var playerMove *int = <-*playerInChan
		var gameResponse *GameResponse = &GameResponse{
			PlayerMover: strings.Title(rpsMap[*playerMove]),
			CpuMove:     strings.Title(rpsMap[*cpuMove]),
			Winner:      *util.ProduceWinner(playerMove, cpuMove),
		}
		*gameChan <- gameResponse
	}

}
