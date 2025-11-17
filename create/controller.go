package main

import (
    "encoding/json"
    "net/http"
)

// Interface que usaremos para poder cambiar el servicio en pruebas
type ClientCreator interface {
    Create(Client) error
}

// Por defecto, este apuntará al servicio real
var clientService ClientCreator = &ClientService{}

func CreateClientHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    var client Client
    if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
        http.Error(w, "Error al parsear el cuerpo", http.StatusBadRequest)
        return
    }

    // 👇 Aquí usamos la variable global clientService (que en test cambiaremos)
    if err := clientService.Create(client); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(client)
}
