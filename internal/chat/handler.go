package chat

import (
	"dating_service/configs"
	"dating_service/pkg/JWT"
	"dating_service/pkg/res"
	"dating_service/pkg/utilits"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type HandlerWs struct {
	upgrader websocket.Upgrader
	service  ChatProvider
	conf     *configs.Config
}

func NewChatHandlerWs(router *http.ServeMux, service ChatProvider, conf *configs.Config) {
	handler := &HandlerWs{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		service: service,
		conf:    conf,
	}
	router.HandleFunc("/chat/ws", handler.ServeWs())
}

func (h *HandlerWs) ServeWs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchIDStr := r.URL.Query().Get("match_id")
		matchID, err := strconv.ParseUint(matchIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid match_id", http.StatusBadRequest)
			return
		}
		tokenStr := r.URL.Query().Get("token")
		userID, err := JWT.NewJWT(h.conf.SecretToken.Token).ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)
			return
		}
		conn, err := h.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Failed to upgrade connection: %v", err)
			return
		}
		h.service.HandleNewConnection(userID, uint(matchID), conn)
	}
}

type Handler struct {
	service ChatProvider
}

func NewChatHandler(router *http.ServeMux, service ChatProvider) {
	handler := &Handler{service: service}
	router.Handle("GET /chat/history", handler.GetHistory())
}

func (h *Handler) GetHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := utilits.GetIdContext(w, r)
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		matchID, _ := strconv.ParseUint(r.URL.Query().Get("match_id"), 10, 64)

		messages, err := h.service.GetMessageHistory(userID, uint(matchID), limit)
		if err != nil {
			if errors.Is(err, ErrForbidden) {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		res.Json(w, MessagesToMessagesDto(messages), http.StatusOK)
	}
}
