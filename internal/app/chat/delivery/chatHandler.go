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
	logger      *zap.SugaredLogger
}

func NewChatHandler(m *mux.Router, chat chatUC.ChatUsecase, logger *zap.SugaredLogger) {
	handler := &ChatHandler{
		chatUsecase: chat,
		logger:      logger,
	}

	m.HandleFunc("/chats/add", handler.HandleCreateChat).Methods(http.MethodPost)
	m.HandleFunc("/chats/get", handler.HandleGetListChats).Methods(http.MethodPost)
}

func (c *ChatHandler) HandleCreateChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := r.Body.Close(); err != nil {
			err = errors.Wrapf(err, "ERROR : HandleCreateChat <- Body.Close")
			c.logger.Info(err.Error())
			respond.Error(w, r, http.StatusInternalServerError, err)
			return
		}
	}()

	thisChat := new(models.Chat)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(thisChat)
	if err != nil {
		err = errors.Wrapf(err, "ERROR : HandleCreateChat <- Decode: ")
		c.logger.Info(err.Error())
		err = errors.New("Invalid format data, example: '{'name': 'chat_1', 'users': ['1', '2 ']} ' ")
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := c.chatUsecase.CreateChat(thisChat)
	if err != nil {
		rerr := err.(*models.HttpError)
		c.logger.Info("ERROR : HandleCreateChat <- CreateChat: " + rerr.Error())
		respond.Error(w, r, rerr.StatusCode, rerr)
		return
	}

	c.logger.Info("/chats/add || HTTP: 200")
	respond.Respond(w, r, http.StatusCreated, id)
	return
}

func (c *ChatHandler) HandleGetListChats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := r.Body.Close(); err != nil {
			err = errors.Wrapf(err, "ERROR : HandleGetListChats <- Body.Close")
			c.logger.Info(err.Error())
			respond.Error(w, r, http.StatusInternalServerError, err)
			return
		}
	}()

	GetChatInput := new(models.GetChat)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(GetChatInput)
	if err != nil {
		err = errors.Wrapf(err, "ERROR : HandleGetListChats <- Decode")
		c.logger.Info(err.Error())
		err = errors.New("Invalid format data, example: '{'user': '1'}' ")
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}

	chats, err := c.chatUsecase.GetListByUser(GetChatInput.User)
	if err != nil {
		rerr := err.(*models.HttpError)
		c.logger.Info("ERROR : HandleGetListChats <- GetListByUser: " + rerr.Error())
		respond.Error(w, r, rerr.StatusCode, rerr)
		return
	}

	c.logger.Info("/chats/get || HTTP: 200")
	respond.Respond(w, r, http.StatusOK, chats)
	return
}
