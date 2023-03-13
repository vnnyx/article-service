package usecase

import (
	"context"

	"github.com/vnnyx/article-service/graph/model"
)

type AuthorUC interface {
	CreateAuthor(ctx context.Context, author *model.NewAuthor) (response *model.Author, err error)
	GetAllAuthor(ctx context.Context) (authors []*model.Author, err error)
	GetAuthorByID(ctx context.Context, id string) (author *model.Author, err error)
	GetAuthorByUsername(ctx context.Context, username string) (authors []*model.Author, err error)
	UpdateAuthor(ctx context.Context, author *model.UpdateAuthor) (bool, error)
	DeleteAuthor(ctx context.Context, id string) (bool, error)
	UpdatePassword(ctx context.Context, updatePassword *model.UpdatePassword) (bool, error)
}
