package entity

import "github.com/vnnyx/article-service/graph/model"

type ArticleList struct {
	Articles []*ArticleEntity `bson:"articles"`
}

type ArticleEntity struct {
	ID      string `bson:"_id"`
	Name    string `bson:"name"`
	Content string `bson:"content"`
}

func (a *ArticleEntity) ToArticleDTO() *model.Article {
	return &model.Article{
		ID:      a.ID,
		Name:    a.Name,
		Content: a.Content,
	}
}
