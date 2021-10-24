package main

import (
	"net/http"
	"rps/rps"
)

func main() {
	var playerInChan chan *int = make(chan *int)
	var gameChan chan *rps.GameResponse = make(chan *rps.GameResponse)
	go rps.RPS(&playerInChan, &gameChan)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./html"))
	mux.Handle("/", fs)

	mux.HandleFunc(
		"/rps",
		func(w http.ResponseWriter, r *http.Request) {
			rps.Handler(w, r, &playerInChan, &gameChan)
		})

	http.ListenAndServe(":8080", mux)

}
