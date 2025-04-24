package server

import (
	"backend-avanzada/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	people := []*models.PersonResponse{}
	for _, v := range s.DB {
		people = append(people, &models.PersonResponse{
			Nombre:        v.Nombre,
			Edad:          v.Edad,
			FechaCreacion: fmt.Sprint(v.FechaCreacion),
		})
	}
	result, err := json.Marshal(people)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (s *Server) handleGetWithId(w http.ResponseWriter, r *http.Request) {
	matches := PeopleRegexWithId.FindStringSubmatch(r.URL.Path)
	id, err := strconv.ParseInt(matches[1], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	p, exists := s.DB[int(id)]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp := &models.PersonResponse{
		Nombre:        p.Nombre,
		Edad:          p.Edad,
		FechaCreacion: fmt.Sprint(p.FechaCreacion),
	}
	response, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	var p models.PersonRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	person := &models.Person{
		Id:            len(s.DB) + 1,
		Nombre:        p.Nombre,
		Edad:          int(p.Edad),
		FechaCreacion: time.Now(),
	}
	s.DB[person.Id] = person
	pResponse := &models.PersonResponse{
		Nombre:        person.Nombre,
		Edad:          person.Edad,
		FechaCreacion: fmt.Sprint(person.FechaCreacion),
	}
	result, err := json.Marshal(pResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (s *Server) handlePut(w http.ResponseWriter, r *http.Request) {
	var p models.PersonRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	matches := PeopleRegexWithId.FindStringSubmatch(r.URL.Path)
	id, err := strconv.ParseInt(matches[1], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	person, exists := s.DB[int(id)]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	person.Nombre = p.Nombre
	person.Edad = int(p.Edad)
	pResponse := &models.PersonResponse{
		Nombre:        person.Nombre,
		Edad:          person.Edad,
		FechaCreacion: fmt.Sprint(person.FechaCreacion),
	}
	result, err := json.Marshal(pResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	matches := PeopleRegexWithId.FindStringSubmatch(r.URL.Path)
	id, err := strconv.ParseInt(matches[1], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	p, exists := s.DB[int(id)]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(s.DB, p.Id)
	w.WriteHeader(http.StatusNoContent)
}
