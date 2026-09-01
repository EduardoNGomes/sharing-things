package routes

import (
	"net/http"

	"github.com/egomes/schedule/internal/application/services/user"
	"github.com/egomes/schedule/internal/infra/http/routes/hc"
	userController "github.com/egomes/schedule/internal/infra/http/routes/user"
)

type Routes struct {
	userService *user.CreateUserService
}

func (r *Routes) New() http.Handler {
	mux := http.NewServeMux()

	hc.HcRoutes(mux)

	userController.RegisterCreateUserRoute(mux, r.userService)
	return mux
}

func NewRoutes(userService *user.CreateUserService) *Routes {
	return &Routes{
		userService: userService,
	}
}
