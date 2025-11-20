package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

// este es un mock para simular el servicio real, asi no usamos mongo en la prueba
type mockClientService struct {
    shouldFail bool
}

func (m *mockClientService) Create(c Client) error {
    if m.shouldFail {
        return assert.AnError // si queremos que falle, devolvemos un error
    }
    return nil
}

// prueba cuando todo sale bien en el handler
func TestCreateClientHandler_Success(t *testing.T) {

    // guardamos el servicio real para luego dejarlo como estaba
    oldService := clientService

    // aqui pongo el mock para que la prueba no toque la BD real
    clientService = &mockClientService{shouldFail: false}

    // al final vuelvo a dejar el servicio original
    defer func() { clientService = oldService }()

    // creo un json como si viniera del cliente
    body, _ := json.Marshal(Client{
        Name:  "Test User",
        Email: "test@example.com",
        Phone: "123456",
    })

    // creo una peticion falsa POST y un recorder para ver la respuesta
    req := httptest.NewRequest(http.MethodPost, "/clients", bytes.NewReader(body))
    rec := httptest.NewRecorder()

    // ejecuto el handler
    CreateClientHandler(rec, req)

    // reviso que responda 201 = creado
    assert.Equal(t, http.StatusCreated, rec.Code)

    // leo el json que respondio el handler
    var resp Client
    err := json.Unmarshal(rec.Body.Bytes(), &resp)

    // verifico que no haya error y que el nombre sea el mismo
    assert.NoError(t, err)
    assert.Equal(t, "Test User", resp.Name)
}
