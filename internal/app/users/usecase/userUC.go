package usecase

import (
	"../../../models"
	"../repository"
	"net/http"
)

type UserUsecase struct {
	UserRep *repository.UserRepository
}

func NewUserUsecase(us *repository.UserRepository) *UserUsecase {
	UserUsecase := &UserUsecase{
		UserRep: us,
	}
	return UserUsecase
}

func (u *UserUsecase) CreateUser(user *models.User) (int64, error) {
	err := new(models.HttpError)

	isExistUsername := u.UserRep.IsExist(user.Username)
	if isExistUsername {
		err.StatusCode = http.StatusInternalServerError
		err.StringErr = "user with this username already exist, try with another one"
		return -1, err
	}


	cerr := u.UserRep.Create(user)
	if cerr != nil {
		err.StatusCode = http.StatusInternalServerError
		err.StringErr = cerr.Error()
		return -1, err
	}


	return user.ID, nil
}

func (u *UserUsecase) IsExistUser(username string) bool {
	return u.UserRep.IsExist(username)
}

func (u *UserUsecase) GetAllUsers() ([]models.User, error) {
	return u.UserRep.GetAllUsers()
}


