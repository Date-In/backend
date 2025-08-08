package profile

import (
	"dating_service/configs"
	"dating_service/pkg/middleware"
	"dating_service/pkg/res"
	"errors"
	"fmt"
	"net/http"
)

type ProfileHandler struct {
	service *ProfileService
	config  configs.Config
}

func NewProfileHandler(router *http.ServeMux, service *ProfileService, config *configs.Config) {
	handler := &ProfileHandler{service: service, config: *config}
	router.Handle("GET /profile", middleware.IsAuthed(handler.GetInfo(), handler.config))
}

func (handler *ProfileHandler) GetInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(middleware.ContextIdKey)
		userID, ok := id.(uint)
		if !ok {
			fmt.Println(ok)
			fmt.Println(userID)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		user, err := handler.service.GetInfo(userID)
		if err != nil {
			switch {
			case errors.Is(err, ErrUserNotFound):
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		res.Json(w, ToProfileResponseDto(user), http.StatusOK)
	}
}
