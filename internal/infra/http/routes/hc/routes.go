package hc

import "net/http"

func HcRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running and healthy!"))
	})
}
