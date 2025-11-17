package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

// ------- Mock del servicio -------
type MockClientService struct {
    GetAllFunc  func() ([]Client, error)
    GetByIDFunc func(id string) (Client, error)
}

func (m *MockClientService) GetAll() ([]Client, error) {
    return m.GetAllFunc()
}

func (m *MockClientService) GetByID(id string) (Client, error) {
    return m.GetByIDFunc(id)
}

// ------- Tests --------

// GET /clients
func TestGetClientsHandler_Success(t *testing.T) {

    mock := &MockClientService{
        GetAllFunc: func() ([]Client, error) {
            return []Client{
                {ID: "1", Name: "John", Email: "john@test.com", Phone: "123"},
            }, nil
        },
    }

    clientService = mock

    req := httptest.NewRequest("GET", "/clients", nil)
    rr := httptest.NewRecorder()

    GetClientsHandler(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var clients []Client
    json.Unmarshal(rr.Body.Bytes(), &clients)

    assert.Equal(t, 1, len(clients))
    assert.Equal(t, "John", clients[0].Name)
}

// GET /clients/:id
func TestGetClientByIDHandler_Success(t *testing.T) {

    mock := &MockClientService{
        GetByIDFunc: func(id string) (Client, error) {
            return Client{ID: id, Name: "Maria"}, nil
        },
    }

    clientService = mock

    req := httptest.NewRequest("GET", "/clients/999", nil)
    rr := httptest.NewRecorder()

    GetClientByIDHandler(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var client Client
    json.Unmarshal(rr.Body.Bytes(), &client)

    assert.Equal(t, "Maria", client.Name)
}
