package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// Http verb POST
	r.HandleFunc("/vote", func(w http.ResponseWriter, r *http.Request) {

		var model Vote

		err := json.NewDecoder(r.Body).Decode(&model)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		// Passing the object to the global channel
		VoteChannel <- model

		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")

	//Http verb GET
	r.HandleFunc("/vote", func(w http.ResponseWriter, _ *http.Request) {

		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(Votes)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}).Methods("GET")

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
	Votes = make([]Vote, 1000)
	WorkerOnBackground()
}

// The global channel
var VoteChannel chan Vote
var Votes []Vote

// The method that will be to process a queue of the channels
func WorkerOnBackground() {
	go func() {
		for {
			vote := <-VoteChannel
			Votes = append(Votes, vote)
		}
	}()
}
