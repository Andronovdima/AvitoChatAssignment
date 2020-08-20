package apiserver

import (
	userH "./users/delivery"
	userR "./users/repository"
	userU "./users/usecase"
	"database/sql"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	Mux    *mux.Router
	Config *Config
	Logger *zap.SugaredLogger
}

func NewServer(config *Config, logger *zap.SugaredLogger) (*Server, error) {
	s := &Server{
		Mux:    mux.NewRouter(),
		Logger: logger,
		Config: config,
	}
	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

func (s *Server) ConfigureServer(db *sql.DB) {
	userRep := userR.NewUserRepository(db)

	userUc := userU.NewUserUsecase(userRep)

	userH.NewUserHandler(s.Mux, *userUc, s.Logger)
}
