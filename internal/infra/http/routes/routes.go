package routes

import (
	"net/http"

	"github.com/egomes/schedule/internal/infra/http/routes/hc"
)

func New() http.Handler {
	mux := http.NewServeMux()

	hc.HcRoutes(mux)

	return mux
}
