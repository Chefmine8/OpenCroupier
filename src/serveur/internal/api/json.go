package api

import (
	"encoding/json"
	"net/http"
)

func ReadJSON[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var req T
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return req, err
	}
	defer r.Body.Close()

	return req, nil
}
