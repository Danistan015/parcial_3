package main

import (
    "log"
    "net/http"
)

func main() {
    // aca el clientService ya esta listo porque viene del controller
    // y apunta al servicio real que crea clientes

    // registro la ruta /clients y le digo que use el handler que ya hice
    http.HandleFunc("/clients", CreateClientHandler)

    // mensaje para saber que el microservicio ya esta corriendo
    log.Println("Create service running on port 8001")

    // aca arranca el servidor web en el puerto 8001
    // si falla, se muere la app con fatal
    log.Fatal(http.ListenAndServe(":8001", nil))
}
