package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type Response struct {
}

func handle(w http.ResponseWriter, r *http.Request) {
	time.Sleep(200 * time.Millisecond)

	json.NewEncoder(w).Encode(Response{})
}

func main() {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":8088", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed")
	} else if err != nil {
		log.Println(err)
	}
}
