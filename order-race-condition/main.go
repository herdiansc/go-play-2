package main

import (
    "encoding/json"
    "fmt"
    "go-play-2/order-race-condition/handlers"
    "net/http"
)

type Response struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func main() {
    port := ":8091"
    fmt.Println("Serving at port ", port)
    stocks := 2
    http.HandleFunc("/orders", func(w http.ResponseWriter, req *http.Request) {
        code := http.StatusOK
        status := "success"
        message := "order successfully created, please proceed to payment!"
        if stocks <=0 {
            code = http.StatusNotFound
            status = "error"
            message = "stock is not found!"
        } else {
            stocks--
        }
        
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(code)
        json.NewEncoder(w).Encode(Response{
            status,
            message,
        })
    })
    http.ListenAndServe(port, nil)
}
