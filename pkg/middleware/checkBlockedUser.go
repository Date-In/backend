package middleware

import (
	"dating_service/internal/user"
	"dating_service/pkg/utilits"
	"log"
	"net/http"
)

func NewCheckBlockedUserMiddleware(userRepository *user.UserRepository) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := utilits.GetIdContext(w, r)
			status, err := userRepository.GetStatusUser(userId)
			if err != nil {
				log.Printf("Error in check block user")
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			if status == 0 {
				log.Printf("User %s hasnt status", userId)
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			if status == 3 {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
