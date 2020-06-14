package main

import (
	"encoding/json"
	"fmt"
	"go-play-2/tenis-ball/models"
	"net/http"
)

type Response struct {
	Container models.Containers `json:"container"`
	Full      int               `json:"full"`
}

func main() {
	port := ":8090"
	fmt.Println("Serving at port ", port)
	http.HandleFunc("/loads", func(w http.ResponseWriter, req *http.Request) {
		c := models.Containers{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
		full := c.Load()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			c,
			full,
		})
	})
	http.ListenAndServe(port, nil)
}
