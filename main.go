package main

import (
	"encoding/json"
	"log"
	"net/http"
	"odd-number-service/number"
	"strconv"
)

/*
{
	"oddNumbers": [1,3,5,7]
}
*/
type oddResponse struct {
	OddNumbers []int `json:"oddNumbers"`
}

func main() {
	http.HandleFunc("/api/odd", oddHandler)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatalf("Listen n Serve error %v", err)
	}

}

func oddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	param := r.URL.Query().Get("max")
	if len(param) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	num, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := number.FindOdds(num)
	odd := oddResponse{
		OddNumbers: result,
	}
	data, err := json.Marshal(&odd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return

}
