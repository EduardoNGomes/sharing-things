package routererrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, routerPath string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
	fmt.Println(fmt.Sprintf("Internal server error on router: %s \nERR: %v", routerPath, err))
}
