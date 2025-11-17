package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/clients", GetClientsHandler)
    http.HandleFunc("/clients/", GetClientByIDHandler)

    log.Println("Read service running on port 8002")
    log.Fatal(http.ListenAndServe(":8002", nil))
}
