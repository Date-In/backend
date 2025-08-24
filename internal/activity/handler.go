package activity

import (
	"dating_service/configs"
	"dating_service/pkg/JWT"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type ActivityHandlerWs struct {
	upgrader websocket.Upgrader
	hub      *Hub
	conf     *configs.Config
}

func NewActivityHandlerWs(router *http.ServeMux, storage ActivityStorage, matchProvider MatchProvider, conf *configs.Config) {
	hub := NewHub(storage, matchProvider)
	go hub.Run()
	handler := &ActivityHandlerWs{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		hub:  hub,
		conf: conf,
	}
	router.HandleFunc("/activity/ws", handler.ServeWs())
}

func (h *ActivityHandlerWs) ServeWs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.URL.Query().Get("token")
		userID, err := JWT.NewJWT(h.conf.SecretToken.Token).ParseToken(tokenStr)
		if err != nil {
			log.Printf("Activity WS: Token parse error for token '%s': %v", tokenStr, err)
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)
			return
		}
		log.Printf("Activity WS: Token for user %d parsed successfully. Upgrading connection...", userID)
		conn, err := h.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Activity WS: Failed to upgrade connection for user %d: %v", userID, err)
			log.Printf("Failed to upgrade activity connection: %v", err)
			return
		}
		log.Printf("Activity WS: Connection for user %d upgraded. Registering client...", userID)
		client := &Client{
			ID:   userID,
			Hub:  h.hub,
			Conn: conn,
			Send: make(chan []byte, 256),
		}

		client.Hub.register <- client

		go client.writePump()
		go client.readPump()
	}
}
