package handlers

import (
    "encoding/json"
    "net/http"
)

func createOrder(w http.ResponseWriter, req *http.Request) {
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
}