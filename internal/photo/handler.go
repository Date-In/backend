package photo

import (
	"dating_service/pkg/res"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type PhotoHandler struct {
	service *PhotoService
}

func NewPhotoHandler(router *http.ServeMux, service *PhotoService) {
	handler := &PhotoHandler{service}
	router.HandleFunc("GET /photo/", handler.GetPhoto())
	router.HandleFunc("GET /photo/{id}/all", handler.GetAllUserPhotos())
}

// GetPhoto godoc
// @Title        Получить файл фотографии
// @Description  Возвращает бинарные данные фотографии по её UUID
// @Param        uuid path string true "UUID фотографии"
// @Success      200 {string} binary "Бинарные данные файла фотографии" // <--- ИЗМЕНЕНИЕ ЗДЕСЬ
// @Failure      400 {string} string "Некорректный запрос (UUID не указан)"
// @Failure      404 {string} string "Фотография не найдена"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Photo
// @Route        /photo/{uuid} [get]
func (h *PhotoHandler) GetPhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := strings.TrimPrefix(r.URL.Path, "/photo/")
		if uuid == "" {
			http.Error(w, "Photo UUID is required", http.StatusBadRequest)
			return
		}
		photo, err := h.service.GetPhoto(uuid)
		if err != nil {
			if errors.Is(err, ErrPhotoNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", photo.FileType)
		w.Header().Set("Content-Length", strconv.Itoa(len(photo.Data)))
		w.Write(photo.Data)
	}
}

// GetAllUserPhotos godoc
// @Title        Получить все ссылки на фото пользователя
// @Description  Возвращает JSON-массив со строками-ссылками на все фотографии пользователя.
// @Param        id path int true "ID Пользователя"
// @Success      200 {array} string "Массив ссылок на фотографии"
// @Failure      400 {string} string "Некорректный ID пользователя"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Photo
// @Route        /photo/{id}/all [get]
func (h *PhotoHandler) GetAllUserPhotos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdString := r.PathValue("id")
		userId, err := strconv.ParseUint(userIdString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		urls, err := h.service.GetUserPhotoURLs(uint(userId))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		res.Json(w, urls, http.StatusOK)
	}
}
