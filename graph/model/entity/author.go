package entity

import "github.com/vnnyx/article-service/graph/model"

type AuthorEntity struct {
	ID       string           `bson:"_id"`
	Username string           `bson:"username"`
	Email    string           `bson:"email"`
	Password string           `bson:"password"`
	Article  []*ArticleEntity `bson:"articles"`
}

func (a *AuthorEntity) ToOneArticleOnly() *ArticleEntity {
	return &ArticleEntity{
		ID:      a.Article[0].ID,
		Name:    a.Article[0].Name,
		Content: a.Article[0].Content,
	}
}

func (a *AuthorEntity) ToAuthorDTO(ae []*ArticleEntity) *model.Author {
	var ar = make([]*model.Article, 0)
	for _, article := range ae {
		ar = append(ar, article.ToArticleDTO())
	}
	return &model.Author{
		ID:       a.ID,
		Username: a.Username,
		Email:    a.Email,
		Articles: ar,
	}
}
