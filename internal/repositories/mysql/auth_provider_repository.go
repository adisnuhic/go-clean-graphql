package repositories

import (
	"github.com/adisnuhic/go-graphql/internal/initialize/mysql"
	models "github.com/adisnuhic/go-graphql/internal/models"
)

// IAuthProviderRepository interface
type IAuthProviderRepository interface {
	Create(authProvider *models.AuthProvider) (*models.AuthProvider, error)
}

type authProviderRepository struct {
	Store mysql.Store
}

// NewMySQLAuthProviderRepository -
func NewMySQLAuthProviderRepository(store mysql.Store) IAuthProviderRepository {
	return &authProviderRepository{
		Store: store,
	}
}

// Create a new auth provider
func (r *authProviderRepository) Create(authProvider *models.AuthProvider) (*models.AuthProvider, error) {

	if err := r.Store.Create(&authProvider).Error; err != nil {
		return nil, err
	}

	return authProvider, nil
}
