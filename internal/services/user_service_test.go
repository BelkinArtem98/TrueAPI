package services

import (
	"testing"
	"trueAPI/internal/models"
	"trueAPI/internal/repository"
)

func TestCreateUserAssignsID(t *testing.T) {
	service := NewUserService(repository.NewUserRepository())

	user, err := service.CreateUser(models.User{Username: "alice", Age: 25})
	if err != nil {
		t.Fatalf("CreateUser returned unexpected error: %v", err)
	}

	if user.ID != 1 {
		t.Fatalf("expected ID 1, got %d", user.ID)
	}
}

func TestCreateUserRejectsInvalidInput(t *testing.T) {
	service := NewUserService(repository.NewUserRepository())

	if _, err := service.CreateUser(models.User{Username: " ", Age: 25}); err == nil {
		t.Fatal("expected error for empty username")
	}

	if _, err := service.CreateUser(models.User{Username: "alice", Age: 0}); err == nil {
		t.Fatal("expected error for non-positive age")
	}
}

func TestUpdateAndDeleteUser(t *testing.T) {
	service := NewUserService(repository.NewUserRepository())

	createdUser, err := service.CreateUser(models.User{Username: "alice", Age: 25})
	if err != nil {
		t.Fatalf("CreateUser returned unexpected error: %v", err)
	}

	updatedUser, err := service.UpdateUser(createdUser.ID, models.User{Username: "bob", Age: 30})
	if err != nil {
		t.Fatalf("UpdateUser returned unexpected error: %v", err)
	}

	if updatedUser.Username != "bob" || updatedUser.Age != 30 {
		t.Fatalf("unexpected updated user: %+v", updatedUser)
	}

	if err := service.DeleteUser(createdUser.ID); err != nil {
		t.Fatalf("DeleteUser returned unexpected error: %v", err)
	}

	if _, err := service.GetUserByID(createdUser.ID); err == nil {
		t.Fatal("expected deleted user to be missing")
	}
}
