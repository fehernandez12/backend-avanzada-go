package models

import "time"

type Person struct {
	Id            int
	Nombre        string
	Edad          int
	FechaCreacion time.Time
}

type PersonRequest struct {
	Nombre string `json:"name"`
	Edad   int32  `json:"age"`
}

type PersonResponse struct {
	Nombre        string `json:"name"`
	Edad          int    `json:"age"`
	FechaCreacion string `json:"created_at"`
}
