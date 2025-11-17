package main

import (
    "net/http"
    "strings"
)

// 🔥 variable global mockeable
var clientService ClientServiceInterface = &ClientService{}

func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodDelete {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 || parts[2] == "" {
        http.Error(w, "ID requerido", http.StatusBadRequest)
        return
    }

    id := parts[2]

    if err := clientService.DeleteClient(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Cliente eliminado"))
}
