package chat

import (
	"dating_service/configs"
	"dating_service/pkg/JWT"
	"dating_service/pkg/res"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

type ChatHandler struct {
	upgrader websocket.Upgrader
	hubs     map[uint]*Hub
	mu       sync.RWMutex
	repo     *ChatRepository
	conf     *configs.Config
}

func NewChatHandler(router *http.ServeMux, repo *ChatRepository, conf *configs.Config) {
	service := &ChatHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		hubs: make(map[uint]*Hub),
		mu:   sync.RWMutex{},
		repo: repo,
		conf: conf,
	}
	router.HandleFunc("/chat/ws", service.ServeWs())
	router.HandleFunc("GET /chat/history", service.GetHistory())
}

func (s *ChatHandler) ServeWs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchIDStr := r.URL.Query().Get("match_id")
		matchID, err := strconv.ParseUint(matchIDStr, 10, 64)
		if err != nil {
			log.Printf("ServeWs Error: invalid match_id param: %s", matchIDStr)
			http.Error(w, "Invalid match_id", http.StatusBadRequest)
			return
		}

		tokenStr := r.URL.Query().Get("token")
		if tokenStr == "" {
			log.Println("ServeWs Error: auth token is missing from query params")
			http.Error(w, "Auth token is missing", http.StatusUnauthorized)
			return
		}

		userID, err := JWT.NewJWT(s.conf.SecretToken.Token).ParseToken(tokenStr)
		if err != nil {
			log.Printf("ServeWs Error: invalid token: %v", err)
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)
			return
		}

		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Failed to upgrade connection: %v", err)
			return
		}

		s.mu.Lock()
		hub, ok := s.hubs[uint(matchID)]
		if !ok {
			hub = NewHub(uint(matchID), s.repo)
			s.hubs[uint(matchID)] = hub
			go hub.Run()
		}
		s.mu.Unlock()
		client := &Client{
			ID:   userID,
			Hub:  hub,
			Conn: conn,
			Send: make(chan []byte, 256),
		}
		client.Hub.register <- client

		go client.writePump()
		go client.readPump()

		log.Printf("Client %d successfully connected to hub %d", userID, matchID)
	}
}

// GetHistory godoc
// @Summary      Получить историю сообщений для чата
// @Description  Возвращает постраничный список сообщений для указанного чата (матча).
// @Tags         Chat
// @Produce      json
// @Param        match_id  query     uint  true  "ID чата (матча)"
// @Param        limit     query     int   true  "Количество сообщений для загрузки"
// @Success      200       {array}   model.Message "Успешный ответ с массивом сообщений"
// @Failure      400       {string}  string        "Неверный запрос (некорректные параметры)"
// @Failure      401       {string}  string        "Пользователь не авторизован"
// @Failure      500       {string}  string        "Внутренняя ошибка сервера"
// @Security     BearerAuth
// @Router       /chat/history [get]
func (s *ChatHandler) GetHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitStr := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		matchIDStr := r.URL.Query().Get("match_id")
		matchID, err := strconv.ParseUint(matchIDStr, 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		messages, err := s.repo.GetMessageHistory(uint(matchID), limit)
		if err != nil {
			log.Printf("Failed to get messages: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		res.Json(w, messages, http.StatusOK)
	}
}
