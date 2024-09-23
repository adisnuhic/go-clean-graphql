package repositories

import (
	"github.com/adisnuhic/go-graphql/internal/initialize/mysql"
	models "github.com/adisnuhic/go-graphql/internal/models"
)

// ITokenRepository
type ITokenRepository interface {
	Create(token *models.Token) (*models.Token, error)
	GetByToken(token string) (*models.Token, error)
}

type tokenRepository struct {
	Store mysql.Store
}

// NewMySQLTokenRepository -
func NewMySQLTokenRepository(store mysql.Store) ITokenRepository {
	return &tokenRepository{
		Store: store,
	}
}

// Create a new token
func (r *tokenRepository) Create(token *models.Token) (*models.Token, error) {

	if err := r.Store.Create(&token).Error; err != nil {
		return nil, err
	}

	return token, nil
}

// GetByToken returns token for provided token string
func (r tokenRepository) GetByToken(token string) (*models.Token, error) {
	model := new(models.Token)

	tx := r.Store.Where("token = ?", token).Find(&model)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return model, nil
}
