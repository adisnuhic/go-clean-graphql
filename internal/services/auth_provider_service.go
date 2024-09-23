package services

import (
	models "github.com/adisnuhic/go-graphql/internal/models"
	"github.com/adisnuhic/go-graphql/internal/repositories"
	"github.com/adisnuhic/go-graphql/pkg/log"
)

// IAuthProviderService represents auth provider service contract
type IAuthProviderService interface {
	Create(provider, uuid string, userID uint64) (*models.AuthProvider, error)
}

type authProviderService struct {
	Logger     log.ILogger
	Repository repositories.IAuthProviderRepository
}

// NewAuthProviderService creates new auth provider service
func NewAuthProviderService(logger log.ILogger, repository repositories.IAuthProviderRepository) IAuthProviderService {
	return &authProviderService{
		Logger:     logger,
		Repository: repository,
	}
}

// Create creates new auth provider
func (s *authProviderService) Create(provider, uuid string, userID uint64) (*models.AuthProvider, error) {
	authProvider := &models.AuthProvider{
		Provider: provider,
		UID:      uuid,
		UserID:   userID,
	}

	return s.Repository.Create(authProvider)
}
