package services

import (
	"context"

	models "github.com/adisnuhic/go-graphql/internal/models"
	"github.com/adisnuhic/go-graphql/internal/repositories"
	"github.com/adisnuhic/go-graphql/pkg/log"
)

// IPostService represents post service contract
type IPostService interface {
	Create(ctx context.Context, input models.NewPost) (*models.Post, error)
	GetAll(ctx context.Context) ([]*models.Post, error)
}

type postService struct {
	Logger         log.ILogger
	Repository     repositories.IPostRepository
	UserRepository repositories.IUserRepository
}

// NewPostService creates new post service
func NewPostService(logger log.ILogger, repository repositories.IPostRepository, userRepo repositories.IUserRepository) IPostService {
	return &postService{
		Logger:         logger,
		Repository:     repository,
		UserRepository: userRepo,
	}
}

// Create creates new post
func (s *postService) Create(ctx context.Context, input models.NewPost) (*models.Post, error) {

	// check if user exists
	_, errUser := s.UserRepository.GetByID(input.UserID)
	if errUser != nil {
		return nil, errUser
	}

	post := &models.Post{
		Title:       input.Title,
		Description: input.Description,
		UserID:      input.UserID,
	}

	return s.Repository.Create(post)
}

// GetPosts returns all posts
func (s *postService) GetAll(ctx context.Context) ([]*models.Post, error) {
	return s.Repository.GetAll()
}
