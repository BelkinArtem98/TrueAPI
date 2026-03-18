package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"trueAPI/internal/repository"
	"trueAPI/internal/services"
)

func newTestHandler() *UserHandler {
	service := services.NewUserService(repository.NewUserRepository())
	return NewUserHandler(service)
}

func TestCreateUser(t *testing.T) {
	handler := newTestHandler()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"username":"alice","age":25}`))
	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
}

func TestGetAllUsers(t *testing.T) {
	handler := newTestHandler()

	createReq := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"username":"alice","age":25}`))
	createRec := httptest.NewRecorder()
	handler.CreateUser(createRec, createReq)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	handler.GetAllUsers(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"username":"alice"`) {
		t.Fatalf("expected response body to contain created user, got %s", rec.Body.String())
	}
}

func TestUpdateAndDeleteUser(t *testing.T) {
	handler := newTestHandler()

	createReq := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"username":"alice","age":25}`))
	createRec := httptest.NewRecorder()
	handler.CreateUser(createRec, createReq)

	updateReq := httptest.NewRequest(http.MethodPut, "/users?id=1", strings.NewReader(`{"username":"bob","age":30}`))
	updateRec := httptest.NewRecorder()
	handler.UpdateUser(updateRec, updateReq)

	if updateRec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, updateRec.Code)
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/users?id=1", nil)
	deleteRec := httptest.NewRecorder()
	handler.DeleteUser(deleteRec, deleteReq)

	if deleteRec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, deleteRec.Code)
	}
}
