package bootstrap

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/vnnyx/article-service/exception"
	"github.com/vnnyx/article-service/graph/di"
	"github.com/vnnyx/article-service/internal/infrastructure"
)

func StartServer() {
	config := infrastructure.NewConfig()
	e := echo.New()
	routes := di.InitializeRoute(e)
	routes.InitRoute()

	if err := e.Start(fmt.Sprintf(":%s", config.AppPort)); err != nil {
		exception.PanicIfNeeded(err)
	}
}
