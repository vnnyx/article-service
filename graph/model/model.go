package model

func (a *NewArticle) ToDTO(id string) *Article {
	return &Article{
		ID:      id,
		Name:    a.Name,
		Content: a.Content,
	}
}

func (a *NewAuthor) ToDTO(id string) *Author {
	article := make([]*Article, 0)
	return &Author{
		ID:       id,
		Username: a.Username,
		Email:    a.Email,
		Articles: article,
	}
}
