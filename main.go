package main

import (
	"backend-avanzada/models"
	"backend-avanzada/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Inicializando base de datos...")
	db := make(map[int]*models.Person)
	fmt.Println("Inicializando servidor web...")
	mux := http.NewServeMux()
	s := &server.Server{
		DB: db,
	}
	mux.Handle("/", s)
	fmt.Println("Escuchando en el puerto 8000...")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
