package server

import (
	"backend-avanzada/api"
	"backend-avanzada/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) HandlePeople(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGet(w, r)
		return
	case http.MethodPost:
		s.handlePost(w, r)
		return
	}
}

func (s *Server) HandlePeopleWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetWithId(w, r)
		return
	case http.MethodPut:
		s.handlePut(w, r)
		return
	case http.MethodDelete:
		s.handleDelete(w, r)
		return
	}
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	result := []*api.PersonResponse{}
	people, err := s.PeopleRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	for _, v := range people {
		result = append(result, &api.PersonResponse{
			ID:            int(v.ID),
			Nombre:        v.Name,
			Edad:          v.Age,
			FechaCreacion: v.CreatedAt.String(),
		})
	}
	response, err := json.Marshal(result)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}

func (s *Server) handleGetWithId(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}
	p, err := s.PeopleRepository.FindById(int(id))
	if p == nil && err == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path, fmt.Errorf("person with id %d not found", id))
		return
	}
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	resp := &api.PersonResponse{
		ID:            int(p.ID),
		Nombre:        p.Name,
		Edad:          p.Age,
		FechaCreacion: p.CreatedAt.String(),
	}
	response, err := json.Marshal(resp)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	var p api.PersonRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	person := &models.Person{
		Name: p.Nombre,
		Age:  int(p.Edad),
	}
	person, err = s.PeopleRepository.Save(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pResponse := &api.PersonResponse{
		ID:            int(person.ID),
		Nombre:        person.Name,
		Edad:          person.Age,
		FechaCreacion: person.CreatedAt.String(),
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
	var p api.PersonRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	person, err := s.PeopleRepository.FindById(int(id))
	if person == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	person.Name = p.Nombre
	person.Age = int(p.Edad)
	person, err = s.PeopleRepository.Save(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pResponse := &api.PersonResponse{
		ID:            int(person.ID),
		Nombre:        person.Name,
		Edad:          person.Age,
		FechaCreacion: person.CreatedAt.String(),
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
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	person, err := s.PeopleRepository.FindById(int(id))
	if person == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = s.PeopleRepository.Delete(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
