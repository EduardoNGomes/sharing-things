package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/egomes/schedule/internal/application/services/user"
	domainErrors "github.com/egomes/schedule/internal/domain/errors"
	routerErrors "github.com/egomes/schedule/internal/infra/http/routes/router-errors"
	"github.com/go-playground/validator/v10"
)

func RegisterCreateUserRoute(mux *http.ServeMux, createUserService *user.CreateUserService) {
	mux.HandleFunc("POST /signup", func(w http.ResponseWriter, r *http.Request) { createUserHandler(w, r, createUserService) })

}

func createUserHandler(w http.ResponseWriter, r *http.Request, createUserService *user.CreateUserService) {
	w.Header().Set("Content-Type", "application/json")
	type body struct {
		Email    string `json:"email" validate:"required,email,max=50"`
		Name     string `json:"name" validate:"required,min=2,max=50"`
		Password string `json:"password" validate:"required,min=6,max=50"`
	}

	var b body

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request"})
		return
	}

	if err := validator.New().Struct(b); err != nil {
		var validationErros validator.ValidationErrors

		w.WriteHeader(http.StatusBadRequest)

		if errors.As(err, &validationErros) {
			fieldErro := validationErros[0]

			message := fmt.Sprintf("Invalid '%s' field\n Reason %s\n", fieldErro.Field(), fieldErro.Tag())

			json.NewEncoder(w).Encode(map[string]string{"message": message})
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request"})
		return
	}

	createUserServiceDTO := user.CreateUserServiceDTO{
		Email:    b.Email,
		Name:     b.Name,
		Password: b.Password,
	}

	if err := createUserService.Handler(r.Context(), createUserServiceDTO); err != nil {
		switch err {
		case domainErrors.UserAlreadyExistsError:
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprint(domainErrors.UserAlreadyExistsError.Error())})
		default:
			routerErrors.InternalServerError(w, "/signup", err)
		}
		return

	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("User %s created successfully ", b.Email)})
}
