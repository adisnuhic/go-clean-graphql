package resolvers

import (
	"github.com/adisnuhic/go-graphql/internal/services"
	"github.com/golobby/container/pkg/container"
)

var (
	AuthProviderSvc services.IAuthProviderService
	TokenSvc        services.ITokenService
	UserSvc         services.IUserService
	PostSvc         services.IPostService
)

func Init(c container.Container) {
	// Resolve dependencies and return concrete type of given abstractions (for services)
	c.Make(&AuthProviderSvc)
	c.Make(&TokenSvc)
	c.Make(&UserSvc)
	c.Make(&PostSvc)
}
