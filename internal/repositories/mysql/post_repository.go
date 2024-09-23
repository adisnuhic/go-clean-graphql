package repositories

import (
	"github.com/adisnuhic/go-graphql/internal/initialize/mysql"
	models "github.com/adisnuhic/go-graphql/internal/models"
)

// IPostRepository interface
type IPostRepository interface {
	Create(post *models.Post) (*models.Post, error)
	GetAll() ([]*models.Post, error)
}

type postRepository struct {
	Store mysql.Store
}

// NewMySQLPostRepository -
func NewMySQLPostRepository(store mysql.Store) IPostRepository {
	return &postRepository{
		Store: store,
	}
}

// Create a new post
func (r *postRepository) Create(post *models.Post) (*models.Post, error) {

	if err := r.Store.Create(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

// GetAll returns all posts
func (r *postRepository) GetAll() ([]*models.Post, error) {

	posts := []*models.Post{}

	if err := r.Store.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
