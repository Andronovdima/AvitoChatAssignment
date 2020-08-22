package usecase

import (
	chatRep "github.com/Andronovdima/AvitoChatAssignment/internal/app/chat/repository"
	messageRep "github.com/Andronovdima/AvitoChatAssignment/internal/app/message/repository"
	userRep "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/repository"
	"github.com/Andronovdima/AvitoChatAssignment/internal/models"
	"net/http"
	"strconv"
)

type MessageUsecase struct {
	chatRep    *chatRep.ChatRepository
	userRep    *userRep.UserRepository
	messageRep *messageRep.MessageRepository
}

func NewMessageUsecase(c *chatRep.ChatRepository, u *userRep.UserRepository, m *messageRep.MessageRepository) *MessageUsecase {
	messageUC := &MessageUsecase{
		chatRep:    c,
		userRep:    u,
		messageRep: m,
	}
	return messageUC
}

func (m *MessageUsecase) CreateMessage(message *models.Message) (int64, error) {
	err := new(models.HttpError)

	chatID_, cerr := strconv.ParseInt(message.ChatID, 10, 64)
	if cerr != nil {
		err.StatusCode = http.StatusBadRequest
		err.StringErr = "chat id isn't a number"
	}

	isExistChat := m.chatRep.IsExistByID(chatID_)
	if !isExistChat {
		err.StatusCode = http.StatusBadRequest
		err.StringErr = "chat with this id doesnt exist"
		return -1, err
	}

	isExistUser := m.userRep.IsExistByID(message.AuthorID)
	if !isExistUser {
		err.StatusCode = http.StatusBadRequest
		err.StringErr = "user with this id doesnt exist"
		return -1, err
	}

	isUserInChat := m.chatRep.IsUserInChat(chatID_, message.AuthorID)
	if !isUserInChat {
		err.StatusCode = http.StatusBadRequest
		err.StringErr = "user with this id doesn't consist in the chat"
		return -1, err
	}

	cerr = m.messageRep.Create(message)
	if cerr != nil {
		err.StatusCode = http.StatusInternalServerError
		err.StringErr = cerr.Error()
		return -1, err
	}

	return message.ID, nil
}

func (m *MessageUsecase) GetMessages(chatID *models.ChatID) ([]models.Message, error) {
	err := new(models.HttpError)

	id, cerr := strconv.ParseInt(chatID.ChatID, 10, 64)
	if cerr != nil {
		err.StatusCode = http.StatusBadRequest
		err.StringErr = "chat id isn't a number"
	}

	isExistChat := m.chatRep.IsExistByID(id)
	if !isExistChat {
		err.StatusCode = http.StatusBadRequest
		err.StringErr = "chat with this id doesnt exist"
		return nil, err
	}

	messages, cerr := m.messageRep.GetMessages(id)
	if cerr != nil {
		err.StatusCode = http.StatusInternalServerError
		err.StringErr = cerr.Error()
		return nil, err
	}

	return messages, nil
}
