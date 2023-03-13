package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/vnnyx/article-service/graph/model"
	"github.com/vnnyx/article-service/graph/model/entity"
	"github.com/vnnyx/article-service/graph/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthorUCImpl struct {
	authorRepository repository.AuthorRepository
}

func NewAuthorUC(authorRepository repository.AuthorRepository) AuthorUC {
	return &AuthorUCImpl{
		authorRepository: authorRepository,
	}
}

func (uc *AuthorUCImpl) CreateAuthor(ctx context.Context, author *model.NewAuthor) (response *model.Author, err error) {
	logrus.SetReportCaller(true)
	id := uuid.New().String()
	password, err := bcrypt.GenerateFromPassword([]byte(author.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err)
		return response, err
	}

	err = uc.authorRepository.InsertAuthor(&entity.AuthorEntity{
		ID:       id,
		Username: author.Username,
		Email:    author.Email,
		Password: string(password),
		Article:  nil,
	})

	if err != nil {
		logrus.Error(err)
		if strings.Contains(err.Error(), "E11000") {
			return response, errors.New("duplicate entry")
		}
		return response, err
	}

	response = author.ToDTO(id)
	return response, nil
}

func (uc *AuthorUCImpl) GetAllAuthor(ctx context.Context) (authors []*model.Author, err error) {
	logrus.SetReportCaller(true)

	got, err := uc.authorRepository.FindAllAuthor()
	if err != nil {
		logrus.Error(err)
	}

	for _, author := range got {
		authors = append(authors, author.ToAuthorDTO(author.Article))
	}

	return authors, nil
}

func (uc *AuthorUCImpl) GetAuthorByID(ctx context.Context, id string) (author *model.Author, err error) {
	logrus.SetReportCaller(true)

	got, err := uc.authorRepository.FindAuthorByID(id)
	if err != nil {
		logrus.Error(err)
		return author, err
	}

	author = got.ToAuthorDTO(got.Article)

	return author, nil
}

func (uc *AuthorUCImpl) GetAuthorByUsername(ctx context.Context, username string) (authors []*model.Author, err error) {
	logrus.SetReportCaller(true)

	got, err := uc.authorRepository.FindAuthorByUsername(username)
	if err != nil {
		logrus.Error(err)
		return authors, err
	}

	for _, author := range got {
		authors = append(authors, author.ToAuthorDTO(author.Article))
	}

	return authors, nil

}

func (uc *AuthorUCImpl) UpdateAuthor(ctx context.Context, author *model.UpdateAuthor) (bool, error) {
	logrus.SetReportCaller(true)

	got, err := uc.authorRepository.FindAuthorByID(author.ID)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return uc.updateNotNullField(author, got)
}

func (uc *AuthorUCImpl) DeleteAuthor(ctx context.Context, id string) (bool, error) {
	logrus.SetReportCaller(true)

	got, err := uc.authorRepository.DeleteAuthor(id)
	if err != nil {
		logrus.Error(err)
		return got, err
	}

	return got, nil
}

func (uc *AuthorUCImpl) updateNotNullField(request *model.UpdateAuthor, author *entity.AuthorEntity) (bool, error) {
	if request.Email == nil {
		request.Email = &author.Email
	}
	if request.Username == nil {
		request.Username = &author.Username
	}

	got, err := uc.authorRepository.UpdateAuthor(&entity.AuthorEntity{
		ID:       request.ID,
		Username: *request.Username,
		Email:    *request.Email,
		Password: author.Password,
		Article:  author.Article,
	})
	if err != nil {
		return false, err
	}

	return got, nil
}

func (uc *AuthorUCImpl) UpdatePassword(ctx context.Context, updatePassword *model.UpdatePassword) (bool, error) {
	logrus.SetReportCaller(true)

	newPassword, err := bcrypt.GenerateFromPassword([]byte(updatePassword.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	got, err := uc.authorRepository.UpdatePassword(updatePassword.ID, string(newPassword))
	if err != nil {
		logrus.Error(err)
		if strings.Contains(err.Error(), "E11000") {
			return false, errors.New("duplicate entry")
		}
		return false, err
	}

	return got, nil
}
