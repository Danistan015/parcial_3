package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/clients/", UpdateClientHandler)

    log.Println("Update service running on port 8003")
    log.Fatal(http.ListenAndServe(":8003", nil))
}
