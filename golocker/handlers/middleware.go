package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Killian264/YTLocker/golocker/services"
)

type ServiceHandler func(w http.ResponseWriter, r *http.Request, s *services.Services) Response
type ErrorHandler func(w http.ResponseWriter, r *http.Request) Response
type Handler func(w http.ResponseWriter, r *http.Request)

// CreateServiceInjector returns a route wrapper that injects services
func CreateServiceInjector(s *services.Services) func(next ServiceHandler) ErrorHandler {
	return func(next ServiceHandler) ErrorHandler {
		return func(w http.ResponseWriter, r *http.Request) Response {
			return next(w, r, s)
		}
	}
}

// CreateErrorHandler returns a route wrapper that handles errors
func CreateResponseWriter(l *log.Logger) func(next ErrorHandler) Handler {
	return func(next ErrorHandler) Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			res := next(w, r)

			if res.Status == http.StatusInternalServerError {
				l.Printf("ERROR: - - ROUTE: %s - - ERROR: %s", r.URL, res.Message)
				res.Message = "An Error Occurred"
			}

			marshaled, err := json.Marshal(res)
			if err != nil {
				l.Printf("Failed to marshal json %v", err)
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(res.Status)
			w.Write(marshaled)

			l.Printf("%s - - %s %s %s %d %d", r.RemoteAddr, r.Method, r.URL, r.Proto, res.Status, len(marshaled))
		}
	}
}

// CreateSubscribeHandler just sets the response header and logs if an error occurs
func CreateSubscribeHandler(l *log.Logger) func(next ErrorHandler) Handler {
	return func(next ErrorHandler) Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			res := next(w, r)

			if res.Status == http.StatusInternalServerError {
				l.Printf(" ERROR:\nROUTE: '%s'\nERROR: '%s'", r.URL, res.Message)
			}

			w.WriteHeader(res.Status)
		}
	}
}
