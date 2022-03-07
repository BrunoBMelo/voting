package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/vote", func(w http.ResponseWriter, r *http.Request) {

		var model Vote

		err := json.NewDecoder(r.Body).Decode(&model)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		// Passing the object to global channel
		VoteChannel <- model

		w.WriteHeader(http.StatusCreated)
	})

	http.ListenAndServe(":80", r)
}

// This`s vote struct
type Vote struct {
	Key         string `json:"key"`
	CandidateId string `json:"candidateId"`
}

// This method creates a channel queue structure on initialization the application
// And call the method that will to process this channel queue as workjob service
func init() {
	VoteChannel = make(chan Vote, 1000)
	WorkerOnBackground()
}

// The global channel
var VoteChannel chan Vote

// The method that will be to process a queue of the channels
func WorkerOnBackground() {
	go func() {
		for {
			job := <-VoteChannel
			fmt.Println("Canal", job)
		}
	}()
}
