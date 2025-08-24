package chat

import (
	"dating_service/configs"
	"dating_service/pkg/JWT"
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

type ChatHandlerWs struct {
	upgrader      websocket.Upgrader
	hubs          map[uint]*Hub
	mu            sync.RWMutex
	chatStorage   ChatStorage
	matchProvider MatchProvider
	conf          *configs.Config
}

type ChatHandler struct {
	chatStorage   ChatStorage
	matchProvider MatchProvider
}

func NewChatHandler(router *http.ServeMux, chatStorage ChatStorage, matchProvider MatchProvider) {
	handler := &ChatHandler{
		chatStorage:   chatStorage,
		matchProvider: matchProvider,
	}
	router.Handle("GET /chat/history", handler.GetHistory())
}

func NewChatHandlerWs(router *http.ServeMux, chatStorage ChatStorage, matchProvider MatchProvider, conf *configs.Config) {
	service := &ChatHandlerWs{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		hubs:          make(map[uint]*Hub),
		mu:            sync.RWMutex{},
		chatStorage:   chatStorage,
		matchProvider: matchProvider,
		conf:          conf,
	}
	router.HandleFunc("/chat/ws", service.ServeWs())
}

func (s *ChatHandlerWs) ServeWs() http.HandlerFunc {
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

		isParticipant, err := s.matchProvider.IsUserInMatch(userID, uint(matchID))
		if err != nil {
			log.Printf("ServeWs Error: failed to check match participation for user %d in match %d: %v", userID, matchID, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if !isParticipant {
			log.Printf("ServeWs Forbidden: user %d tried to access chat for match %d", userID, matchID)
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
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
			hub = NewHub(uint(matchID), s.chatStorage)
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
// @Title        Получить историю сообщений для чата
// @Description  Возвращает постраничный список сообщений для указанного чата (матча).
// @Param        match_id query uint true "ID чата (матча)"
// @Param        limit query int true "Количество сообщений для загрузки"
// @Success      200 {array} string "Успешный ответ с массивом сообщений"
// @Failure      400 {string} string "Неверный запрос (некорректные параметры)"
// @Failure      401 {string} string "Пользователь не авторизован"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Security     AuthorizationHeader
// @Resource     Chat
// @Route        /chat/history [get]
func (s *ChatHandler) GetHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := utilits.GetIdContext(w, r)
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

		isParticipant, err := s.matchProvider.IsUserInMatch(userID, uint(matchID))
		if err != nil {
			log.Printf("GetHistory Error: failed to check match participation for user %d in match %d: %v", userID, matchID, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if !isParticipant {
			log.Printf("GetHistory Forbidden: user %d tried to access history for match %d", userID, matchID)
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		messages, err := s.chatStorage.GetMessageHistory(uint(matchID), limit)
		if err != nil {
			log.Printf("Failed to get messages: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		res.Json(w, messages, http.StatusOK)
	}
}
