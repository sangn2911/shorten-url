package port

import "net/http"

type Handler interface {
	Encode(w http.ResponseWriter, r *http.Request)
	Decode(w http.ResponseWriter, r *http.Request)
}
