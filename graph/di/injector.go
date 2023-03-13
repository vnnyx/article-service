//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/vnnyx/article-service/graph/repository"
	"github.com/vnnyx/article-service/graph/routes"
	"github.com/vnnyx/article-service/graph/usecase"
	"github.com/vnnyx/article-service/internal/infrastructure"
)

func InitializeRoute(e *echo.Echo) *routes.Route {
	wire.Build(
		infrastructure.NewConfig,
		infrastructure.NewMongoDatabase,
		repository.NewAuthorRepository,
		repository.NewArticleRepository,
		usecase.NewAuthorUC,
		usecase.NewArticleUC,
		routes.NewRoute,
	)
	return nil
}
