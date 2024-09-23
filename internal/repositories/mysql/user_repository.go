package repositories

import (
	"github.com/adisnuhic/go-graphql/internal/initialize/mysql"
	models "github.com/adisnuhic/go-graphql/internal/models"
)

// IUserRepository interface
type IUserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetAll() ([]*models.User, error)
	GetByID(userID uint64) (*models.User, error)
}

type userRepository struct {
	Store mysql.Store
}

// NewMySQLUserRepository -
func NewMySQLUserRepository(store mysql.Store) IUserRepository {
	return &userRepository{
		Store: store,
	}
}

// Create a new user
func (r *userRepository) Create(user *models.User) (*models.User, error) {

	if err := r.Store.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetAll returns all users
func (r *userRepository) GetAll() ([]*models.User, error) {
	users := []*models.User{}

	if err := r.Store.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// GetByID returns user for provided ID
func (r *userRepository) GetByID(userID uint64) (*models.User, error) {
	user := &models.User{}

	if err := r.Store.Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
