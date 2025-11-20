package main

// este es el servicio, basicamente aqui hago la logica para crear un cliente
type ClientService struct{}

// Create recibe un cliente y llama al repo para guardarlo en mongo
func (s *ClientService) Create(client Client) error {

    // creo el repo que es el que habla directo con la base de datos
    repo := ClientRepository{}

    // simplemente llamo a Insert y devuelvo el error si pasa algo
    return repo.Insert(client)
}
