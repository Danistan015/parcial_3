package main

import (
    "encoding/json"
    "net/http"
    "strings"
)

// variable global que puedo cambiar en pruebas (mock)
// asi los tests no usan la bd real
var clientService ClientServiceInterface = &ClientService{}

// handler para actualizar un cliente por id
func UpdateClientHandler(w http.ResponseWriter, r *http.Request) {

    // solo permito el metodo PUT
    if r.Method != http.MethodPut {
        http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
        return
    }

    // saco el id desde la ruta
    // ej: /clients/123 -> id = "123"
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 || parts[2] == "" {
        http.Error(w, "ID requerido", http.StatusBadRequest)
        return
    }
    id := parts[2]

    // leo el json que viene en el body
    var client Client
    if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
        http.Error(w, "Error al leer el JSON", http.StatusBadRequest)
        return
    }

    // llamo al servicio para actualizar
    // si el servicio falla, devuelvo 500
    if err := clientService.UpdateClient(id, client); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // respuesta correcta
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(client)
}
