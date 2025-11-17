package main

type ClientService struct{}

func (s *ClientService) Create(client Client) error {
    repo := ClientRepository{}
    return repo.Insert(client)
}
