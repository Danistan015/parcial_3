package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/clients/", DeleteClientHandler)

    log.Println("Delete service running on port 8004")
    log.Fatal(http.ListenAndServe(":8004", nil))
}
