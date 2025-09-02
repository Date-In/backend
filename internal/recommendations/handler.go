package recommendations

import (
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"errors"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewRecommendationHandler(router *http.ServeMux, service *Service) {
	handler := &Handler{service: service}
	router.Handle("GET /recommendations", handler.GetRecommendations())
}

// GetRecommendations godoc
// @Title        Получение списка рекомендаций
// @Description  Возвращает отсортированный по релевантности (match score) список пользователей. Параметры пагинации передаются через query-параметры в URL.
// @Param        page query int false "Номер страницы. По умолчанию: 1"
// @Param        pageSize query int false "Количество элементов на странице. По умолчанию: 20"
// @Success      200 {object} GetRecommendationsRes "Успешный ответ со списком рекомендованных пользователей"
// @Failure      400 {string} string "Некорректный формат параметров пагинации"
// @Failure      404 {string} string "Текущий пользователь или его фильтр не найден"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Recommendations
// @Route        /recommendations [get]
func (handler *Handler) GetRecommendations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := utilits.GetIdContext(w, r)
		pageStr := r.URL.Query().Get("page")
		if pageStr == "" {
			pageStr = "0"
		}
		pageSizeStr := r.URL.Query().Get("pageSize")
		if pageSizeStr == "" {
			pageSizeStr = "0"
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, ErrQueryParam.Error(), http.StatusBadRequest)
			return
		}
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			http.Error(w, ErrQueryParam.Error(), http.StatusBadRequest)
			return
		}
		userRecommendation, err := handler.service.GetRecommendations(userId, page, pageSize)
		if err != nil {
			switch {
			case errors.Is(err, ErrUserNotFound):
				http.Error(w, err.Error(), http.StatusNotFound)
			case errors.Is(err, ErrFilterNotFound):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		resp := ScoredUserToGetRecommendationResponse(userRecommendation)
		res.Json(w, resp, http.StatusOK)
	}
}
