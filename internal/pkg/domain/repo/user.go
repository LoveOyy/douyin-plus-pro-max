package repo

import (
	"context"

	"douyin/internal/pkg/domain/po"
)

type User interface {
	CreateUser(ctx context.Context, args *po.User) error
}
