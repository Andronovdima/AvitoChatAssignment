package delivery

import (
	"encoding/json"
	messageUC "github.com/Andronovdima/AvitoChatAssignment/internal/app/message/usecase"
	"github.com/Andronovdima/AvitoChatAssignment/internal/app/respond"
	"github.com/Andronovdima/AvitoChatAssignment/internal/models"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

type MessageHandler struct {
	messageUC messageUC.MessageUsecase
	logger    *zap.SugaredLogger
}

func NewMessageHandler(m *mux.Router, message messageUC.MessageUsecase, logger *zap.SugaredLogger) {
	handler := &MessageHandler{
		messageUC: message,
		logger:    logger,
	}

	m.HandleFunc("/messages/add", handler.HandleCreateMessage).Methods(http.MethodPost)
}

func (m *MessageHandler) HandleCreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := r.Body.Close(); err != nil {
			err = errors.Wrapf(err, "HandleCreateMessage<-Body.Close")
			respond.Error(w, r, http.StatusInternalServerError, err)
			return
		}
	}()

	thisMessage := new(models.Message)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(thisMessage)
	if err != nil {
		err = errors.Wrapf(err, "HandleCreateMessage:")
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := m.messageUC.CreateMessage(thisMessage)
	if err != nil {
		rerr := err.(*models.HttpError)
		respond.Error(w, r, rerr.StatusCode, rerr)
		return
	}

	respond.Respond(w, r, http.StatusCreated, id)
	return
}
