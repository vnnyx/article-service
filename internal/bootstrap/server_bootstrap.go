package bootstrap

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/vnnyx/article-service/exception"
	"github.com/vnnyx/article-service/graph/di"
	"github.com/vnnyx/article-service/internal/infrastructure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartServer() {
	config := infrastructure.NewConfig()

	authConn, err := grpc.Dial(config.GRPCHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	exception.PanicIfNeeded(err)
	e := echo.New()
	routes, err := di.InitializeRoute(e, authConn)
	exception.PanicIfNeeded(err)
	routes.InitRoute()

	if err := e.Start(fmt.Sprintf(":%s", config.AppPort)); err != nil {
		exception.PanicIfNeeded(err)
	}
}
