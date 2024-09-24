package resolvers

import (
	"context"

	models "github.com/adisnuhic/go-graphql/internal/models"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input models.NewPost) (*models.Post, error) {
	return r.PostService.Create(ctx, input)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	return r.PostService.GetAll(ctx)
}

// User is a field level resolver for Post
func (r *postResolver) User(ctx context.Context, post *models.Post) (*models.User, error) {
	return r.PostService.GetUserByPostID(ctx, post.ID)
}
