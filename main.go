package main

import (
    "fmt"
    "net/http"
    tenisBall "go-play-2/tenisball"
    shopOrder "go-play-2/shoporder"
    "encoding/json"
)

type Data interface {}

type BaseResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    Data   `json:"data"`
}

func main() {
    port := ":8090"
    fmt.Println("Serving on port", port)

    http.HandleFunc("/loads", func(w http.ResponseWriter, req *http.Request) {
        // assumptions:
        // - there are 5 containers
        // - each container will be full when it contains 3 balls
        c := tenisBall.Containers{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
        full := c.Load()

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(BaseResponse{
            Status: "success",
            Message: "A container has been full",
            Data: tenisBall.Response{
                ListContainers: c,
                FullContainerNo: full,
            },
        })
    })

    initialStock := 2
    http.HandleFunc("/orders", func(w http.ResponseWriter, req *http.Request) {
        stock := shopOrder.NewStock(initialStock)
        currentStock, status, message := stock.Order()
        code := http.StatusOK
        if status != "success" {
            code = http.StatusNotFound
        }
        initialStock = currentStock

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(code)
        json.NewEncoder(w).Encode(BaseResponse{
            Status: status,
            Message: message,
            Data: nil,
        })
    })
    http.ListenAndServe(port, nil)
}
