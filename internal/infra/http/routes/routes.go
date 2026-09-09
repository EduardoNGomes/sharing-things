package routes

import (
	"net/http"

	"github.com/egomes/schedule/internal/application/services/user"
	"github.com/egomes/schedule/internal/infra/http/routes/hc"
	userController "github.com/egomes/schedule/internal/infra/http/routes/user"
)

type Routes struct {
	createUserService *user.CreateUserService
	loginService      *user.LoginService
}

type Services struct {
	CreateUserService *user.CreateUserService
	LoginService      *user.LoginService
}

func (r *Routes) New() http.Handler {
	mux := http.NewServeMux()

	hc.HcRoutes(mux)

	userController.RegisterCreateUserRoute(mux, r.createUserService)
	userController.LoginRoute(mux, r.loginService)

	return mux
}

func NewRoutes(
	services *Services) *Routes {
	return &Routes{
		createUserService: services.CreateUserService,
		loginService:      services.LoginService,
	}
}
