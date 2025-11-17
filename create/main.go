package main

import (
    "log"
    "net/http"
)

func main() {
    // clientService ya apunta a &ClientService{} por defecto
    http.HandleFunc("/clients", CreateClientHandler)

    log.Println("Create service running on port 8001")
    log.Fatal(http.ListenAndServe(":8001", nil))
}
