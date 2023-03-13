package repository

import (
	"github.com/vnnyx/article-service/graph/model/entity"
)

type AuthorRepository interface {
	InsertAuthor(author *entity.AuthorEntity) error
	FindAllAuthor() (authors []*entity.AuthorEntity, err error)
	FindAuthorByID(id string) (author *entity.AuthorEntity, err error)
	FindAuthorByUsername(username string) (authors []*entity.AuthorEntity, err error)
	UpdateAuthor(author *entity.AuthorEntity) (bool, error)
	DeleteAuthor(id string) (bool, error)
	UpdatePassword(id, newPassword string) (bool, error)
}
