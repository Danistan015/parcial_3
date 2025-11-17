package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

// -------- Mock del servicio --------
type MockUpdateService struct {
    ShouldFail bool
}

func (m *MockUpdateService) UpdateClient(id string, client Client) error {
    if m.ShouldFail {
        return assert.AnError
    }
    return nil
}

// -------- TESTS --------

// ÉXITO — 200 OK
func TestUpdateClientHandler_Success(t *testing.T) {

    mock := &MockUpdateService{ShouldFail: false}
    clientService = mock

    body := Client{
        Name:  "UpdatedName",
        Email: "new@test.com",
        Phone: "999",
    }
    jsonBody, _ := json.Marshal(body)

    req := httptest.NewRequest(http.MethodPut, "/clients/12345", bytes.NewReader(jsonBody))
    rr := httptest.NewRecorder()

    UpdateClientHandler(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var resp Client
    json.Unmarshal(rr.Body.Bytes(), &resp)

    assert.Equal(t, "UpdatedName", resp.Name)
}

// ERROR — 500
func TestUpdateClientHandler_Fail(t *testing.T) {

    mock := &MockUpdateService{ShouldFail: true}
    clientService = mock

    body := Client{Name: "Test"}
    jsonBody, _ := json.Marshal(body)

    req := httptest.NewRequest(http.MethodPut, "/clients/12345", bytes.NewReader(jsonBody))
    rr := httptest.NewRecorder()

    UpdateClientHandler(rr, req)

    assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

// ERROR — 400 BAD REQUEST POR FALTA DE ID
func TestUpdateClientHandler_BadRequest(t *testing.T) {

    mock := &MockUpdateService{}
    clientService = mock

    req := httptest.NewRequest(http.MethodPut, "/clients/", nil)
    rr := httptest.NewRecorder()

    UpdateClientHandler(rr, req)

    assert.Equal(t, http.StatusBadRequest, rr.Code)
}
