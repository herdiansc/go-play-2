package main

import (
	"encoding/json"
	"fmt"
	"go-play-2/order-race-condition/models"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	port := ":8091"
	fmt.Println("Serving at port ", port)

	stock := 2
	http.HandleFunc("/orders", func(w http.ResponseWriter, req *http.Request) {
		stockModel := models.NewStock(stock)
		currentStock, status, message := stockModel.Order()
		code := http.StatusOK
		if status != "success" {
			code = http.StatusNotFound
		}
		stock = currentStock

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(Response{
			status,
			message,
		})
	})
	http.ListenAndServe(port, nil)
}
