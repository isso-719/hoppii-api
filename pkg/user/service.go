package user

//go:generate mockgen -source=service.go -destination=service_mock.go -package=user

import (
	"github.com/isso-719/hoppii-api/pkg/networking"
	"net/http"
)

type IFUserService interface {
	// CreateUser : User を作成
	CreateUser(input *UserInput) (*UserOutput, error)
}

type ifPrivateUserService interface {
	// login : Hoppii にログインを行い、Cookie を取得する
	login(username, password string) (*UserCookie, error)
}

type UserService struct {
	networkingService networking.IFNetworkingService

	privateService ifPrivateUserService
}

type UserInput struct {
	StudentId string
	Password  string
}

type UserOutput struct {
	StudentId string
	Password  string
	Cookie    *UserCookie
}

type UserCookie struct {
	Cookie http.CookieJar
}

func NewUserService() *UserService {
	srv := &UserService{
		networkingService: networking.NewNetworkingService(),
	}
	srv.privateService = srv

	return srv
}
