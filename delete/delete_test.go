package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

// este es un mock para no usar la base de datos real
// solo simula si el delete funciona o si falla
type MockDeleteService struct {
    ShouldFail bool
}

func (m *MockDeleteService) DeleteClient(id string) error {
    if m.ShouldFail {
        return assert.AnError // simulo un error si quiero que falle
    }
    return nil
}

// TEST: cuando todo sale bien
func TestDeleteClientHandler_Success(t *testing.T) {
    mock := &MockDeleteService{ShouldFail: false}
    clientService = mock // uso el mock en vez del servicio real

    // creo una peticion DELETE falsa con un id
    req := httptest.NewRequest(http.MethodDelete, "/clients/123", nil)
    rr := httptest.NewRecorder()

    // ejecuto el handler
    DeleteClientHandler(rr, req)

    // verifico que devuelva 200 OK
    assert.Equal(t, http.StatusOK, rr.Code)

    // reviso que el mensaje diga "Cliente eliminado"
    assert.Contains(t, rr.Body.String(), "Cliente eliminado")
}

// TEST: si no mando el id, debe dar 400
func TestDeleteClientHandler_BadRequest(t *testing.T) {
    mock := &MockDeleteService{}
    clientService = mock

    // aqui mando la ruta sin id
    req := httptest.NewRequest(http.MethodDelete, "/clients/", nil)
    rr := httptest.NewRecorder()

    DeleteClientHandler(rr, req)

    // como falta el id, debe ser 400
    assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// TEST: si el servicio falla, debe devolver 500
func TestDeleteClientHandler_Fail(t *testing.T) {
    mock := &MockDeleteService{ShouldFail: true}
    clientService = mock // mock que falla

    req := httptest.NewRequest(http.MethodDelete, "/clients/123", nil)
    rr := httptest.NewRecorder()

    DeleteClientHandler(rr, req)

    // si el servicio fallo, debe dar 500
    assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
