package filter

import (
	"dating_service/pkg/req"
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"errors"
	"net/http"
)

type FilterHandler struct {
	service *FilterService
}

func NewFilterHandler(router *http.ServeMux, service *FilterService) {
	handler := &FilterHandler{service: service}
	router.Handle("GET /filter", handler.GetFilter())
	router.Handle("POST /filter/create", handler.CreateFilter())
	router.Handle("PATCH /filter/update", handler.UpdateFilter())
}

// CreateFilter godoc
// @Summary      Создание фильтра поиска
// @Description  Создает новый фильтр поиска для текущего аутентифицированного пользователя.
// @Tags         Filters
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        filterData   body      CreateFilterDto     true  "Данные для создания фильтра"
// @Success      201          {object}  map[string]string   "Фильтр успешно создан"
// @Failure      400          {object}  map[string]string   "Некорректные данные запроса"
// @Failure      500          {object}  map[string]string   "Внутренняя ошибка сервера"
// @Router       /filter/create [post]
func (handler *FilterHandler) CreateFilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := utilits.GetIdContext(w, r)
		body, err := req.HandleBody[CreateFilterDto](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.service.CreateFilter(userId, body.MinAge, body.MaxAge, body.SexID, body.Location)
		if err != nil {
			switch {
			case errors.Is(err, ErrMaxAndMinValue):
				http.Error(w, err.Error(), http.StatusBadRequest)
			case errors.Is(err, ErrFilterExists):
				http.Error(w, err.Error(), http.StatusConflict)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// UpdateFilter godoc
// @Summary      Обновление фильтра поиска
// @Description  Обновляет существующий фильтр поиска для текущего пользователя. Позволяет частичное обновление: можно передать только те поля, которые нужно изменить.
// @Tags         Filters
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        filterUpdateData body      UpdateFilterDto     true  "Данные для обновления фильтра (можно передавать только изменяемые поля)"
// @Success      200              {object}  map[string]string   "Фильтр успешно обновлен"
// @Failure      400              {object}  map[string]string   "Некорректные данные запроса"
// @Failure      404              {object}  map[string]string   "Фильтр не найден"
// @Failure      500              {object}  map[string]string   "Внутренняя ошибка сервера"
// @Router       /filter/update [patch]
func (handler *FilterHandler) UpdateFilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := utilits.GetIdContext(w, r)
		body, err := req.HandleBody[UpdateFilterDto](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = handler.service.UpdateUserFilter(userId, body.MinAge, body.MaxAge, body.SexId, body.Location)
		if err != nil {
			switch {
			case errors.Is(err, ErrNotFoundFilter):
				http.Error(w, err.Error(), http.StatusNotFound)
			case errors.Is(err, ErrMaxAndMinValue):
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// GetFilter godoc
// @Summary      Получение фильтра поиска
// @Description  Возвращает текущие настройки фильтра поиска для аутентифицированного пользователя.
// @Tags         Filters
// @Security     ApiKeyAuth
// @Produce      json
// @Success      200  {object}  GetFilterDto        "Успешный ответ с данными фильтра"
// @Failure      404  {object}  map[string]string   "Фильтр для данного пользователя не найден"
// @Failure      500  {object}  map[string]string   "Внутренняя ошибка сервера"
// @Router       /filter [get]
func (handler *FilterHandler) GetFilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := utilits.GetIdContext(w, r)
		filter, err := handler.service.GetFilter(userId)
		if err != nil {
			switch {
			case errors.Is(err, ErrNotFoundFilter):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		res.Json(w, GetFilterDto{
			MinAge:   filter.MinAge,
			MaxAge:   filter.MaxAge,
			SexID:    filter.SexID,
			Location: filter.Location,
		}, http.StatusOK)
	}
}
