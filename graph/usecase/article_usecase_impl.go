package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/vnnyx/article-service/graph/model"
	"github.com/vnnyx/article-service/graph/model/entity"
	"github.com/vnnyx/article-service/graph/repository"
)

type ArticleUCImpl struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUC(articleRepository repository.ArticleRepository) ArticleUC {
	return &ArticleUCImpl{articleRepository: articleRepository}
}

func (uc *ArticleUCImpl) CreateArticle(ctx context.Context, article *model.NewArticle) (response *model.Article, err error) {
	logrus.SetReportCaller(true)
	id := uuid.New().String()

	err = uc.articleRepository.CreateArticle(article.AuthorID, &entity.ArticleEntity{
		ID:      id,
		Name:    article.Name,
		Content: article.Content,
	})
	if err != nil {
		logrus.Error(err)
		return response, err
	}

	response = &model.Article{
		ID:      id,
		Name:    article.Name,
		Content: article.Content,
	}

	return response, nil
}

func (uc *ArticleUCImpl) GetArticleByName(ctx context.Context, name string) (authors []*model.Author, err error) {
	logrus.SetReportCaller(true)

	got, err := uc.articleRepository.FindArticleByName(name)
	if err != nil {
		logrus.Error(err)
		return authors, err
	}

	for _, author := range got {
		authors = append(authors, author.ToAuthorDTO(author.Article))
	}

	return authors, nil
}

func (uc *ArticleUCImpl) GetArticleByID(ctx context.Context, id string) (author *model.Author, err error) {
	logrus.SetReportCaller(true)
	got, err := uc.articleRepository.FindArticleByID(id)
	if err != nil {
		logrus.Error(err)
		return author, err
	}

	return got.ToAuthorDTO(got.Article), nil
}

func (uc *ArticleUCImpl) UpdateArticle(ctx context.Context, request *model.UpdateArticle) (bool, error) {
	logrus.SetReportCaller(true)
	author, err := uc.articleRepository.FindArticleByID(request.ID)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return uc.updateNotNullField(request, author.ToOneArticleOnly())
}

func (uc *ArticleUCImpl) DeleteArticle(ctx context.Context, articleID string) (bool, error) {
	logrus.SetReportCaller(true)
	got, err := uc.articleRepository.DeleteArticle(articleID)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	return got, nil
}

func (uc *ArticleUCImpl) updateNotNullField(request *model.UpdateArticle, article *entity.ArticleEntity) (bool, error) {
	if request.Content == nil {
		request.Content = &article.Content
	}
	if request.Name == nil {
		request.Name = &article.Name
	}

	got, err := uc.articleRepository.UpdateArticle(&entity.ArticleEntity{
		ID:      request.ID,
		Name:    *request.Name,
		Content: *request.Content,
	})
	if err != nil {
		return false, err
	}

	return got, nil
}
