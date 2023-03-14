package auth

import (
	"context"

	"github.com/vnnyx/article-service/graph/model"
)

type AuthUC interface {
	Login(ctx context.Context, req *model.LoginRequest) (res *model.Auth, err error)
}
