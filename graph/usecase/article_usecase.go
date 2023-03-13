package usecase

import (
	"context"

	"github.com/vnnyx/article-service/graph/model"
)

type ArticleUC interface {
	CreateArticle(ctx context.Context, article *model.NewArticle) (response *model.Article, err error)
	GetArticleByName(ctx context.Context, name string) (authors []*model.Author, err error)
	GetArticleByID(ctx context.Context, id string) (author *model.Author, err error)
	UpdateArticle(ctx context.Context, article *model.UpdateArticle) (bool, error)
	DeleteArticle(ctx context.Context, articleID string) (bool, error)
}
