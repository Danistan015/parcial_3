package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

// Mock
type MockDeleteService struct {
    ShouldFail bool
}

func (m *MockDeleteService) DeleteClient(id string) error {
    if m.ShouldFail {
        return assert.AnError
    }
    return nil
}

// TEST: éxito
func TestDeleteClientHandler_Success(t *testing.T) {
    mock := &MockDeleteService{ShouldFail: false}
    clientService = mock

    req := httptest.NewRequest(http.MethodDelete, "/clients/123", nil)
    rr := httptest.NewRecorder()

    DeleteClientHandler(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Contains(t, rr.Body.String(), "Cliente eliminado")
}

// TEST: falta ID → 400
func TestDeleteClientHandler_BadRequest(t *testing.T) {
    mock := &MockDeleteService{}
    clientService = mock

    req := httptest.NewRequest(http.MethodDelete, "/clients/", nil)
    rr := httptest.NewRecorder()

    DeleteClientHandler(rr, req)

    assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// TEST: error del servicio → 500
func TestDeleteClientHandler_Fail(t *testing.T) {
    mock := &MockDeleteService{ShouldFail: true}
    clientService = mock

    req := httptest.NewRequest(http.MethodDelete, "/clients/123", nil)
    rr := httptest.NewRecorder()

    DeleteClientHandler(rr, req)

    assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
