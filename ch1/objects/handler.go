package objects

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	switch m {
	case http.MethodGet:
		get(w, r)
		return
	case http.MethodPut:
		put(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
