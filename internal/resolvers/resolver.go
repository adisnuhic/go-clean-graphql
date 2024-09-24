package resolvers

import (
	"github.com/adisnuhic/go-graphql/graph"
	models "github.com/adisnuhic/go-graphql/internal/models"
	"github.com/adisnuhic/go-graphql/internal/services"
)

// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	Users         []*models.User
	Posts         []*models.Post
	AuthProviders []*models.AuthProvider
	Tokens        []*models.Token

	// Services
	AuthProviderService services.IAuthProviderService
	TokenService        services.ITokenService
	UserService         services.IUserService
	PostService         services.IPostService
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// Field level resolvers
type userResolver struct{ *Resolver }

func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type postResolver struct{ *Resolver }

func (r *Resolver) Post() graph.PostResolver { return &postResolver{r} }
