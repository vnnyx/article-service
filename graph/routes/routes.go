package routes

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/vnnyx/article-service/graph"
	"github.com/vnnyx/article-service/graph/usecase"
	"github.com/vnnyx/article-service/graph/usecase/auth"
)

type Route struct {
	Route     *echo.Echo
	AuthorUC  usecase.AuthorUC
	ArticleUC usecase.ArticleUC
	AuthUC    auth.AuthUC
}

func NewRoute(router *echo.Echo, authorUC usecase.AuthorUC, articleUC usecase.ArticleUC, authUC auth.AuthUC) *Route {
	return &Route{
		Route:     router,
		AuthorUC:  authorUC,
		ArticleUC: articleUC,
		AuthUC:    authUC,
	}
}

func (r *Route) InitRoute() {
	r.Route.POST("/query", echo.WrapHandler(graphQLHandler(r.AuthorUC, r.ArticleUC, r.AuthUC)))
	r.Route.GET("/", echo.WrapHandler(playgroundQLHandler("/query")))
}

func graphQLHandler(authorUC usecase.AuthorUC, articleUC usecase.ArticleUC, authUC auth.AuthUC) http.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		AuthorUC:  authorUC,
		ArticleUC: articleUC,
		AuthUC:    authUC,
	}}))

	return h.ServeHTTP
}

func playgroundQLHandler(endpoint string) http.HandlerFunc {
	playgroundHandler := playground.Handler("GraphQL playground", endpoint)

	return playgroundHandler
}
