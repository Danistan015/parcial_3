package main

import (
    "encoding/json"
    "net/http"
)

// Interfaz para cambiar el servicio en pruebas
type ClientCreator interface {
    Create(Client) error
}

//  apunta al servicio real
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

    //  usamos la variable global clientService 
    if err := clientService.Create(client); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(client)
}
