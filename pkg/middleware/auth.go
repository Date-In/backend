package middleware

import (
	"dating_service/configs"
	"dating_service/pkg/JWT"
	"dating_service/pkg/appcontext"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func NewAuthMiddleware(conf configs.Config) Middleware {
	return func(next http.Handler) http.Handler {
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
			ctx := context.WithValue(r.Context(), appcontext.ContextIdKey, userID)
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		})
	}
}
