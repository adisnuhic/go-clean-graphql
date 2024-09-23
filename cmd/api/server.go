package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/adisnuhic/go-graphql/config"
	"github.com/adisnuhic/go-graphql/graph"
	initLog "github.com/adisnuhic/go-graphql/internal/initialize/log"
	"github.com/adisnuhic/go-graphql/internal/initialize/mysql"
	repositories "github.com/adisnuhic/go-graphql/internal/repositories/mysql"
	"github.com/adisnuhic/go-graphql/internal/resolvers"
	"github.com/adisnuhic/go-graphql/internal/services"
	"github.com/adisnuhic/go-graphql/pkg/log"
	"github.com/golobby/container"
)

func main() {

	// init logger
	logger := log.NewZeroLogger(initLog.InitZeroLogger())

	// load app config
	cfg := config.Load()

	// initialize mysql database
	mysql.Init(cfg, logger)

	// create new DI container
	c := container.NewContainer()

	// init repositories
	repositories.Init(c)

	// init services/usecases
	services.Init(c, logger)

	// init resolver
	resolvers.Init(c)

	// server settings
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		AuthProviderService: resolvers.AuthProviderSvc,
		TokenService:        resolvers.TokenSvc,
		UserService:         resolvers.UserSvc,
		PostService:         resolvers.PostSvc,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Printf("connect to %s:%s/ for GraphQL playground", cfg.Url, cfg.Port)
	logger.Fatalf("%s", http.ListenAndServe(":"+cfg.Port, nil))
}
