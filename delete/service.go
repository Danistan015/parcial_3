package main

type ClientServiceInterface interface {
    DeleteClient(id string) error
}

type ClientService struct{}

func (s *ClientService) DeleteClient(id string) error {
    repo := ClientRepository{}
    return repo.Delete(id)
}
