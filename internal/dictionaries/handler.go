package dictionaries

import (
	"dating_service/pkg/res"
	"log"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(router *http.ServeMux, service *Service) {
	handler := &Handler{service}
	router.HandleFunc("GET /dict/sexes", handler.GetSexes())
	router.HandleFunc("GET /dict/educations", handler.GetEducations())
	router.HandleFunc("GET /dict/zodiac-signs", handler.GetZodiacSigns())
	router.HandleFunc("GET /dict/worldviews", handler.GetWorldViews())
	router.HandleFunc("GET /dict/type-of-dating", handler.GetTypeOfDating())
	router.HandleFunc("GET /dict/attitude-to-alcohol", handler.GetAttitudeToAlcohol())
	router.HandleFunc("GET /dict/attitude-to-smoking", handler.GetAttitudeToSmoking())
	router.HandleFunc("GET /dict/interests", handler.GetInterests())
	router.HandleFunc("GET /dict/statuses", handler.GetStatuses())
}

// GetSexes godoc
// @Title       Получить список полов
// @Description Возвращает справочник полов
// @Success     200 {array} model.Sex "Список полов"
// @Failure     500 {string} string "Ошибка сервера"
// @Resource    Dictionaries
// @Route       /dict/sexes [get]
func (h *Handler) GetSexes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetSexes()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, resp, http.StatusOK)
	}
}

// GetEducations godoc
// @Title       Получить список образований
// @Description Возвращает справочник образований
// @Success     200 {array} model.Education "Список образований"
// @Failure     500 {string} string "Ошибка сервера"
// @Resource    Dictionaries
// @Route       /dict/educations [get]
func (h *Handler) GetEducations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetEducations()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, resp, http.StatusOK)
	}
}

// GetZodiacSigns godoc
// @Title       Получить список знаков зодиака
// @Description Возвращает справочник знаков зодиака
// @Success     200 {array} model.ZodiacSign "Список знаков зодиака"
// @Failure     500 {string} string "Ошибка сервера"
// @Resource    Dictionaries
// @Route       /dict/zodiac-signs [get]
func (h *Handler) GetZodiacSigns() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetZodiacSigns()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, resp, http.StatusOK)
	}
}

// GetWorldViews godoc
// @Title       Получить список мировоззрений
// @Description Возвращает справочник мировоззрений
// @Success     200 {array} model.Worldview "Список мировоззрений"
// @Failure     500 {string} string "Ошибка сервера"
// @Resource    Dictionaries
// @Route       /dict/worldviews [get]
func (h *Handler) GetWorldViews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetWorldViews()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, resp, http.StatusOK)
	}
}

// GetTypeOfDating godoc
// @Title       Получить список типов знакомств
// @Description Возвращает справочник типов знакомств
// @Success     200 {array} model.TypeOfDating "Список типов знакомств"
// @Failure     500 {string} string "Ошибка сервера"
// @Resource    Dictionaries
// @Route       /dict/type-of-dating [get]
func (h *Handler) GetTypeOfDating() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetTypeOfDating()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, resp, http.StatusOK)
	}
}

// GetAttitudeToAlcohol godoc
// @Title       Получить список вариантов отношения к алкоголю
// @Description Возвращает справочник отношения к алкоголю
// @Success     200 {array} model.AttitudeToAlcohol "Список вариантов отношения к алкоголю"
// @Failure     500 {string} string "Ошибка сервера"
// @Resource    Dictionaries
// @Route       /dict/attitude-to-alcohol [get]
func (h *Handler) GetAttitudeToAlcohol() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetAttitudeToAlcohol()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, resp, http.StatusOK)
	}
}

// GetAttitudeToSmoking godoc
// @Title       Получить список вариантов отношения к курению
// @Description Возвращает справочник отношения к курению
// @Success     200 {array} model.AttitudeToSmoking "Список вариантов отношения к курению"
// @Failure     500 {string} string "Ошибка сервера"
// @Resource    Dictionaries
// @Route       /dict/attitude-to-smoking [get]
func (h *Handler) GetAttitudeToSmoking() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetAttitudeToSmoking()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, resp, http.StatusOK)
	}
}

// GetInterests godoc
// @Title       Получить список интересов
// @Description Возвращает справочник интересов
// @Success     200 {array} model.Interest "Список интересов"
// @Failure     500 {string} string "Ошибка сервера"
// @Resource    Dictionaries
// @Route       /dict/interests [get]
func (h *Handler) GetInterests() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetInterests()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, resp, http.StatusOK)
	}
}

// GetStatuses godoc
// @Title       Получить список статусов
// @Description Возвращает справочник статусов
// @Success     200 {array} model.Status "Список статусов"
// @Failure     500 {string} string "Ошибка сервера"
// @Resource    Dictionaries
// @Route       /dict/statuses [get]
func (h *Handler) GetStatuses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetStatuses()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Json(w, resp, http.StatusOK)
	}
}
