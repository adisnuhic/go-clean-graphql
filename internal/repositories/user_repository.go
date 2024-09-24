package repositories

import models "github.com/adisnuhic/go-graphql/internal/models"

// IUserRepository contract
type IUserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetAll() ([]*models.User, error)
	GetByID(userID uint64) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByPostID(postID uint64) (*models.User, error)
}
