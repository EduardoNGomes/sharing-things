package routes

import (
	"net/http"

	"github.com/egomes/schedule/internal/application/services/user"
	"github.com/egomes/schedule/internal/infra/http/routes/hc"
	userController "github.com/egomes/schedule/internal/infra/http/routes/user"
)

type Routes struct {
	createUserService *user.CreateUserService
}

type Services struct {
	CreateUserService *user.CreateUserService
}

func (r *Routes) New() http.Handler {
	mux := http.NewServeMux()

	hc.HcRoutes(mux)

	userController.RegisterCreateUserRoute(mux, r.createUserService)
	return mux
}

func NewRoutes(services *Services) *Routes {
	return &Routes{
		createUserService: services.CreateUserService,
	}
}
