package middleware

import (
	"context"
	"dating_service/configs"
	"dating_service/pkg/JWT"
	"net/http"
	"strings"
)

type key string

const (
	ContextPhoneKey key = "ContextPhoneKey"
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
		isValid, data := JWT.NewJWT(conf.SecretToken.Token).ParseToken(token)
		if !isValid {
			writeUnauthorized(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextPhoneKey, data.Phone)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
