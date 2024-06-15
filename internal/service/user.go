package service

import (
	"context"

	"github.com/chaninlaw/auth/internal/repository"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `validate:"required,min=3,max=32"`
	Password string `validate:"required,min=8,max=32"`
	IsAdmin  bool
}

type UserService struct {
	repo repository.UserRepository
	validate *validator.Validate
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo, 
		validate: validator.New(),
	}
}

func (s *UserService) FetchAllUsers(ctx context.Context) ([]repository.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) FetchUser(ctx context.Context, id string) (*repository.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, user User) error {
	if err := s.validate.Struct(user); err != nil {
		return err
	}

	repoUser := repository.User{
		ID:       0,
		Username: user.Username,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}

	return s.repo.Create(ctx, &repoUser)
}