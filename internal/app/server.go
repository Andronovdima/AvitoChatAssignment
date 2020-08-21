package apiserver

import (
	"database/sql"
	userH "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/delivery"
	userR "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/repository"
	userU "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/usecase"
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
