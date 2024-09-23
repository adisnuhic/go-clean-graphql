package services

import (
	"time"

	models "github.com/adisnuhic/go-graphql/internal/models"
	"github.com/adisnuhic/go-graphql/internal/repositories"
	"github.com/adisnuhic/go-graphql/pkg/log"
	"github.com/rs/xid"
)

const (
	// TokenTypeRefreshToken - refresh token
	TokenTypeRefreshToken = 1

	// RefreshTokenDuration holds duration value in minutes for refresh token
	RefreshTokenDuration = 43200 // 30 days
)

// ITokenService represents token service contract
type ITokenService interface {
	CreateRefreshToken(userID uint64) (*models.Token, error)
	GetByToken(token string) (*models.Token, error)
}

type tokenService struct {
	Logger     log.ILogger
	Repository repositories.ITokenRepository
}

// NewTokenService creates new token service
func NewTokenService(logger log.ILogger, repository repositories.ITokenRepository) ITokenService {
	return &tokenService{
		Logger:     logger,
		Repository: repository,
	}
}

// CreateRefreshToken creates new refresh token
func (s *tokenService) CreateRefreshToken(userID uint64) (*models.Token, error) {

	expireAt := time.Now().Add(time.Minute * time.Duration(RefreshTokenDuration))
	token := &models.Token{}
	token.UserID = userID
	token.TokenTypeID = TokenTypeRefreshToken
	token.Token = xid.New().String()
	token.ExpiresAt = expireAt

	return s.Repository.Create(token)
}

// GetByToken returns token
func (s *tokenService) GetByToken(token string) (*models.Token, error) {
	return s.Repository.GetByToken(token)
}
