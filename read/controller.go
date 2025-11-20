package main

import (
    "encoding/json"
    "net/http"
    "strings"
)

// variable global que puedo cambiar en pruebas (mock)
// asi los tests no usan la BD real
var clientService ClientServiceInterface = &ClientService{}

// handler para traer todos los clientes
func GetClientsHandler(w http.ResponseWriter, r *http.Request) {

    // solo dejo usar GET
    if r.Method != http.MethodGet {
        http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
        return
    }

    // llamo al servicio para traer todo
    clients, err := clientService.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // devuelvo el json con los clientes
    json.NewEncoder(w).Encode(clients)
}

// handler para traer un cliente por id
func GetClientByIDHandler(w http.ResponseWriter, r *http.Request) {

    // solo GET funciona aqui
    if r.Method != http.MethodGet {
        http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
        return
    }

    // saco el id de la ruta
    // ej: /clients/123 -> ["", "clients", "123"]
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 {
        http.Error(w, "ID no proporcionado", http.StatusBadRequest)
        return
    }
    id := parts[2]

    // llamo al servicio para buscar el cliente
    client, err := clientService.GetByID(id)
    if err != nil {
        http.Error(w, "Cliente no encontrado", http.StatusNotFound)
        return
    }

    // mando el json del cliente encontrado
    json.NewEncoder(w).Encode(client)
}
