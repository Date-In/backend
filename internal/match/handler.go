package match

import (
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"net/http"
)

type MatchHandler struct {
	service *MatchService
}

func NewMatchHandler(router *http.ServeMux, service *MatchService) {
	handler := &MatchHandler{service}
	router.HandleFunc("GET /matches/all", handler.GetAll())
}

// GetAll godoc
// @Title        Получить все мэтчи пользователя
// @Description  Возвращает список всех мэтчей, в которых участвует текущий авторизованный пользователь.
// @Success      200 {object} GetAllDto "Успешный ответ со списком мэтчей"
// @Failure      401 {string} string "Пользователь не авторизован"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Matches
// @Route        /matches/all [get]
func (handler *MatchHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := utilits.GetIdContext(w, r)
		matches, err := handler.service.GetAll(userId)
		if err != nil {
			switch {
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		matchesResponse := make([]MatchDto, 0, len(matches))
		for _, match := range matches {
			matchesResponse = append(matchesResponse, MatchDto{
				ID:      match.ID,
				User1ID: match.User1ID,
				User2ID: match.User2ID,
			})
		}
		res.Json(w, GetAllDto{
			Matches: matchesResponse,
		}, http.StatusOK)
	}
}
