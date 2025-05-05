package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) router() http.Handler {
	router := mux.NewRouter()
	router.Use(s.logger.RequestLogger)
	router.HandleFunc("/people", s.HandlePeople).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc(
		"/people/{id}",
		s.HandlePeopleWithId,
	).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	return router
}
