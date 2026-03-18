package repository

import "trueAPI/internal/models"

type UserRepository struct {
	users  []models.User
	nextID int
}

type UserRepositoryInterface interface {
	CreateUser(username string, age int) models.User
	GetUserByID(id int) (models.User, bool)
	UpdateUser(id int, username string, age int) (models.User, bool)
	DeleteUser(id int) bool
	GetAllUsers() []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  []models.User{},
		nextID: 1,
	}
}

func (r *UserRepository) CreateUser(username string, age int) models.User {
	user := models.User{
		ID:       r.nextID,
		Username: username,
		Age:      age,
	}
	r.users = append(r.users, user)
	r.nextID++
	return user
}

func (r *UserRepository) GetUserByID(id int) (models.User, bool) {
	for _, user := range r.users {
		if user.ID == id {
			return user, true
		}
	}
	return models.User{}, false
}

func (r *UserRepository) UpdateUser(id int, username string, age int) (models.User, bool) {
	for i, user := range r.users {
		if user.ID == id {
			r.users[i] = models.User{
				ID:       id,
				Username: username,
				Age:      age,
			}
			return r.users[i], true
		}
	}
	return models.User{}, false
}

func (r *UserRepository) DeleteUser(id int) bool {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return true
		}
	}
	return false
}

func (r *UserRepository) GetAllUsers() []models.User {
	return r.users
}

