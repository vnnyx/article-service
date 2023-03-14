package auth

import (
	"context"

	"github.com/vnnyx/article-service/graph/model"
	"github.com/vnnyx/article-service/graph/repository"
	authPB "github.com/vnnyx/auth-service/pb/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthUCImpl struct {
	authorRepository repository.AuthorRepository
	authClient       authPB.AuthServiceClient
}

func NewAuthUC(authorRepository repository.AuthorRepository, authClient authPB.AuthServiceClient) AuthUC {
	return &AuthUCImpl{authorRepository: authorRepository, authClient: authClient}
}

func (uc *AuthUCImpl) Login(ctx context.Context, req *model.LoginRequest) (res *model.Auth, err error) {
	users, err := uc.authorRepository.FindAuthorByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	got, err := uc.authClient.Login(ctx, &authPB.AuthRequest{
		User: &authPB.User{
			Id:       users[0].ID,
			Username: users[0].Username,
		},
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	res = &model.Auth{
		AccessToken: got.AccessToken,
	}

	return res, nil
}
