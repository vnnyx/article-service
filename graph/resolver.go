package graph

import (
	"github.com/vnnyx/article-service/graph/usecase"
	"github.com/vnnyx/article-service/graph/usecase/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthorUC  usecase.AuthorUC
	ArticleUC usecase.ArticleUC
	AuthUC    auth.AuthUC
}
