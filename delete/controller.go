package main

import (
    "net/http"
    "strings"
)

// variable global que puedo cambiar en las pruebas (para mockear)
// asi no tengo que usar el servicio real
var clientService ClientServiceInterface = &ClientService{}

// este handler borra un cliente segun el id que venga en la ruta
func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {

    // solo dejo usar DELETE, los demas metodos no sirven aqui
    if r.Method != http.MethodDelete {
        http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
        return
    }

    // divido la ruta por "/" para sacar el id
    // ej: /clients/123 -> ["", "clients", "123"]
    parts := strings.Split(r.URL.Path, "/")

    // si no viene el id, mando error
    if len(parts) < 3 || parts[2] == "" {
        http.Error(w, "ID requerido", http.StatusBadRequest)
        return
    }

    id := parts[2]

    // llamo al servicio para borrar, si falla mando 500
    if err := clientService.DeleteClient(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // si todo sale bien mando estado 200 con mensaje
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Cliente eliminado"))
}
