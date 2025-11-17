package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

// Mock del servicio que NO va a Mongo
type mockClientService struct {
    shouldFail bool
}

func (m *mockClientService) Create(c Client) error {
    if m.shouldFail {
        return assert.AnError
    }
    return nil
}

func TestCreateClientHandler_Success(t *testing.T) {
    // Guardamos el servicio real y lo restauramos al final
    oldService := clientService
    clientService = &mockClientService{shouldFail: false}
    defer func() { clientService = oldService }()

    body, _ := json.Marshal(Client{
        Name:  "Test User",
        Email: "test@example.com",
        Phone: "123456",
    })

    req := httptest.NewRequest(http.MethodPost, "/clients", bytes.NewReader(body))
    rec := httptest.NewRecorder()

    CreateClientHandler(rec, req)

    assert.Equal(t, http.StatusCreated, rec.Code)

    var resp Client
    err := json.Unmarshal(rec.Body.Bytes(), &resp)
    assert.NoError(t, err)
    assert.Equal(t, "Test User", resp.Name)
}
