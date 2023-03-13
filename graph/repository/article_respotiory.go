package repository

import "github.com/vnnyx/article-service/graph/model/entity"

type ArticleRepository interface {
	CreateArticle(authorID string, article *entity.ArticleEntity) error
	FindArticleByID(articleID string) (author *entity.AuthorEntity, err error)
	FindArticleByName(name string) (authors []*entity.AuthorEntity, err error)
	UpdateArticle(article *entity.ArticleEntity) (bool, error)
	DeleteArticle(articleID string) (bool, error)
}
