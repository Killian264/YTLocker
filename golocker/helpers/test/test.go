package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/gorilla/mux"
)

// FakeRequest is used to send a fake request
type FakeRequest struct {
	Services *services.Services
	Request  *http.Request
	Route    string
	Handler  func(w http.ResponseWriter, r *http.Request, s *services.Services) error
}

// SendFakeRequest sends a fake request the request will return 500 status code on error
func SendFakeRequest(request FakeRequest) *http.Response {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	handler := func(w http.ResponseWriter, r *http.Request) {
		err := request.Handler(w, r, request.Services)

		if err != nil {
			w.WriteHeader(500)
		}
	}

	router.HandleFunc(request.Route, handler)
	router.ServeHTTP(rr, request.Request)
	return rr.Result()
}
