package notifier

import (
	"dating_service/configs"
	"dating_service/pkg/JWT"
	"github.com/gorilla/websocket"
	"net/http"
)

type HandlerWs struct {
	upgrader        websocket.Upgrader
	notifierService *Service
	conf            *configs.Config
}

func NewHandlerWs(
	router *http.ServeMux,
	service *Service,
	conf *configs.Config,
) {
	handler := &HandlerWs{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		notifierService: service,
		conf:            conf,
	}

	router.HandleFunc("/notifier/ws", handler.ServeHTTP())
}

func (h *HandlerWs) ServeHTTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.URL.Query().Get("token")
		userID, err := JWT.NewJWT(h.conf.SecretToken.Token).ParseToken(tokenStr)
		if err != nil {
			return
		}
		conn, err := h.upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		client := NewClient(userID, conn)
		h.notifierService.HandleUserConnect(client)
		defer func() {
			h.notifierService.HandleUserDisconnect(client)
		}()
		go client.WritePump()
		client.ReadPump()
	}
}
