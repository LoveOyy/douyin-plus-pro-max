package service

import (
	"context"

	"douyin/internal/pkg/domain/vo"
)

type UserService interface {
	Register(ctx context.Context, args *vo.UserRegister) error
}
