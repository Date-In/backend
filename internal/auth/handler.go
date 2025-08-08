package auth

import (
	"dating_service/pkg/req"
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"errors"
	"net/http"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(router *http.ServeMux, service *AuthService) {
	handler := &AuthHandler{service}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequestDto](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		normalPhoneNumber, err := utilits.FormatPhoneNumber(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := handler.service.Register(normalPhoneNumber, body.Name, body.Password, body.SexID, body.Age)
		if err != nil {
			switch {
			case errors.Is(err, ErrUserAlreadyExists):
				http.Error(w, err.Error(), http.StatusConflict)
			case errors.Is(err, ErrInvalidSexID):
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		res.Json(w, token, http.StatusCreated)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequestDto](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		normalPhoneNumber, err := utilits.FormatPhoneNumber(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		token, err := handler.service.Login(normalPhoneNumber, body.Password)
		if err != nil {
			switch {
			case errors.Is(err, ErrIncorrectPasswordOrPhone):
				http.Error(w, err.Error(), http.StatusConflict)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			}
			return
		}
		res.Json(w, token, http.StatusOK)
	}
}
