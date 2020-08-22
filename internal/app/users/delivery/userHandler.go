package delivery

import (
	"encoding/json"
	"github.com/Andronovdima/AvitoChatAssignment/internal/app/respond"
	"github.com/Andronovdima/AvitoChatAssignment/internal/app/users/usecase"
	"github.com/Andronovdima/AvitoChatAssignment/internal/models"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
	Logger      *zap.SugaredLogger
}

func NewUserHandler(m *mux.Router, uc usecase.UserUsecase, logger *zap.SugaredLogger) {
	handler := &UserHandler{
		UserUsecase: uc,
		Logger:      logger,
	}

	m.HandleFunc("/users/add", handler.HandleCreateUser).Methods(http.MethodPost)
}

func (u *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := r.Body.Close(); err != nil {
			err = errors.Wrapf(err, "ERROR : HandleCreateUser<-Body.Close")
			u.Logger.Info(err.Error())
			respond.Error(w, r, http.StatusInternalServerError, err)
			return
		}
	}()

	thisUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(thisUser)
	if err != nil {
		err = errors.Wrapf(err, "ERROR : HandleCreateUser<-Decode: ")
		u.Logger.Info(err.Error())
		err = errors.New("Invalid format data, example: '{'username': 'user_1'}' ")
		respond.Error(w, r, http.StatusBadRequest, err)
		return
	}
	us, err := u.UserUsecase.CreateUser(thisUser)
	if err != nil {
		rerr := err.(*models.HttpError)
		u.Logger.Info("ERROR : HandleCreateUser<-Create: " + rerr.Error())
		respond.Error(w, r, rerr.StatusCode, rerr)
		return
	}

	u.Logger.Info("/users/add || HTTP: 200")
	respond.Respond(w, r, http.StatusCreated, us)
	return
}
