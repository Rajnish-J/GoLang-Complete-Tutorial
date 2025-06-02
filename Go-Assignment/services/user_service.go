package services

import (
	"go-assignment/models"
	"go-assignment/repository"
)

type UserService interface {
    CreateUser(models.User) error
    GetUserByID(string) (models.User, error)
    UpdateUser(models.User) error
    DeleteUser(string) error
    ListUsers() ([]models.User, error)
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo}
}

func (s *userService) CreateUser(u models.User) error {
    return s.repo.Create(u)
}

func (s *userService) GetUserByID(id string) (models.User, error) {
    return s.repo.GetByID(id)
}

func (s *userService) UpdateUser(u models.User) error {
    return s.repo.Update(u)
}

func (s *userService) DeleteUser(id string) error {
    return s.repo.Delete(id)
}

func (s *userService) ListUsers() ([]models.User, error) {
    return s.repo.GetAll()
}
