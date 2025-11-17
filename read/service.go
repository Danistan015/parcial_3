package main

type ClientServiceInterface interface {
    GetAll() ([]Client, error)
    GetByID(id string) (Client, error)
}

type ClientService struct{}

func (s *ClientService) GetAll() ([]Client, error) {
    repo := ClientRepository{}
    return repo.GetAll()
}

func (s *ClientService) GetByID(id string) (Client, error) {
    repo := ClientRepository{}
    return repo.GetByID(id)
}
