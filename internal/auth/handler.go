package auth

import (
	"dating_service/pkg/req"
	"dating_service/pkg/utilits"
	"errors"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(router *http.ServeMux, service *Service) {
	handler := &Handler{service}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

// Register godoc
// @Title        Регистрация нового пользователя
// @Description  Создает нового пользователя и возвращает JWT токен для доступа к защищенным ресурсам
// @Param        credentials body RegisterRequestDto true "Данные для регистрации"
// @Success      201 {string} string "JWT токен"
// @Resource     Authentication
// @Route        /auth/register [post]
func (h *Handler) Register() http.HandlerFunc {
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
		token, err := h.service.Register(normalPhoneNumber, body.Name, body.Password, body.SexID, body.Age)
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
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(*token))
		w.WriteHeader(http.StatusCreated)
	}
}

// Login godoc
// @Title        Вход пользователя в систему
// @Description  Аутентифицирует пользователя и возвращает JWT токен в теле ответа как обычный текст (plain text)
// @Param        credentials body LoginRequestDto true "Данные для входа"
// @Success      200 {string} string "eyJhbGciOiJIU..."
// @Resource     Authentication
// @Route        /auth/login [post]
func (h *Handler) Login() http.HandlerFunc {
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
		token, err := h.service.Login(normalPhoneNumber, body.Password)
		if err != nil {
			switch {
			case errors.Is(err, ErrIncorrectPasswordOrPhone):
				http.Error(w, err.Error(), http.StatusConflict)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			}
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(*token))
		w.WriteHeader(http.StatusOK)
	}
}
