package repo

import (
	"context"

	"douyin/internal/pkg/domain/po"
)

type User struct {
}

func (u *User) CreateUser(ctx context.Context, args *po.User) error {

	return db.Create(args).Error

}
