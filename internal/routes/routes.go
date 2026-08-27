package routes

import (
	"net/http"

	"github.com/egomes/schedule/internal/hc"
)

func New() http.Handler {
	mux := http.NewServeMux()

	hc.HcRoutes(mux)

	return mux
}
