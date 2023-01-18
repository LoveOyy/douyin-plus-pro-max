package service

import (
	"context"
	"douyin/internal/pkg/domain/po"
	"douyin/internal/pkg/domain/repo"
	"douyin/internal/pkg/domain/vo"
)

type User struct {
	userRepo repo.User
}

func NewUser(userRepo repo.User) *User {
	return &User{userRepo: userRepo}

}
func (u *User) Register(ctx context.Context, args *vo.UserRegister) error {
	// todo 检查用户是否存在
	// todo 用户密码MD5加密 可以加盐两边
	// todo ...

	err := u.userRepo.CreateUser(ctx, &po.User{
		Username: args.Username,
		Password: args.Password,
	})
	if err != nil {
		return err
	}
	return nil
}
