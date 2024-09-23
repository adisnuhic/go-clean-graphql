package resolvers

import (
	"context"

	models "github.com/adisnuhic/go-graphql/internal/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input models.NewRegisterUser) (*models.User, error) {
	return r.UserService.Create(ctx, input)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UserService.GetAll(ctx)
}
