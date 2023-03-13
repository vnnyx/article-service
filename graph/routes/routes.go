package routes

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/vnnyx/article-service/graph"
	"github.com/vnnyx/article-service/graph/usecase"
)

type Route struct {
	Route     *echo.Echo
	AuthorUC  usecase.AuthorUC
	ArticleUC usecase.ArticleUC
}

func NewRoute(router *echo.Echo, authorUC usecase.AuthorUC, articleUC usecase.ArticleUC) *Route {
	return &Route{
		Route:     router,
		AuthorUC:  authorUC,
		ArticleUC: articleUC,
	}
}

func (r *Route) InitRoute() {
	r.Route.POST("/query", echo.WrapHandler(graphQLHandler(r.AuthorUC, r.ArticleUC)))
	r.Route.GET("/", echo.WrapHandler(playgroundQLHandler("/query")))
}

func graphQLHandler(authorUC usecase.AuthorUC, articleUC usecase.ArticleUC) http.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		AuthorUC:  authorUC,
		ArticleUC: articleUC,
	}}))

	return h.ServeHTTP
}

func playgroundQLHandler(endpoint string) http.HandlerFunc {
	playgroundHandler := playground.Handler("GraphQL playground", endpoint)

	return playgroundHandler
}
