package server

import (
	"backend-avanzada/config"
	"backend-avanzada/logger"
	"backend-avanzada/models"
	"backend-avanzada/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	DB               *gorm.DB
	Config           *config.Config
	Handler          http.Handler
	PeopleRepository repository.Repository[models.Person]
	logger           *logger.Logger
}

func NewServer() *Server {
	s := &Server{
		logger: logger.NewLogger(),
	}
	var config config.Config
	configFile, err := os.ReadFile("config/config.json")
	if err != nil {
		s.logger.Fatal(err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		s.logger.Fatal(err)
	}
	s.Config = &config
	return s
}

func (s *Server) StartServer() {
	fmt.Println("Inicializando base de datos...")
	s.initDB()
	fmt.Println("Inicializando mux...")
	srv := &http.Server{
		Addr:    s.Config.Address,
		Handler: s.router(),
	}
	fmt.Println("Escuchando en el puerto ", s.Config.Address)
	if err := srv.ListenAndServe(); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) initDB() {
	switch s.Config.Database {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			s.logger.Fatal(err)
		}
		s.DB = db
	case "postgres":
		dsn := "host=ep-snowy-sunset-a2d8ilrq.eu-central-1.aws.neon.tech user=neondb_owner password=npg_KlJRLSnBT65j dbname=neondb sslmode=require"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			s.logger.Fatal(err)
		}
		s.DB = db
	}
	fmt.Println("Aplicando migraciones...")
	s.DB.AutoMigrate(&models.Person{})
	s.PeopleRepository = repository.NewPeopleRepository(s.DB)
}
