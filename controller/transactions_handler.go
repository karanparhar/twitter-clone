package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	ethereum "github.com/ethereum_project/usecase/ethereum_repository"
	"github.com/gorilla/mux"
)

type transactionhandler struct {
	ethereum.Usecase
}

func NewHandler(r *mux.Router, u ethereum.Usecase) *mux.Router {

	t := &transactionhandler{
		u,
	}
	r.HandleFunc("/fetchblocks", t.fetchblocks).Methods("GET", "OPTIONS")
	r.HandleFunc("/gettransaction", t.gettransaction).Methods("GET", "OPTIONS")
	//this route will be used for instancegroup healthcheck
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("This is a catch-all route")
	})

	return r

}

func (t *transactionhandler) fetchblocks(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	for k, v := range query {
		if v[0] == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(k + " is missing")
			return
		}
	}

	blocks := query["blocks"][0]

	n, err := strconv.Atoi(blocks)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Not a valid number")
		return

	}

	go t.GetBlocks(uint64(n))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Processing the request")

}
func (t *transactionhandler) gettransaction(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	for k, v := range query {
		if v[0] == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(k + " is missing")
			return
		}
	}

	transaction := query["trhash"][0]

	result, err := t.GetTransactions(transaction)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
