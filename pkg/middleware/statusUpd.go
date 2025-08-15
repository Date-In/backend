package middleware

import (
	"dating_service/internal/action"
	"dating_service/pkg/utilits"
	"log"
	"net/http"
)

func NewStatusUpdateMiddleware(actionService *action.ActionsService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := utilits.GetIdContext(w, r)
			go func() {
				err := actionService.Update(userId)
				if err != nil {
					log.Printf("ERROR: background user activity update failed for user %d: %v", userId, err)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
