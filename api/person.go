package api

type PersonRequest struct {
	Nombre string `json:"name"`
	Edad   int32  `json:"age"`
}

type PersonResponse struct {
	ID            int    `json:"person_id"`
	Nombre        string `json:"name"`
	Edad          int    `json:"age"`
	FechaCreacion string `json:"created_at"`
}

type ErrorResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Message     string `json:"message"`
}
