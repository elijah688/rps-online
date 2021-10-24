package rps

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Handler(
	w http.ResponseWriter,
	r *http.Request,
	playerInChan *chan *int,
	gameChan *chan *GameResponse,
) {

	rpsIntVal, err := strconv.Atoi(r.URL.Query().Get("rps"))
	if err != nil || rpsIntVal < 0 || rpsIntVal > 2 {
		http.Error(w, "Invalid rps parameter!", 400)
		return
	}

	*playerInChan <- &rpsIntVal
	e := json.NewEncoder(w)
	e.SetIndent("", "     ")
	err = e.Encode(*<-*gameChan)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
