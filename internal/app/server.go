package apiserver

import (
	"database/sql"
	chatD "github.com/Andronovdima/AvitoChatAssignment/internal/app/chat/delivery"
	chatR "github.com/Andronovdima/AvitoChatAssignment/internal/app/chat/repository"
	chatU "github.com/Andronovdima/AvitoChatAssignment/internal/app/chat/usecase"

	userD "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/delivery"
	userR "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/repository"
	userU "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/usecase"

	messageD "github.com/Andronovdima/AvitoChatAssignment/internal/app/message/delivery"
	messageR "github.com/Andronovdima/AvitoChatAssignment/internal/app/message/repository"
	messageU "github.com/Andronovdima/AvitoChatAssignment/internal/app/message/usecase"

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
	chatRep := chatR.NewChatRepository(db)
	messageRep := messageR.NewMessageRepository(db)

	userUc := userU.NewUserUsecase(userRep)
	chatUC := chatU.NewChatUsecase(chatRep, userRep)
	messageUc := messageU.NewMessageUsecase(chatRep, userRep, messageRep)

	userD.NewUserHandler(s.Mux, *userUc, s.Logger)
	chatD.NewChatHandler(s.Mux, *chatUC, s.Logger)
	messageD.NewMessageHandler(s.Mux, *messageUc, s.Logger)
}
