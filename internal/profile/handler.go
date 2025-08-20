package profile

import (
	"dating_service/pkg/req"
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ProfileHandler struct {
	service *ProfileService
}

func NewProfileHandler(router *http.ServeMux, service *ProfileService) {
	handler := &ProfileHandler{service: service}
	router.Handle("GET /profile", handler.GetInfo())
	router.Handle("PATCH /profile", handler.UpdateProfile())
	router.Handle("PUT /profile/interests", handler.UpdateInterests())
	router.Handle("POST /profile/photos", handler.AddPhoto())
	router.Handle("DELETE /profile/photo/{photoId}", handler.DeletePhoto())
	router.Handle("PATCH /profile/photo/change-avatar/{photoId}", handler.UpdateAvatar())
	router.Handle("GET /profile/avatar", handler.getAvatar())
}

// GetInfo godoc
// @Title        Получение информации о профиле
// @Description  Возвращает данные профиля текущего пользователя
// @Success      200 {object} GetInfoResponseDto "Информация о профиле"
// @Failure      404 {string} string "Пользователь не найден"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Profile
// @Route        /profile [get]
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

// UpdateProfile godoc
// @Title        Обновить данные о профиле
// @Description  Возвращает данные профиля текущего пользователя
// @Param        credentials body UpdateInfoRequestDto true "Данные для обновления"
// @Success      200 {object} GetInfoResponseDto "Информация о профиле"
// @Security     AuthorizationHeader
// @Resource     Profile
// @Route        /profile [patch]
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
			body.City,
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

// UpdateInterests godoc
// @Title        Обновить данные об интересах
// @Description  Возвращает данные профиля текущего пользователя
// @Param        credentials body UpdateInterestRequestDto true "Данные для обновления"
// @Success      200 {object} model.Interest "Список интересов"
// @Security     AuthorizationHeader
// @Resource     Profile
// @Route        /profile/interests [put]
func (handler *ProfileHandler) UpdateInterests() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := utilits.GetIdContext(w, r)
		body, err := req.HandleBody[UpdateInterestRequestDto](r)
		if err != nil {
			return
		}
		updatedInterests, err := handler.service.UpdateInterests(userID, body.InterestIDs)
		if err != nil {
			if errors.Is(err, ErrInvalidInterestId) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		res.Json(w, updatedInterests, http.StatusOK)
	}
}

// AddPhoto godoc
// @Title        Добавить фотографию в профиль
// @Description  Загружает файл фотографии для текущего пользователя. Принимает multipart/form-data с ключом "photo".
// @Param        photo formData file true "Файл фотографии для загрузки"
// @Success      201 {string} string "UUID созданной фотографии"
// @Failure      400 {string} string "Некорректный запрос (например, файл не предоставлен)"
// @Failure      401 {string} string "Пользователь не авторизован"
// @Failure      404 {string} string "Пользователь не найден"
// @Failure      409 {string} string "Достигнут лимит на количество фотографий"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Profile
// @Route        /profile/photos [post]
func (handler *ProfileHandler) AddPhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := utilits.GetIdContext(w, r)
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, ErrCannotParse.Error(), http.StatusBadRequest)
			return
		}

		file, fileHeader, err := r.FormFile("photo")
		if err != nil {
			http.Error(w, ErrNotFoundKeyPhoto.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		photoID, err := handler.service.AddPhoto(fileHeader.Filename, fileHeader.Header.Get("Content-Type"), data, userID)
		if err != nil {
			switch {
			case errors.Is(err, ErrLimitPhoto):
				http.Error(w, ErrLimitPhoto.Error(), http.StatusConflict)
			case errors.Is(err, ErrUserNotFound):
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			default:
				fmt.Println(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		res.Json(w, photoID, http.StatusCreated)
	}
}

// DeletePhoto godoc
// @Title Удаление фотографии пользователя
// @Description Удаляет фотографию, принадлежащую текущему авторизованному пользователю.
// @Param photoId path string true "UUID фотографии для удаления"
// @Success 204 {string} string "No Content - фотография успешно удалена"
// @Failure 401 string string "Пользователь не авторизован"
// @Failure 404 string string "Фотография не найдена или нет прав на удаление"
// @Failure 500 string string "Внутренняя ошибка сервера"
// @Security AuthorizationHeader
// @Resource Profile
// @Route /profile/photo/{photoId} [delete]
func (handler *ProfileHandler) DeletePhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		photoID := r.PathValue("photoId")
		userID := utilits.GetIdContext(w, r)
		err := handler.service.DeletePhoto(photoID, userID)
		if err != nil {
			if errors.Is(err, ErrPhotoNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// UpdateAvatar godoc
// @Title        Обновление аватара пользователя
// @Description  Устанавливает указанную фотографию как аватар пользователя. Только фотографии, принадлежащие пользователю, могут быть установлены как аватар.
// @Param        photoId path string true "UUID фотографии для установки как аватар"
// @Success      200 {string} string "JSON-объект с ID нового аватара"
// @Failure      401 string string "Пользователь не авторизован"
// @Failure      404 string string "Фотография не найдена или не принадлежит пользователю"
// @Failure      500 string string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Profile
// @Route        /profile/photo/change-avatar/{photoId} [patch]
func (handler *ProfileHandler) UpdateAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := utilits.GetIdContext(w, r)
		photoID := r.PathValue("photoId")
		newAvatar, err := handler.service.UpdateAvatar(photoID, userID)
		if err != nil {
			switch {
			case errors.Is(err, ErrPhotoNotFound):
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		res.Json(w, newAvatar, http.StatusOK)
	}
}

// getAvatar godoc
// @Title        Получить аватар пользователя
// @Description  Возвращает ID текущего аватара авторизованного пользователя
// @Success      200 {string} string "ID аватара пользователя"
// @Failure      404 {string} string "Not Found - аватар не установлен"
// @Failure      401 {string} string "Unauthorized - пользователь не авторизован"
// @Failure      500 {string} string "Internal Server Error - ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Profile
// @Route        /profile/avatar [get]
func (handler *ProfileHandler) getAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := utilits.GetIdContext(w, r)
		avatarId, err := handler.service.GetAvatar(userID)
		if avatarId == "" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		res.Json(w, avatarId, 200)
	}
}
