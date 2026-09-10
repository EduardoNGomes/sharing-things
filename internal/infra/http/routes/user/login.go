package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/egomes/sharing-things/internal/application/services/user"
	domainErrors "github.com/egomes/sharing-things/internal/domain/errors"
	routerErrors "github.com/egomes/sharing-things/internal/infra/http/routes/router-errors"
	"github.com/go-playground/validator/v10"
)

func LoginRoute(mux *http.ServeMux, loginService *user.LoginService) {
	mux.HandleFunc("POST /signin", func(w http.ResponseWriter, r *http.Request) { LoginHandler(w, r, loginService) })

}

func LoginHandler(w http.ResponseWriter, r *http.Request, loginService *user.LoginService) {
	w.Header().Set("Content-Type", "application/json")
	type body struct {
		Email    string `json:"email" validate:"required,email,max=50"`
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

	loginServiceDTO := user.LoginServiceDTO{
		Email:    b.Email,
		Password: b.Password,
	}

	reply, err := loginService.Handler(r.Context(), loginServiceDTO)

	if err != nil {
		switch err {
		case domainErrors.UserNotFoundError:
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprint(domainErrors.UserNotFoundError.Error())})
		case domainErrors.InvalidPasswordError:
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"message": "E-mail or password is invalid"})
		default:
			routerErrors.InternalServerError(w, "/login", err)
		}
		return
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   reply.Token,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24 * 7),
		MaxAge:  int(time.Hour * 24 * 7),
		Secure:  true,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)

}
