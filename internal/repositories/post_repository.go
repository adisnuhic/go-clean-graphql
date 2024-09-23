package repositories

import models "github.com/adisnuhic/go-graphql/internal/models"

// IPostRepository contract
type IPostRepository interface {
	Create(post *models.Post) (*models.Post, error)
	GetAll() ([]*models.Post, error)
}
