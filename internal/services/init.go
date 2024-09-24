package services

import (
	"github.com/adisnuhic/go-graphql/config"
	"github.com/adisnuhic/go-graphql/internal/repositories"

	"github.com/adisnuhic/go-graphql/pkg/log"
	"github.com/golobby/container/pkg/container"
)

var (

	// services
	authSvc         IAuthService
	authProviderSvc IAuthProviderService

	// repositories
	authProviderRepo repositories.IAuthProviderRepository
	tokenRepo        repositories.ITokenRepository
	postRepo         repositories.IPostRepository
	userRepo         repositories.IUserRepository
)

// Bind services to IoC (dependency injection) container
func Init(c container.Container, logger log.ILogger) {

	// Resolve dependencies and return concrete type of given abstractions (for repos)

	// Bind auth service
	c.Singleton(func() IAuthService {
		return NewAuthService(config.Load().JWTConf.Secret)
	})

	// authProvider repository concrete type
	c.Make(&authProviderRepo)

	// Bind auth provider service
	c.Singleton(func() IAuthProviderService {
		return NewAuthProviderService(logger, authProviderRepo)
	})

	c.Make(&authSvc)
	c.Make(&authProviderSvc)

	c.Make(&tokenRepo)
	c.Make(&postRepo)
	c.Make(&userRepo)

	// Bind token service
	c.Singleton(func() ITokenService {
		return NewTokenService(logger, tokenRepo)
	})

	// Bind post service
	c.Singleton(func() IPostService {
		return NewPostService(logger, postRepo, userRepo)
	})

	// Bind user service
	c.Singleton(func() IUserService {
		return NewUserService(logger, authSvc, authProviderSvc, userRepo, postRepo)
	})

}
