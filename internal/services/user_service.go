package services

import (
	"errors"
	"strings"
	"trueAPI/internal/models"
	"trueAPI/internal/repository"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() []models.User
	GetUserByID(id int) (models.User, error)
}

type userService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	if strings.TrimSpace(user.Username) == "" {
		return models.User{}, errors.New("username cannot be empty")
	}
	if user.Age <= 0 {
		return models.User{}, errors.New("age must be a positive integer")
	}

	createdUser := s.repo.CreateUser(user.Username, user.Age)
	return createdUser, nil
}

func (s *userService) GetAllUsers() []models.User {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id int) (models.User, error) {
	user, found := s.repo.GetUserByID(id)
	if !found {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}
