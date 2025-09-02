package match

import (
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(router *http.ServeMux, service *Service) {
	handler := &Handler{service}
	router.HandleFunc("GET /matches/all", handler.GetAll())
}

// GetAll godoc
// @Title        Получить все мэтчи пользователя
// @Description  Возвращает список всех мэтчей, в которых участвует текущий авторизованный пользователь.
// @Success      200 {array} MatchPreviewDTO "Успешный ответ со списком мэтчей"
// @Failure      401 {string} string "Пользователь не авторизован"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Matches
// @Route        /matches/all [get]
func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := utilits.GetIdContext(w, r)
		matches, err := h.service.GetUserMatches(userId)
		if err != nil {
			switch {
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		response := ToMatchPreviewDTOs(matches, userId)
		res.Json(w, response, http.StatusOK)
	}
}
