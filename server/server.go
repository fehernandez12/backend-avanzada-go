package server

import (
	"backend-avanzada/models"
	"net/http"
	"regexp"
)

var (
	PeopleRegex       = regexp.MustCompile("^/people/*$")
	PeopleRegexWithId = regexp.MustCompile("^/people/([0-9]+)$")
)

type Server struct {
	DB map[int]*models.Person
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && PeopleRegex.MatchString(r.URL.Path):
		s.handleGet(w, r)
		return
	case r.Method == http.MethodGet && PeopleRegexWithId.MatchString(r.URL.Path):
		s.handleGetWithId(w, r)
		return
	case r.Method == http.MethodPost && PeopleRegex.MatchString(r.URL.Path):
		s.handlePost(w, r)
		return
	case r.Method == http.MethodPut && PeopleRegexWithId.MatchString(r.URL.Path):
		s.handlePut(w, r)
		return
	case r.Method == http.MethodDelete && PeopleRegexWithId.MatchString(r.URL.Path):
		s.handleDelete(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
