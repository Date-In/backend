package like

import (
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"errors"
	"net/http"
	"strconv"
)

type LikeHandler struct {
	service *LikeService
}

func NewLikeHandler(router *http.ServeMux, service *LikeService) {
	handler := &LikeHandler{service}
	router.HandleFunc("POST /like/{target_id}", handler.CreateLike())
	router.HandleFunc("GET /like/all", handler.GetLike())
}

// CreateLike godoc
// @Title        Создать лайк
// @Description  Текущий авторизованный пользователь ставит лайк другому пользователю. Если лайк взаимный, создается мэтч.
// @Param        target_id path int true "ID пользователя, которого нужно лайкнуть"
// @Success      201 {string} string "Created - Лайк успешно создан"
// @Failure      400 {string} string "Неверный ID пользователя в URL"
// @Failure      401 {string} string "Пользователь не авторизован"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Likes
// @Route        /like/{target_id} [post]
func (handler *LikeHandler) CreateLike() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := utilits.GetIdContext(w, r)
		targetIdStr := r.PathValue("target_id")
		targetId, err := strconv.ParseUint(targetIdStr, 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		err = handler.service.CreateLike(userId, uint(targetId))
		if err != nil {
			switch {
			case errors.Is(err, ErrNotFoundUser):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// GetLike godoc
// @Title        Получить список своих лайков
// @Description  Возвращает список всех лайков, поставленных текущим авторизованным пользователем.
// @Success      200 {array} LikeDto "Список лайков"
// @Failure      401 {string} string "Пользователь не авторизован"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Likes
// @Route        /like/all [get]
func (handler *LikeHandler) GetLike() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := utilits.GetIdContext(w, r)
		likes, err := handler.service.GetLikes(userId)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		likesResponse := LikeToDto(likes)
		res.Json(w, likesResponse, http.StatusOK)
	}
}
