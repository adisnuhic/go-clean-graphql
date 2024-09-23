package repositories

import models "github.com/adisnuhic/go-graphql/internal/models"

// IAuthProviderRepository contract
type IAuthProviderRepository interface {
	Create(authProvider *models.AuthProvider) (*models.AuthProvider, error)
}
