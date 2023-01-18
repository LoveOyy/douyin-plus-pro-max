package douyin

import (
	"douyin/internal/pkg/adapter/http"
	"douyin/internal/pkg/usecase/repo"
	"douyin/internal/pkg/usecase/service"
)

func Init() {

	repo.Init()
	http.Init()

}

func Run() {
	var userRepo = &repo.User{}
	var userService = service.NewUser(userRepo)
	http.NewUserAdapter(userService)

	http.Run()
}
