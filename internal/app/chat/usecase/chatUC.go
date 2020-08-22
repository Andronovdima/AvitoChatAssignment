package usecase

import (
	chatRep "github.com/Andronovdima/AvitoChatAssignment/internal/app/chat/repository"
	userRep "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/repository"
	"github.com/Andronovdima/AvitoChatAssignment/internal/models"
	"net/http"
)

type ChatUsecase struct {
	chatRep *chatRep.ChatRepository
	userRep *userRep.UserRepository
}

func NewChatUsecase(c *chatRep.ChatRepository, u *userRep.UserRepository) *ChatUsecase {
	chatUsecase := &ChatUsecase{
		chatRep: c,
		userRep: u,
	}
	return chatUsecase
}

func (c *ChatUsecase) CreateChat(chat *models.Chat) (int64, error) {
	err := new(models.HttpError)

	isExistChatName := c.chatRep.IsExist(chat.Name)
	if isExistChatName {
		err.StatusCode = http.StatusBadRequest
		err.StringErr = "chat with this name already exist, try with another one"
		return -1, err
	}

	for _, i := range chat.UsersID {
		isExist := c.userRep.IsExistByID(i)
		if !isExist {
			err.StatusCode = http.StatusBadRequest
			err.StringErr = "at least one id doesnt correct"
			return -1, err
		}

	}

	cerr := c.chatRep.Create(chat)
	if cerr != nil {
		err.StatusCode = http.StatusInternalServerError
		err.StringErr = cerr.Error()
		return -1, err
	}

	return chat.ID, nil
}

func (c *ChatUsecase) GetListByUser(userID string) ([]models.Chat, error) {
	err := new(models.HttpError)

	isExist := c.userRep.IsExistByID(userID)
	if !isExist {
		err.StatusCode = http.StatusBadRequest
		err.StringErr = "user id doesnt correct"
		return nil, err
	}

	chats, cerr := c.chatRep.GetList(userID)
	if cerr != nil {
		err.StatusCode = http.StatusInternalServerError
		err.StringErr = cerr.Error()
		return nil, err
	}

	return chats, nil
}
