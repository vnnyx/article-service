//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/vnnyx/article-service/graph/repository"
	"github.com/vnnyx/article-service/graph/routes"
	"github.com/vnnyx/article-service/graph/usecase"
	"github.com/vnnyx/article-service/graph/usecase/auth"
	"github.com/vnnyx/article-service/internal/infrastructure"
	authPB "github.com/vnnyx/auth-service/pb/auth"
	"google.golang.org/grpc"
)

func InitializeRoute(e *echo.Echo, a grpc.ClientConnInterface) (*routes.Route, error) {
	wire.Build(
		infrastructure.NewConfig,
		infrastructure.NewMongoDatabase,
		repository.NewAuthorRepository,
		repository.NewArticleRepository,
		usecase.NewAuthorUC,
		authPB.NewAuthServiceClient,
		auth.NewAuthUC,
		usecase.NewArticleUC,
		routes.NewRoute,
	)
	return nil, nil
}
