package main

import (
    "encoding/json"
    "net/http"
    "strings"
)

// 🔥 VARIABLE GLOBAL MOCKEABLE
var clientService ClientServiceInterface = &ClientService{}

func GetClientsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    clients, err := clientService.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(clients)
}

func GetClientByIDHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 {
        http.Error(w, "ID no proporcionado", http.StatusBadRequest)
        return
    }
    id := parts[2]

    client, err := clientService.GetByID(id)
    if err != nil {
        http.Error(w, "Cliente no encontrado", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(client)
}
