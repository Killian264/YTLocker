package handlers

import (
	"log"
	"net/http"

	"github.com/Killian264/YTLocker/golocker/services"
)

func CreateServiceInjector(s *services.Services) func(next func(w http.ResponseWriter, r *http.Request, s *services.Services) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(next func(w http.ResponseWriter, r *http.Request, s *services.Services) error) func(w http.ResponseWriter, r *http.Request) error {
		return func(w http.ResponseWriter, r *http.Request) error {
			return next(w, r, s)
		}
	}
}

func CreateErrorHandler(l *log.Logger) func(next func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(next func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			err := next(w, r)

			if err == nil {
				return
			}

			l.Printf("ERROR on ROUTE: '%s' \n ERROR: '%s'", r.URL, err.Error())

			w.WriteHeader(500)

			w.Write([]byte("An Error Occurred"))

		}
	}
}
