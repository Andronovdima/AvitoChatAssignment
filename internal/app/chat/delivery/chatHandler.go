package delivery

import (
	"encoding/json"
	chatUC "github.com/Andronovdima/AvitoChatAssignment/internal/app/chat/usecase"
	"github.com/Andronovdima/AvitoChatAssignment/internal/app/respond"
	"github.com/Andronovdima/AvitoChatAssignment/internal/models"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)
type ChatHandler struct {
	chatUsecase chatUC.ChatUsecase
	logger         *zap.SugaredLogger
}

func NewChatHandler(m *mux.Router, chat chatUC.ChatUsecase, logger *zap.SugaredLogger) {
	handler := &ChatHandler{
		chatUsecase:    chat,
		logger:         logger,
	}

	m.HandleFunc("/chats/add", handler.HandleCreateChat).Methods(http.MethodPost)
}


func (c *ChatHandler) HandleCreateChat (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := r.Body.Close(); err != nil {
			err = errors.Wrapf(err, "HandleCreateChat<-Body.Close")
			respond.Error(w, r, http.StatusInternalServerError, err)
			return
		}
	}()

	thisChat := new(models.Chat)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(thisChat)
	if err != nil {
		err = errors.Wrapf(err, "HandleCreateChat:")
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}


	id, err := c.chatUsecase.CreateChat(thisChat)
	if err != nil {
		rerr := err.(*models.HttpError)
		respond.Error(w, r, rerr.StatusCode, rerr)
		return
	}

	respond.Respond(w, r, http.StatusCreated, id)
	return
}
