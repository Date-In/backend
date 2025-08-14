package recommendations

import (
	"dating_service/configs"
	"dating_service/pkg/middleware"
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"errors"
	"net/http"
	"strconv"
)

type RecommendationHandler struct {
	service *RecommendationService
	conf    configs.Config
}

func NewRecommendationHandler(router *http.ServeMux, service *RecommendationService, config *configs.Config) {
	handler := &RecommendationHandler{service: service, conf: *config}
	router.Handle("GET /recommendations", middleware.IsAuthed(handler.GetRecommendations(), handler.conf))
}

// GetRecommendations godoc
// @Summary      Получение списка рекомендаций
// @Description  Возвращает отсортированный по релевантности (match score) список пользователей. Параметры пагинации передаются через query-параметры в URL.
// @Tags         Recommendations
// @Security     ApiKeyAuth
// @Produce      json
// @Param        page      query     int  false  "Номер страницы. По умолчанию: 1"
// @Param        pageSize  query     int  false  "Количество элементов на странице. По умолчанию: 20"
// @Success      200       {array}   ScoredUser            "Успешный ответ со списком рекомендованных пользователей"
// @Failure      400       {object}  map[string]string     "Некорректный формат параметров пагинации"
// @Failure      404       {object}  map[string]string     "Текущий пользователь или его фильтр не найден"
// @Failure      500       {object}  map[string]string     "Внутренняя ошибка сервера"
// @Router       /recommendations [get]
func (handler *RecommendationHandler) GetRecommendations() http.HandlerFunc {
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
		res.Json(w, userRecommendation, http.StatusOK)
	}
}
