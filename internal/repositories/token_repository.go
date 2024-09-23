package repositories

import (
	models "github.com/adisnuhic/go-graphql/internal/models"
)

// ITokenRepository contract
type ITokenRepository interface {
	Create(token *models.Token) (*models.Token, error)
	GetByToken(token string) (*models.Token, error)
}
