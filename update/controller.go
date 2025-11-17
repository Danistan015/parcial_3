package main

import (
    "encoding/json"
    "net/http"
    "strings"
)

// 🔥 variable global mockeable
var clientService ClientServiceInterface = &ClientService{}

func UpdateClientHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPut {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 || parts[2] == "" {
        http.Error(w, "ID requerido", http.StatusBadRequest)
        return
    }

    id := parts[2]

    var client Client
    if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
        http.Error(w, "Error al leer el JSON", http.StatusBadRequest)
        return
    }

    // 🔥 ahora usa clientService (mockeable)
    if err := clientService.UpdateClient(id, client); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(client)
}
