// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Auth struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

type Author struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Articles []*Article `json:"articles"`
}

type DeleteArticle struct {
	ID string `json:"id"`
}

type DeleteAuthor struct {
	ID string `json:"id"`
}

type GetArticleByAuthorID struct {
	ID string `json:"id"`
}

type GetArticleByID struct {
	ID string `json:"id"`
}

type GetArticleByName struct {
	Name string `json:"name"`
}

type GetArticleByUsername struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewArticle struct {
	AuthorID string `json:"author_id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
}

type NewAuthor struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateArticle struct {
	ID      string  `json:"id"`
	Name    *string `json:"name"`
	Content *string `json:"content"`
}

type UpdateAuthor struct {
	ID       string  `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

type UpdatePassword struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
