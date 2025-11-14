package httputils

import (
	"encoding/json"
	"net/http"
)

func DecodeRequest(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(&v)
}
