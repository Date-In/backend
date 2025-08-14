package middleware

import (
	"dating_service/internal/action"
	"dating_service/pkg/utilits"
	"net/http"
	"time"
)

func NewStatusUpdateMiddleware(actionService *action.ActionsService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := utilits.GetIdContext(w, r)
			err := actionService.Update(userId, time.Now())
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
