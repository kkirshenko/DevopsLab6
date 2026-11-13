package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kkirshenko/DevopsLab6/internal/restaurant"
)

func DecodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var item restaurant.Restaurant
	err := json.NewDecoder(r.Body).Decode(&item)
	return createRequest{Restaurant:item}, err
}
func DecodePathRequest(_ context.Context, r *http.Request) (interface{}, error) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 || parts[0] != "restaurants" { return nil, restaurant.ErrInvalid }
	id := parts[1]
	if r.Method == http.MethodGet { return idRequest{ID:id}, nil }
	if r.Method == http.MethodPatch {
		var body struct { Rating float32 `json:"rating"` }
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil { return nil, err }
		return ratingRequest{ID:id, Rating:body.Rating}, nil
	}
	return nil, restaurant.ErrInvalid
}
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err, ok := response.(error); ok { w.WriteHeader(http.StatusBadRequest); return json.NewEncoder(w).Encode(map[string]string{"error":err.Error()}) }
	return json.NewEncoder(w).Encode(map[string]interface{}{"data":response})
}
