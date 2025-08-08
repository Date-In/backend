package middleware

import (
	"dating_service/configs"
	"dating_service/pkg/JWT"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

type key string

const (
	ContextIdKey key = "ContextIdKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, conf configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnauthorized(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")

		userID, err := JWT.NewJWT(conf.SecretToken.Token).ParseToken(token)
		if err != nil {
			writeUnauthorized(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextIdKey, userID)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
