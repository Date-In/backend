package profile

import (
	"dating_service/configs"
	"dating_service/pkg/middleware"
	"dating_service/pkg/req"
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"errors"
	"net/http"
)

type ProfileHandler struct {
	service *ProfileService
	config  configs.Config
}

func NewProfileHandler(router *http.ServeMux, service *ProfileService, config *configs.Config) {
	handler := &ProfileHandler{service: service, config: *config}
	router.Handle("GET /profile", middleware.IsAuthed(handler.GetInfo(), handler.config))
	router.Handle("PATCH /profile", middleware.IsAuthed(handler.UpdateProfile(), handler.config))
}

func (handler *ProfileHandler) GetInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := utilits.GetIdContext(w, r)
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

func (handler *ProfileHandler) UpdateProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := utilits.GetIdContext(w, r)
		body, err := req.HandleBody[UpdateInfoRequestDto](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedUser, err := handler.service.Update(userID,
			body.Name,
			body.Age,
			body.Bio,
			body.Children,
			body.Height,
			body.SexId,
			body.ZodiacSignId,
			body.WorldviewId,
			body.TypeOfDatingId,
			body.EducationId,
			body.AttitudeToAlcoholId,
			body.AttitudeToSmokingId)
		if err != nil {
			switch {
			case errors.Is(err, ErrUserNotFound):
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			case errors.Is(err, ErrInvalidSexID):
				http.Error(w, err.Error(), http.StatusBadRequest)
			case errors.Is(err, ErrInvalidTypeOfDatingId):
				http.Error(w, err.Error(), http.StatusBadRequest)
			case errors.Is(err, ErrInvalidEducationId):
				http.Error(w, err.Error(), http.StatusBadRequest)
			case errors.Is(err, ErrInvalidZodiacID):
				http.Error(w, err.Error(), http.StatusBadRequest)
			case errors.Is(err, ErrInvalidWordViewID):
				http.Error(w, err.Error(), http.StatusBadRequest)
			case errors.Is(err, ErrInvalidAttitudeToAlcoholicId):
				http.Error(w, err.Error(), http.StatusBadRequest)
			case errors.Is(err, ErrInvalidAttitudeToSmokingId):
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		res.Json(w, ToProfileResponseDto(updatedUser), http.StatusOK)
	}
}
