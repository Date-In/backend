package utilits

import (
	"dating_service/pkg/appcontext"
	"net/http"
)

func GetIdContext(w http.ResponseWriter, r *http.Request) uint {
	id := r.Context().Value(appcontext.ContextIdKey)
	userID, ok := id.(uint)
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 0
	}
	return userID
}
