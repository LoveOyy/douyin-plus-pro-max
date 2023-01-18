package http

import (
	"context"
	"douyin/internal/pkg/domain/service"
	"douyin/internal/pkg/domain/vo"

	"github.com/gin-gonic/gin"
)

type UserAdapter struct {
	service service.UserService
}

// func (UserAdapter)

func NewUserAdapter(service service.UserService) *UserAdapter {
	u := &UserAdapter{service: service}
	user := dyHTTPRouter.Group("/user")
	user.POST("/register/", u.Register)

	return u

}

func (userAdapter *UserAdapter) Register(c *gin.Context) {

	err := userAdapter.service.Register(context.TODO(), &vo.UserRegister{
		Username: c.Query("username"),
		Password: c.Query("password"),
	})
	if err != nil {
		panic(err)
	}

}
