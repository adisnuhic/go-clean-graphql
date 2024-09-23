package services

import (
	"context"

	models "github.com/adisnuhic/go-graphql/internal/models"
	"github.com/adisnuhic/go-graphql/internal/repositories"
	"github.com/adisnuhic/go-graphql/pkg/log"
)

// IUserService represents user service contract
type IUserService interface {
	Create(ctx context.Context, input models.NewRegisterUser) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
}

type userService struct {
	Logger     log.ILogger
	Repository repositories.IUserRepository
}

// NewUserService creates new user service
func NewUserService(logger log.ILogger, repository repositories.IUserRepository) IUserService {
	return &userService{
		Logger:     logger,
		Repository: repository,
	}
}

// CreateUser creates new user
func (s *userService) Create(ctx context.Context, input models.NewRegisterUser) (*models.User, error) {
	user := &models.User{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		IsConfirmed: input.IsConfirmed,
		AcceptedTos: &input.AcceptedTos,
	}
	return s.Repository.Create(user)
}

// GetUsers returns all users
func (s *userService) GetAll(ctx context.Context) ([]*models.User, error) {
	return s.Repository.GetAll()
}
