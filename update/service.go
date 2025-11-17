package main

type ClientServiceInterface interface {
    UpdateClient(id string, client Client) error
}

type ClientService struct{}

func (s *ClientService) UpdateClient(id string, client Client) error {
    repo := ClientRepository{}
    return repo.Update(id, client)
}
