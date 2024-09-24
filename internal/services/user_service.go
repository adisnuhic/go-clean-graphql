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
	GetPostsByUserID(ctx context.Context, userID uint64) ([]*models.Post, error)
}

type userService struct {
	Logger              log.ILogger
	AuthService         IAuthService
	AuthProviderService IAuthProviderService

	Repository     repositories.IUserRepository
	PostRepository repositories.IPostRepository
}

// NewUserService creates new user service
func NewUserService(logger log.ILogger, authSvc IAuthService, authProviderSvc IAuthProviderService, repository repositories.IUserRepository, postRepository repositories.IPostRepository) IUserService {
	return &userService{
		Logger:              logger,
		AuthService:         authSvc,
		AuthProviderService: authProviderSvc,
		Repository:          repository,
		PostRepository:      postRepository,
	}
}

// CreateUser creates new user
func (s *userService) Create(ctx context.Context, input models.NewRegisterUser) (*models.User, error) {

	// fake confirmation
	isConfirmed := true

	// check if user already exists
	existingUser, errUser := s.Repository.GetByEmail(input.Email)
	if errUser == nil && existingUser != nil {
		return nil, errUser
	}

	// generate hash
	hash, errHash := s.AuthService.GeneratePasswordHash(input.Password)
	if errHash != nil {
		return nil, errHash
	}

	// create user
	user := &models.User{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		IsConfirmed: &isConfirmed,
		AcceptedTos: isConfirmed,
	}

	u, errUser := s.Repository.Create(user)
	if errUser != nil {
		return nil, errUser
	}

	// create auth provider
	_, errAuth := s.AuthProviderService.Create("local", hash, u.ID)
	if errAuth != nil {
		return nil, errAuth
	}

	return u, nil
}

// GetUsers returns all users
func (s *userService) GetAll(ctx context.Context) ([]*models.User, error) {
	return s.Repository.GetAll()
}

// GetPostsByUserID return posts for provided userID
func (s *userService) GetPostsByUserID(ctx context.Context, userID uint64) ([]*models.Post, error) {
	return s.PostRepository.GetAllByUserID(userID)
}
